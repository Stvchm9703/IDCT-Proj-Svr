package serverctlNoRedis

import (
	cm "RoomStatus/common"
	pb "RoomStatus/proto"
	"context"
	"errors"
	"time"
)

// CreateRoom :
func (b *RoomStatusBackend) CreateRoom(ctx context.Context, req *pb.RoomCreateRequest) (*pb.Room, error) {
	printReqLog(ctx, req)
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, vr := range b.Roomlist {
		if vr.HostId == req.HostId || vr.DuelerId == req.HostId {
			return nil, errors.New("GameSetExistWithCurrentPlayer")
		}
	}

	// wkbox := b.searchAliveClient()
	// for loop it
	tmptime := time.Now().String() + req.HostId
	var f = ""
	for {
		f = cm.HashText(tmptime)
		// ------
		var l []*string
		for _, v := range b.Roomlist {
			if v.Key == f {
				l = append(l, &v.Key)
			}
		}
		if len(l) == 0 {
			break
		}
		// -----
	}
	rmTmp := pb.Room{
		Key:        "Rm" + f,
		HostId:     req.HostId,
		DuelerId:   "",
		Status:     0,
		Round:      0,
		Cell:       -1,
		CellStatus: nil,
	}
	b.Roomlist = append(b.Roomlist, &rmTmp)
	return &rmTmp, nil
}
