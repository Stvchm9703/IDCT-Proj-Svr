//
package authServer

import (
	pb "RoomStatus/proto"
	"bytes"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// CheckCred(context.Context, *CredReq) (*CheckCredResp, error)

func (CAB *CreditsAuthBackend) CheckCred(ctx context.Context, crq *pb.CredReq) (*pb.CheckCredResp, error) {
	var result []*UserCredMod
	CAB.DBconn.Where(&UserCredMod{
		Username: crq.Username,
	}).Find(&result)

	if len(result) != 1 {
		return nil, errors.New("UNKNOWN_DB_RECORD")
	}
	pwParse, _ := bcrypt.GenerateFromPassword([]byte(crq.Password), bcrypt.DefaultCost)
	actual, _ := bcrypt.GenerateFromPassword([]byte(result[0].Password), bcrypt.DefaultCost)
	if !bytes.Equal(actual, pwParse) {
		return &pb.CheckCredResp{
			ResponseCode: 500,
			ErrorMsg: &pb.ErrorMsg{
				MsgInfo: "PASSWORD_INVALID",
				MsgDesp: "password invalid",
			},
		}, nil
	}

	return &pb.CheckCredResp{
		ResponseCode: 200,
		ErrorMsg:     nil,
	}, nil
}
