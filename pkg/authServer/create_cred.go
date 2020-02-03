//
package authServer

import (
	pb "RoomStatus/proto"
	"errors"
)

// CreateCred(*CredReq, CreditsAuth_CreateCredServer) error

func (CAB *CreditsAuthBackend) CreateCred(req *pb.CredReq, stream pb.CreditsAuth_CreateCredServer) error {
	CAB.mu.Lock()
	defer CAB.mu.Unlock()

	// var result []*UserCredMod
	tmpUser := &UserCredMod{
		Username: &req.Username,
		Password: &req.Password,
	}

	if !CAB.DBconn.NewRecord(tmpUser) {
		return errors.New("UserIsExist")
	}
	return errors.New("NOT_IMPLEMENT")
}
