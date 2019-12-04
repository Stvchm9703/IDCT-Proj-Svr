package serverctl

import (
	pb "RoomStatus/proto"
	"context"
	"log"
	"time"

	"github.com/gogo/status"
)

// GetRoomInfo :
func (b *RoomStatusBackend) GetRoomInfo(ctx context.Context, req *pb.RoomRequest) (*pb.Room, error) {
	start := time.Now()
	b.mu.Lock()
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()
	printReqLog(ctx, req)
	wkbox := b.searchAliveClient()
	var tmp pb.Room
	if _, err := wkbox.GetPara(&req.Key, &tmp); err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(500, err.Error())
	}
	(wkbox).Preserve(false)
	return &tmp, nil
}
