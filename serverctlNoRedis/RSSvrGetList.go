package serverctlNoRedisl

import (
	pb "RoomStatus/proto"
	"context"
	"log"
	"reflect"
)

// GetRoomList :
func (b *RoomStatusBackend) GetRoomList(ctx context.Context, req *pb.RoomListRequest) (res *pb.RoomListResponse, err error) {
	printReqLog(ctx, req)

	var RmList []*pb.Room

	log.Println("list:", RmList)
	log.Println("typeof:", reflect.TypeOf(RmList))
	res = &pb.RoomListResponse{
		Result: b.Roomlist,
	}
	return
}
