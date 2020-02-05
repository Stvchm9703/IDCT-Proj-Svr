package serverctlNoRedis

import (
	"RoomStatus/config"
	aSvr "RoomStatus/pkg/authServer"
	"context"
	"errors"
	"log"
	"reflect"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

// /https://davidsbond.github.io/2019/06/14/creating-grpc-interceptors-in-go.html
var (
	auth_db  *gorm.DB
	base_key []byte
)

func (this *RoomStatusBackend) InitDB(config *config.CfTDatabase) (*gorm.DB, error) {
	log.Println("in InitDB")
	this.mu.Lock()
	defer this.mu.Unlock()
	log.Println("\t Open DB")
	dburl := config.Connector + "://" +
		config.Username + ":" + config.Password +
		"@" + config.Host + ":" + strconv.Itoa(config.Port) +
		"/" + config.Database +
		"?sslmode=disable"
	db, err := gorm.Open(config.Connector, dburl)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Open Success")
	auth_db = db
	base_key = []byte(time.Now().String() + "/" + config.Database)
	_, err = init_check(db)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func init_check(dbc *gorm.DB) (bool, error) {
	log.Println("start checking")
	if dbc == nil {
		return false, errors.New("NULL_SESSION")
	}

	log.Println("table-", aSvr.UserCredMod{}.TableName())
	if dbc.HasTable(&aSvr.UserCredMod{}) == false {
		log.Println("table->create", aSvr.UserCredMod{}.TableName())

		dbc.CreateTable(&aSvr.UserCredMod{})
	}

	log.Println("table-", aSvr.CredSessionMod{}.TableName())

	if dbc.HasTable(&aSvr.CredSessionMod{}) == false {
		log.Println("table->create", aSvr.CredSessionMod{}.TableName())

		dbc.CreateTable(&aSvr.CredSessionMod{})
		dbc.Model(&aSvr.CredSessionMod{}).AddForeignKey("user_id", "user_cred(id)", "CASCADE", "CASCADE")
	}
	return true, nil
}

func (this *RoomStatusBackend) CloseDB() (bool, error) {
	log.Println("BeforeCloseDB")
	if auth_db == nil {
		log.Println("CloseSession")
		return false, errors.New("NIL_SESSION")
	}
	err := auth_db.Close()
	if err != nil {
		return false, err
	}
	return true, nil
}

// ----------------------------------------------------------------------------
// Token Hash

func parseToken(tokenString string) (*jwt.Token, error) {
	testSet, err := jwt.Parse(
		tokenString,
		func(tk *jwt.Token) (interface{}, error) {
			return base_key, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if testSet.Valid {
		return testSet, nil
	}
	return nil, errors.New("INVALID")
}

func userClaimFromToken(chaim *jwt.StandardClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, chaim)
	tokenString, _ := token.SignedString(base_key)
	return tokenString
}

func updateDBToken(userSet *aSvr.UserCredMod, tokenTime *time.Time, tokenString *string) error {
	credMod := []*aSvr.CredSessionMod{}
	if userSet != nil {
		auth_db.Where("user_id = ?", userSet.ID).Find(&credMod)
	} else {
		auth_db.Where("token = ?", tokenString).Find(&credMod)
	}
	if len(credMod) == 0 {
		err := auth_db.Create(&aSvr.CredSessionMod{
			UserId:  userSet.ID,
			Token:   *tokenString,
			Timeout: *tokenTime,
		}).Error
		if err != nil {
			log.Println(err)
		}
		return err
	} else if len(credMod) == 1 {
		err := auth_db.Model(credMod[0]).Update(&aSvr.CredSessionMod{
			Token:   *tokenString,
			Timeout: *tokenTime,
		}).Where(
			"user_id = ?", userSet.ID,
		).Error
		return err
	}

	return status.Error(codes.Unknown, "RECORD_ISSUE")

}

// ----------------------------------------------------------------------------

func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("<<<<<< Start of TokenCheck")

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")

	if err != nil {
		log.Println("grpc_auth.AuthFromMD:", err)
	}
	timeSnap := time.Now()
	regenString := ""
	tokenInfo, err := parseToken(token)
	if err != nil {
		md, exist := metadata.FromIncomingContext(ctx)
		if !exist {
			return nil, status.Errorf(codes.Unauthenticated, "META_ERROR")
		}

		var appKey, appSecret string

		if key, ok := md["username"]; ok {
			appKey = key[0]
		}
		if secret, ok := md["password"]; ok {
			appSecret = secret[0]
		}
		mod, err := checkBasicAuth(appKey, appSecret)
		if err != nil {
			return nil, err
		}
		// d, _ := time.ParseDuration("1d")
		newChaims := &jwt.StandardClaims{
			ExpiresAt: timeSnap.AddDate(0, 0, 1).Unix(),
			IssuedAt:  timeSnap.Unix(),
			Issuer:    "RoomSvr",
		}
		tokenInfo := jwt.NewWithClaims(jwt.SigningMethodHS256, newChaims)
		regenString, _ = tokenInfo.SignedString(base_key)

		dayAfter := timeSnap.AddDate(0, 0, 1)
		err1 := updateDBToken(mod, &dayAfter, &regenString)
		log.Println("err1:", err1)
	} else {
		regenString = token
	}

	grpc_ctxtags.Extract(ctx).Set("auth.sub", regenString)

	newCtx := context.WithValue(ctx, reflect.TypeOf(tokenInfo), &tokenInfo)
	grpc.SetHeader(newCtx, metadata.Pairs("Authorization", "Bearer "+regenString))
	grpc.SetHeader(newCtx, metadata.Pairs("username", ""))
	grpc.SetHeader(newCtx, metadata.Pairs("password", ""))

	log.Println(">>>>>> End of TokenCheck")
	return handler(newCtx, req)
}

func checkBasicAuth(username string, password string) (*aSvr.UserCredMod, error) {
	if username == "" || password == "" {
		return nil, status.Error(codes.InvalidArgument, "EMPTY_USERNAME")
	}
	if auth_db == nil {
		return nil, status.Error(codes.Internal, "DB_SESSION_NIL")
	}
	var ucm []*aSvr.UserCredMod

	err := auth_db.Model(
		&aSvr.UserCredMod{},
	).Where(&aSvr.UserCredMod{
		Username: username,
	}).Find(&ucm).Error

	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if len(ucm) != 1 {
		log.Println(ucm)
		return nil, status.Error(codes.Internal, ("UNKNOWN_DB_RECORD"))
	}
	err = bcrypt.CompareHashAndPassword([]byte(ucm[0].Password), []byte(password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "INVALID_ACCOUNT")
	}
	return ucm[0], nil
}

// https://davidsbond.github.io/2019/06/14/creating-grpc-interceptors-in-go.html
