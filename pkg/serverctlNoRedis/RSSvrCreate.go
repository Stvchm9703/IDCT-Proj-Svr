package serverctlNoRedis

import (
	cm "RoomStatus/common"
	pb "RoomStatus/proto"
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

// CreateRoom :
func (b *RoomStatusBackend) CreateRoom(ctx context.Context, req *pb.RoomCreateReq) (*pb.RoomResp, error) {
	cm.PrintReqLog(ctx, req)
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, vr := range b.Roomlist {
		if vr.HostId == req.UserId || vr.DuelerId == req.UserId {
			return &pb.RoomResp{
				Timestamp: time.Now().String(),
				ResponseMsg: &pb.RoomResp_Error{
					Error: &pb.ErrorMsg{
						MsgInfo: "GameSetExistWithCurrentPlayer",
						MsgDesp: "Current Player <" + req.UserId + "> is in other Room <" + vr.Key + ">",
					},
				},
			}, errors.New("GameSetExistWithCurrentPlayer")
		}
	}

	tmptime := time.Now().String() + req.UserId
	var f = ""
	for {
		f = cm.HashText(tmptime)

		// ------
		var l []*string
		for _, v := range b.Roomlist {
			if v.Key == f {
				l = append(l, &v.Key)
			}
		}
		if len(l) == 0 {
			break
		}
		// -----
	}
	rmTmp := pb.Room{
		Key:        "Rm" + f,
		HostId:     req.UserId,
		DuelerId:   "",
		Status:     0,
		Round:      0,
		Cell:       -1,
		CellStatus: nil,
	}
	rmTmp1 := RoomMgr{
		Room:            rmTmp,
		get_only_stream: make(map[string]*pb.RoomStatus_GetRoomStreamServer),
		conn_pool:       &sync.Map{},
	}

	b.Roomlist = append(b.Roomlist, &rmTmp1)
	log.Println("Created Room : <", rmTmp)
	return &pb.RoomResp{
		Timestamp: time.Now().String(),
		ResponseMsg: &pb.RoomResp_RoomInfo{
			RoomInfo: &rmTmp,
		},
	}, nil
}
