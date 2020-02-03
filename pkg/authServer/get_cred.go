//
package authServer

import (
	pb "RoomStatus/proto"
	"bytes"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// GetCred(context.Context, *CredReq) (*CreateCredResp, error)

func (CAB *CreditsAuthBackend) GetCred(ctx context.Context, cq *pb.CredReq) (*pb.CreateCredResp, error) {
	CAB.mu.Lock()
	defer CAB.mu.Unlock()

	var result []*UserCredMod
	CAB.DBconn.Where(&UserCredMod{
		Username: cq.Username,
	}).Find(&result)

	if len(result) != 1 {
		return nil, errors.New("UNKNOWN_DB_RECORD")
	}
	pwParse, _ := bcrypt.GenerateFromPassword([]byte(cq.Password), bcrypt.DefaultCost)
	actual, _ := bcrypt.GenerateFromPassword([]byte(result[0].Password), bcrypt.DefaultCost)
	if !bytes.Equal(actual, pwParse) {
		return nil, errors.New("PASSWORD_INVALID")
	}

	return &pb.CreateCredResp{
		Code:     200,
		File:     result[0].KeyPem,
		ErrorMsg: nil,
	}, nil
}
