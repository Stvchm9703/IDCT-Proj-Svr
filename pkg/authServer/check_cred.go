//
package authServer

import (
	pb "RoomStatus/proto"
	"context"
)

// CheckCred(context.Context, *CredReq) (*CheckCredResp, error)

func (CAB *CreditsAuthBackend) CheckCred(ctx context.Context, crq *pb.CredReq) (*pb.CheckCredResp, error) {

	return nil, nil
}
