package serverctlNoRedis

import (
	pb "RoomStatus/proto"
	"context"
	"time"
)

// GetRoomList :
func (b *RoomStatusBackend) GetRoomList(ctx context.Context, req *pb.RoomListReq) (res *pb.RoomListResp, err error) {
	printReqLog(ctx, req)

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
