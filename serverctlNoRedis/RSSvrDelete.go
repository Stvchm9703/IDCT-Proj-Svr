package serverctlNoRedis

import (
	pb "RoomStatus/proto"
	"context"
	"errors"
	"log"
	"time"
)

// DeleteRoom :
func (b *RoomStatusBackend) DeleteRoom(ctx context.Context, req *pb.RoomReq) (*pb.RoomResp, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	b.mu.Lock()
	defer b.mu.Unlock()

	done := false
	var room_tmp *RoomMgr

	for k, v := range b.Roomlist {
		if v.Key == req.Key {
			log.Println(b.Roomlist[k])
			room_tmp = b.Roomlist[k]
			b.Roomlist = append(b.Roomlist[:k], b.Roomlist[k+1:]...)
			done = true
		}
	}
	if !done {
		return &pb.RoomResp{
			Timestamp: time.Now().String(),
			ResponseMsg: &pb.RoomResp_Error{
				Error: &pb.ErrorMsg{
					MsgInfo: "RoomNotExist",
					MsgDesp: "Room<" + req.Key + "> is not exist",
				},
			},
		}, errors.New("RoomNotExist")
	}
	room_tmp.BroadCast("RoomSvrMgr",
		&pb.CellStatusResp{
			UserId:    "RoomSvrMgr",
			Key:       room_tmp.Room.Key,
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
			RoomInfo: &room_tmp.Room,
		},
	}, nil

}

//
