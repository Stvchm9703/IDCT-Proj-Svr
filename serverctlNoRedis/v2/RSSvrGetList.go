package v2

import (
	pb "RoomStatus/proto/v2"
	"context"
	"log"
	"reflect"
	"time"
)

// GetRoomList :
func (b *RoomStatusBackend) GetRoomList(ctx context.Context, req *pb.RoomListReq) (res *pb.RoomListResp, err error) {
	printReqLog(ctx, req)

	var RmList []*pb.Room

	log.Println("list:", RmList)
	log.Println("typeof:", reflect.TypeOf(RmList))
	var tmp []*pb.Room
	for _, v := range b.Roomlist {
		tmp = append(tmp, &v.Room)
	}
	res = &pb.RoomListResp{
		Timestamp: time.Now().String(),
		Result:    tmp,
		ErrorMsg:  nil,
	}
	return
}
