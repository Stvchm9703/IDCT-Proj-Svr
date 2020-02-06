//
package authServer

import (
	pb "RoomStatus/proto"
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// CreateCred(*CredReq, CreditsAuth_CreateCredServer) error

func (CAB *CreditsAuthBackend) CreateCred(ctx context.Context, req *pb.CredReq) (*pb.CheckCredResp, error) {
	CAB.mu.Lock()
	defer CAB.mu.Unlock()

	pwParse, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	tmpUser := &UserCredMod{
		Username: req.Username,
		Password: string(pwParse),
	}

	count := 0
	if CAB.DBconn.Model(&UserCredMod{}).Where(tmpUser).Count(&count); count > 0 {
		return nil, errors.New("UserIsExist")
	}

	if err := createCredTx(CAB.DBconn, tmpUser); err != nil {
		return nil, err
	}

	return &pb.CheckCredResp{
		ResponseCode: 200,
		ErrorMsg:     nil,
	}, nil
}

func createCredTx(dbc *gorm.DB, ucm *UserCredMod) error {
	tx := dbc.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(ucm).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
