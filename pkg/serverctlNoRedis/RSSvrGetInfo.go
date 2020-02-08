package serverctlNoRedis

import (
	"RoomStatus/common"
	pb "RoomStatus/proto"
	"context"
	"log"
	"time"
)

// GetRoomInfo :
func (b *RoomStatusBackend) GetRoomInfo(ctx context.Context, req *pb.RoomReq) (*pb.RoomResp, error) {
	common.PrintReqLog(ctx, req)
	for _, v := range b.Roomlist {
		log.Println(v.Key)
		if (v).Key == req.Key {
			return &pb.RoomResp{
				Timestamp: time.Now().String(),
				ResponseMsg: &pb.RoomResp_RoomInfo{
					RoomInfo: &v.Room,
				},
			}, nil
		}
	}
	return &pb.RoomResp{
		Timestamp: time.Now().String(),
		ResponseMsg: &pb.RoomResp_Error{
			Error: &pb.ErrorMsg{
				MsgInfo: "RoomNoFound",
				MsgDesp: "No Room Found With Key :" + req.Key,
			},
		},
	}, nil
}
