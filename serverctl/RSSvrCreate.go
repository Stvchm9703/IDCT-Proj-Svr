package serverctl

import (
	cm "RoomStatus/common"
	pb "RoomStatus/proto"
	"context"
	"errors"
	"log"
	"time"
)

/// ===>>> Worker Goroutine function
// createWkTask:
func (b *RoomStatusBackend) createWkTask(payload interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()

	req, ok := payload.(WkTask).In.(*pb.RoomCreateRequest)
	if !ok {
		return
	}

	wkbox := b.searchAliveClient()

	// for loop it
	tmptime := time.Now().String() + req.HostId
	var f = ""

	for {
		f = cm.HashText(tmptime)
		l, err := (wkbox).ListRem(&f)
		if err != nil {
			log.Println(err)
			return
		}
		if len(*l) == 0 {
			break
		}
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
	if _, err := wkbox.SetPara(&rmTmp.Key, rmTmp); err != nil {
		log.Println(err)
		return
	}
	wkbox.Preserve(false)

	payload.(WkTask).Out <- rmTmp
	return
}

// TestCreateWkTask : Test Unit
// func (b *RoomStatusBackend) TestCreateWkTask(pl interface{}) (rmTmp *pb.Room, err error) {
// 	if err := b.createWk.Invoke(pl.(WkTask)); err != nil {
// 		log.Println("err in create Wk", err)
// 		return nil, err
// 	}
// 	// ====== Worker End =======
// 	plc := <-(pl.(WkTask)).Out
// 	rmTmpa := plc.(pb.Room)
// 	rmTmp = &rmTmpa
// 	// create room success
// 	b.Roomlist = append(b.Roomlist, rmTmp)
// 	return
// }

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

	wkbox := b.searchAliveClient()
	// for loop it
	tmptime := time.Now().String() + req.HostId
	var f = ""
	for {
		f = cm.HashText(tmptime)
		l, err := (wkbox).ListRem(&f)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if len(*l) == 0 {
			break
		}
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
	if _, err := wkbox.SetPara(&rmTmp.Key, rmTmp); err != nil {
		log.Println(err)
		return nil, err
	}
	wkbox.Preserve(false)
	b.Roomlist = append(b.Roomlist, &rmTmp)
	return &rmTmp, nil
}
