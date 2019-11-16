package serverctl

import (
	pb "RoomStatus/proto"
	"context"
	"errors"
	"log"
	"time"

	types "github.com/gogo/protobuf/types"
)

// QuitRoom :
// !incomplete
func (b *RoomStatusBackend) QuitRoom(ctx context.Context, req *pb.RoomCreateRequest) (*types.Empty, error) {
	start := time.Now()
	b.mu.Lock()
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()
	var tmpr *pb.Room
	for _, v := range b.Roomlist {
		if v.HostId == req.HostId || v.DuelerId == req.HostId {
			tmpr = v
			break
		}
	}
	if tmpr == nil {
		return nil, errors.New("RoomNotExist")
	}
	if tmpr.HostId == req.HostId {
		// Self quit, not start play yet

	} else {
		// Dueler quit game,
	}
	return nil, errors.New("NotImplement")

	// return nil,
}
