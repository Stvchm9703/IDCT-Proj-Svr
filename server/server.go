package server

import (
	pb "RoomStatus/pkg/protos"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoomWorker struct {
}
type RoomStatusServer struct {
	pb.UnimplementedRoomStatusServer
	hashmap map[string]RoomWorker
}

func (*RoomStatusServer) CreateRoom(ctx context.Context, req *pb.Empty) (*pb.Room, error) {

	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (*RoomStatusServer) GetRoomList(ctx context.Context, req *pb.RoomListRequest) (*pb.RoomListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomList not implemented")
}
func (*RoomStatusServer) GetRoomCurrentInfo(ctx context.Context, req *pb.RoomRequest) (*pb.Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomCurrentInfo not implemented")
}
func (*RoomStatusServer) GetRoomStream(req *pb.RoomRequest, srv pb.RoomStatus_GetRoomStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRoomStream not implemented")
}
func (*RoomStatusServer) UpdateRoomStatus(ctx context.Context, req *pb.CellStatus) (*pb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoomStatus not implemented")
}
func (*RoomStatusServer) DeleteRoom(ctx context.Context, req *pb.RoomRequest) (*pb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
