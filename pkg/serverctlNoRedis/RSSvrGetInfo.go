package serverctlNoRedis

import (
	"RoomStatus/pkg/common"
	pb "RoomStatus/proto"
	"context"
	"log"
	"time"
)

// GetRoomInfo :
func (b *RoomStatusBackend) GetRoomInfo(ctx context.Context, req *pb.RoomReq) (*pb.RoomResp, error) {
	common.PrintReqLog(ctx, "GetRoom-Info", req)
	for _, v := range b.Roomlist {
		log.Println(v.Key)
		if (v).Key == req.Key {
			return &pb.RoomResp{
				Timestamp: time.Now().String(),
				ResponseMsg: &pb.RoomResp_RoomInfo{
					RoomInfo: v,
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
