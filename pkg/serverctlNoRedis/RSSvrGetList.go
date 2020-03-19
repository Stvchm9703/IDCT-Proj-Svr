package serverctlNoRedis

import (
	"RoomStatus/common"
	pb "RoomStatus/proto"
	"context"
	"time"
)

// GetRoomList :
func (b *RoomStatusBackend) GetRoomList(ctx context.Context, req *pb.RoomListReq) (res *pb.RoomListResp, err error) {
	// md, _ := metadata.FromIncomingContext(ctx)
	// log.Println(md)
	common.PrintReqLog(ctx, "Get-Room-List", req)

	var tmp []*pb.Room
	for _, v := range b.Roomlist {
		tmp = append(tmp, v)
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
