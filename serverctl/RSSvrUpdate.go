package serverctl

import (
	pb "RoomStatus/proto"
	"context"
	"log"
	"time"

	"github.com/gogo/status"
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

// // TestUpdateWkTask : Test Unit
// func (b *RoomStatusBackend) TestUpdateWkTask(pl interface{}) (rmTmp *pb.Room, err error) {
// 	if err := b.UpdateWk.Invoke(pl.(WkTask)); err != nil {
// 		log.Println("err in create Wk", err)
// 		return nil, err
// 	}
// 	// ====== Worker End =======
// 	plc := <-(pl.(WkTask)).Out
// 	log.Println("plc;", plc)
// 	rmTmp = plc.(*pb.Room)
// 	return
// }

// pb.RoomStatus.

// UpdateRoom :
func (b *RoomStatusBackend) UpdateRoom(ctx context.Context, req *pb.CellStatus) (*pb.CellStatus, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	printReqLog(ctx, req)
	start := time.Now()
	b.mu.Lock()
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()
	wkbox := b.searchAliveClient()
	var room pb.Room
	if _, err := (wkbox).GetPara(&req.Key, &room); err != nil {
		log.Println(err)
		return nil, status.Errorf(500, err.Error())
	}
	keynum := -1
	if len(room.CellStatus) == 9 && room.Round == 9 {
		log.Println("the game should be end")
		return nil, status.Errorf(500, "GameEnd")
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
			break
		}
	}
	if _, err := wkbox.UpdatePara(&room.Key, &room); err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(500, err.Error())
	}
	(wkbox).Preserve(false)
	log.Println("b.RoomList", b.Roomlist)
	// log.Println("r", r)
	return room.CellStatus[len(room.CellStatus)], nil
}
