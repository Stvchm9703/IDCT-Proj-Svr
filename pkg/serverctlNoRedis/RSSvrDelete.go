package serverctlNoRedis

import (
	cm "RoomStatus/pkg/common"
	pb "RoomStatus/proto"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

// DeleteRoom :
func (b *RoomStatusBackend) DeleteRoom(ctx context.Context, req *pb.RoomReq) (*pb.RoomResp, error) {
	cm.PrintReqLog(ctx, "delete-room", req)

	b.mu.Lock()
	defer b.mu.Unlock()

	done := false
	var room_tmp *pb.Room

	for k, v := range b.Roomlist {
		if v.Key == req.Key {
			log.Println(b.Roomlist[k])
			room_tmp = b.Roomlist[k]
			b.Roomlist = append(b.Roomlist[:k], b.Roomlist[k+1:]...)
			done = true
		}
	}
	if !done {
		return nil, errors.New("RoomNotExist")
	}
	// !Broadcast
	go b.BroadCast(&pb.CellStatusResp{
		UserId:    "RoomSvrMgr",
		Key:       room_tmp.Key,
		Timestamp: time.Now().String(),
		Status:    510,
		ResponseMsg: &pb.CellStatusResp_ErrorMsg{
			ErrorMsg: &pb.ErrorMsg{
				MsgInfo: "RoomClose",
				MsgDesp: "Room Close By Server with Request <UserID:" + req.Key + ">",
			},
		},
	})

	log.Println("b.RoomList", b.Roomlist)
	return &pb.RoomResp{
		Timestamp: time.Now().String(),
		ResponseMsg: &pb.RoomResp_RoomInfo{
			RoomInfo: room_tmp,
		},
	}, nil

}

//
func (rsb *RoomStatusBackend) RemoveRoom(req *pb.RoomReq) error {
	done := false
	var room_tmp *pb.Room

	for k, v := range rsb.Roomlist {
		if v.Key == req.Key {
			log.Println(rsb.Roomlist[k])
			room_tmp = rsb.Roomlist[k]
			rsb.Roomlist = append(rsb.Roomlist[:k], rsb.Roomlist[k+1:]...)
			done = true
		}
	}
	if !done {
		return errors.New("RoomNotExist")
	}
	// !Broadcast
	go rsb.BroadCast(&pb.CellStatusResp{
		UserId:    "RoomSvrMgr",
		Key:       room_tmp.Key,
		Timestamp: time.Now().String(),
		Status:    510,
		ResponseMsg: &pb.CellStatusResp_ErrorMsg{
			ErrorMsg: &pb.ErrorMsg{
				MsgInfo: "RoomClose",
				MsgDesp: fmt.Sprintf("Room Close By Server with Request <RoomID:%s>", req.Key),
			},
		},
	})

	log.Println("b.RoomList", rsb.Roomlist)
	return nil

}
