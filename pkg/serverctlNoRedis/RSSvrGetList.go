package serverctlNoRedis

import (
	"RoomStatus/common"
	pb "RoomStatus/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc/metadata"
)

// GetRoomList :
func (b *RoomStatusBackend) GetRoomList(ctx context.Context, req *pb.RoomListReq) (res *pb.RoomListResp, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Println(md)
	common.PrintReqLog(ctx, req)

	var tmp []*pb.Room
	for _, v := range b.Roomlist {
		var y pb.Room
		y = v.Room
		tmp = append(tmp, &y)
	}
	// log.Println("list:", tmp)
	// log.Println("typeof:", reflect.TypeOf(tmp))

	res = &pb.RoomListResp{
		Timestamp: time.Now().String(),
		Result:    tmp,
		ErrorMsg:  nil,
	}
	return
}
