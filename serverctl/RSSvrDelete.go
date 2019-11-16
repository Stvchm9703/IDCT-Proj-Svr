package serverctl

import (
	pb "RoomStatus/proto"
	"context"
	"log"
	"time"

	types "github.com/gogo/protobuf/types"
	"github.com/gogo/status"
)

// DeleteRoom :
func (b *RoomStatusBackend) DeleteRoom(ctx context.Context, req *pb.RoomRequest) (*types.Empty, error) {
	start := time.Now()
	b.mu.Lock()
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()

	wkbox := b.searchAliveClient()

	if _, err := (wkbox).RemovePara(&req.Key); err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(500, err.Error())
	}
	(wkbox).Preserve(false)
	done := false
	for k, v := range b.Roomlist {
		if v.Key == req.Key {
			// rmTmp = b.Roomlist[k]
			log.Println(b.Roomlist[k])
			b.Roomlist = append(b.Roomlist[:k], b.Roomlist[k+1:]...)
			done = true
		}
	}
	if !done {
		return nil, status.Errorf(500, "RoomNotExist")
	}
	log.Println("b.RoomList", b.Roomlist)
	return nil, nil
}

//
