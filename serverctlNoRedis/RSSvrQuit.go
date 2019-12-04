package serverctlNoRedisl

import (
	pb "RoomStatus/proto"
	"context"
	"errors"

	types "github.com/gogo/protobuf/types"
)

func (b *RoomStatusBackend) QuitRoom(ctx context.Context, req *pb.RoomCreateRequest) (*types.Empty, error) {
	return nil, errors.New("NotImplement")
}
