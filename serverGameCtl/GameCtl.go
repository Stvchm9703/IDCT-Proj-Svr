package serverGameCtl

import (
	cm "RoomStatus/common"
	cf "RoomStatus/config"
	pb "RoomStatus/proto"
	rd "RoomStatus/store/redis"
	"context"
	"strconv"
	"sync"

	types "github.com/gogo/protobuf/types"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

type Backend struct {
	mu       *sync.RWMutex
	Roomlist []*pb.Room
	redhdlr  []*rd.RdsCliBox
	CoreKey  string
}

var _ pb.RoomStatusServer = (*Backend)(nil)

// New : Create new backend
func New(conf *cf.ConfTmp) *Backend {
	ck := "RSCore" + cm.HashText(conf.Server.IP)
	rdfl := []*rd.RdsCliBox{}
	for i := 0; i < conf.Database.WorkerNode; i++ {
		rdf := rd.RdsCliBox{
			CoreKey: ck,
			Key:     "wKU" + cm.HashText("num"+strconv.Itoa(i)),
		}
		if _, err := rdf.Connect(conf); err == nil {
			rdfl = append(rdfl, &rdf)
		}
	}

	return &Backend{
		mu:       &sync.RWMutex{},
		Roomlist: nil,
		redhdlr:  rdfl,
	}
}

// 	Impletement from GameCtl.pb.go(auto-gen file)
// 		CreateRoom(context.Context, *types.Empty) (*Room, error)
// 		GetRoomList(context.Context, *RoomListRequest) (*RoomListResponse, error)
// 		GetRoomCurrentInfo(context.Context, *RoomRequest) (*Room, error)
// 		GetRoomStream(*RoomRequest, RoomStatus_GetRoomStreamServer) error
// 		UpdateRoomStatus(context.Context, *CellStatus) (*types.Empty, error)
// 		DeleteRoom(context.Context, *RoomRequest) (*types.Empty, error)

// CreateRoom :
func (b *Backend) CreateRoom(context.Context, *types.Empty) (*pb.Room, error) {
	return nil, nil
}

// GetRoomList :
func (b *Backend) GetRoomList(ctx context.Context, req *pb.RoomListRequest) (*pb.RoomListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomList not implemented")
}

// GetRoomCurrentInfo:
func (b *Backend) GetRoomCurrentInfo(ctx context.Context, req *pb.RoomRequest) (*pb.Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomCurrentInfo not implemented")
}

// GetRoomStream:
func (b *Backend) GetRoomStream(req *pb.RoomRequest, srv pb.RoomStatus_GetRoomStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRoomStream not implemented")
}

// UpdateRoomStatus:
func (b *Backend) UpdateRoomStatus(ctx context.Context, req *pb.CellStatus) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoomStatus not implemented")
}

// DeleteRoom:
func (b *Backend) DeleteRoom(ctx context.Context, req *pb.RoomRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
