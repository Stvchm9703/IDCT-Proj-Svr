package serverctlNoRedis

import (
	pb "RoomStatus/proto"
	"context"
	"errors"
	"log"

	types "github.com/gogo/protobuf/types"
)

// DeleteRoom :
func (b *RoomStatusBackend) DeleteRoom(ctx context.Context, req *pb.RoomRequest) (*types.Empty, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	b.mu.Lock()
	defer b.mu.Unlock()

	done := false
	for k, v := range b.Roomlist {
		if v.Key == req.Key {
			log.Println(b.Roomlist[k])
			b.Roomlist = append(b.Roomlist[:k], b.Roomlist[k+1:]...)
			done = true
		}
	}
	if !done {
		return nil, errors.New("RoomNotExist")
	}
	log.Println("b.RoomList", b.Roomlist)
	return nil, nil
}

//
