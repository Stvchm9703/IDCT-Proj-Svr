package serverctlNoRedisl

import (
	pb "RoomStatus/proto"
	"context"
)

// GetRoomInfo :
func (b *RoomStatusBackend) GetRoomInfo(ctx context.Context, req *pb.RoomRequest) (*pb.Room, error) {

	printReqLog(ctx, req)

	for _, v := range b.Roomlist {
		if &v.Key == &req.Key {
			return v, nil
		}
	}
	return nil, nil
}
