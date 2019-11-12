package serverGameCtl

import (
	pb "RoomStatus/proto"
	"context"
	"log"
)

func (b *RoomStatusBackend) updateWkTask(payload interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()

	req, ok := payload.(WkTask).In.(*pb.CellStatus)
	if !ok {
		return
	}

	wkbox := b.searchAliveClient()
	var room pb.Room
	if _, err := (wkbox).GetPara(&req.Key, &room); err != nil {
		log.Fatalln(err)
		return
	}
	keynum := -1
	if len(room.CellStatus) == 9 && room.Round == 9 {
		log.Println("the game should be end")
		payload.(WkTask).Out <- nil
		return
	}

	for k, v := range room.CellStatus {
		if v.Turn == req.Turn {
			room.Cell = int32(k)
			v.CellNum = req.CellNum
			keynum = k
			break
		}
	}

	if keynum == -1 {
		room.CellStatus = append(room.CellStatus, req)
		room.Cell = int32(len(room.CellStatus))
		room.Round++
	}

	for _, v := range b.Roomlist {
		if v.Key == room.Key {
			v = &room
		}
	}
	if _, err := wkbox.UpdatePara(&req.Key, &room); err != nil {
		log.Fatalln(err)
		return
	}
	(wkbox).Preserve(false)
	payload.(WkTask).Out <- &room
	return
}

// TestUpdateWkTask : Test Unit
func (b *RoomStatusBackend) TestUpdateWkTask(pl interface{}) (rmTmp *pb.Room, err error) {
	if err := b.deleteWk.Invoke(pl.(WkTask)); err != nil {
		log.Println("err in create Wk", err)
		return nil, err
	}
	// ====== Worker End =======
	plc := <-(pl.(WkTask)).Out
	log.Println("plc;", plc)
	rmTmp = plc.(*pb.Room)
	return
}

// pb.RoomStatus.

// UpdateRoom :
func (b *RoomStatusBackend) UpdateRoom(ctx context.Context, req *pb.CellStatus) (*pb.CellStatus, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	printReqLog(ctx, req)
	// var k chan pb.Room
	// ====== Worker Start =======
	pl := &WkTask{In: req, Out: make(chan interface{})}
	if err := b.deleteWk.Invoke(pl); err != nil {
		log.Println(err)
		return nil, err
	}
	// ====== Worker End =======
	plc := <-(pl).Out
	r := plc.(*pb.Room).CellStatus
	qq := r[len(r)-1]
	log.Println("b.RoomList", b.Roomlist)
	log.Println("r", r)
	return qq, nil
}
