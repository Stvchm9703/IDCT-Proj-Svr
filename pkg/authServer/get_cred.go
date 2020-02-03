//
package authServer

import (
	pb "RoomStatus/proto"
	"context"
	"errors"
)

// GetCred(context.Context, *CredReq) (*CreateCredResp, error)

func (CAB *CreditsAuthBackend) GetCred(ctx context.Context, cq *pb.CredReq) (*pb.CreateCredResp, error) {
	CAB.mu.Lock()
	defer CAB.mu.Unlock()

	var result []*UserCredMod
	CAB.DBconn.Where(&UserCredMod{
		Username: &cq.Username,
		Password: &cq.Password,
	}).Find(&result)

	if len(result) != 1 {
		return &pb.CreateCredResp{
			Code: 500,
			File: nil,
			ErrorMsg: &pb.ErrorMsg{
				MsgInfo: "UNKNOWN_DB_RECORD",
				MsgDesp: "DB Record seem have no exist good record for this account",
			},
		}, errors.New("UNKNOWN_DB_RECORD")
	}

	return &pb.CreateCredResp{
		Code:     200,
		File:     *result[0].KeyPem,
		ErrorMsg: nil,
	}, nil
}
