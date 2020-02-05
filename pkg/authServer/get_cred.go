//
package authServer

import (
	"RoomStatus/insecure"
	pb "RoomStatus/proto"
	"context"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(codes.Internal, "UNKNOWN_DB_RECORD")
	}

	err := bcrypt.CompareHashAndPassword([]byte(result[0].Password), []byte(cq.Password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, ("PASSWORD_INVALID"))
	}

	return &pb.CreateCredResp{
		Code:     200,
		ErrorMsg: nil,
		File:     insecure.GetCertPemFile(),
	}, nil
}
