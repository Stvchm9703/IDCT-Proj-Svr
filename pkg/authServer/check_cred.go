//
package authServer

import (
	pb "RoomStatus/proto"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	err := bcrypt.CompareHashAndPassword([]byte(result[0].Password), []byte(crq.Password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, ("PASSWORD_INVALID"))
	}

	return &pb.CheckCredResp{
		ResponseCode: 200,
		ErrorMsg:     nil,
	}, nil
}
