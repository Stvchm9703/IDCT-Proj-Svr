package serverGameCtl

import (
	pb "RoomStatus/proto"
	"context"
	"log"
)

// pb.GetRoomStream
// func (*UnimplementedRoomStatusServer) GetRoomStream(req *RoomRequest, srv RoomStatus_GetRoomStreamServer) error {
// 	return status.Errorf(codes.Unimplemented, "method GetRoomStream not implemented")
// }

// GetRoomList :
func (b *RoomStatusBackend) GetRoomStream(ctx context.Context, req *pb.RoomRequest, srv RoomStatusBackend) (res *pb.RoomListResponse, err error) {
	printReqLog(ctx, req)
	// ===== Worker Start ======
	pl := &WkTask{
		In:  req,
		Out: make(chan interface{})}
	if err = b.getListWk.Invoke(pl); err != nil {
		log.Println("err in GetList Wk", err)
		return
	}
	// ====== Worker End =======
	rm := <-(pl).Out
	fg := rm.([]*pb.Room)
	res = &pb.RoomListResponse{
		Result: fg,
	}
	return
}
