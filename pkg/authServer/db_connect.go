package authServer

import (
	"RoomStatus/config"
	"log"
	"strconv"
	"time"

	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (CAB *CreditsAuthBackend) InitDB(config *config.CfTDatabase) (*gorm.DB, error) {
	log.Println("in InitDB")
	CAB.mu.Lock()
	defer CAB.mu.Unlock()
	log.Println("\t Open DB")
	dburl := config.Connector + "://" +
		config.Username + ":" + config.Password +
		"@" + config.Host + ":" + strconv.Itoa(config.Port) +
		"/" + config.Database +
		"?sslmode=disable"
	db, err := gorm.Open(config.Connector, dburl)

	// "host="+config.Host+
	// 	" port="+strconv.Itoa(config.Port)+
	// 	" user="+config.Username+
	// 	" dbname="+config.Database+
	// 	" password="+config.Password+
	// 	" sslmode=disable"

	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Open Success")
	CAB.DBconn = db

	_, err = init_check(db)

	if err != nil {
		return nil, err
	}
	return db, nil

}

func (CAB *CreditsAuthBackend) CloseDB() (bool, error) {

	log.Println("BeforeCloseDB")

	if CAB.DBconn == nil {
		log.Println("CloseSession")
		return false, errors.New("NIL_SESSION")
	}
	err := CAB.DBconn.Close()
	if err != nil {
		return false, err
	}
	return true, nil
}

type UserCredMod struct {
	gorm.Model
	Username string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`
}

func (UserCredMod) TableName() string {
	return "user_cred"
}

type CredSessionMod struct {
	gorm.Model
	UserId     uint   `gorm:"column:user_id;"`
	DeviceName string `gorm:"column:device_name;type:varchar(100)"`
	Token      string
	Timeout    time.Time
}

func (CredSessionMod) TableName() string {
	return "cred_session"
}

func init_check(dbc *gorm.DB) (bool, error) {
	log.Println("start checking")
	if dbc == nil {
		return false, errors.New("NULL_SESSION")
	}

	log.Println("table-", UserCredMod{}.TableName())
	if dbc.HasTable(&UserCredMod{}) == false {
		log.Println("table->create", UserCredMod{}.TableName())

		dbc.CreateTable(&UserCredMod{})
	}

	log.Println("table-", CredSessionMod{}.TableName())

	if dbc.HasTable(&CredSessionMod{}) == false {
		log.Println("table->create", CredSessionMod{}.TableName())

		dbc.CreateTable(&CredSessionMod{})
		// dbc.Model(&CredSessionMod{}).AddForeignKey("user_id", "user_cred(id)", "CASCADE", "CASCADE")
	}
	return true, nil
}
