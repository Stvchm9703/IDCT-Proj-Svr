package serverctlNoRedisl

import (
	pb "RoomStatus/proto"
	"context"
	"errors"
	"log"
)

// UpdateRoom :
func (b *RoomStatusBackend) UpdateRoom(ctx context.Context, req *pb.CellStatus) (*pb.CellStatus, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	printReqLog(ctx, req)
	// wkbox := b.searchAliveClient()
	var room pb.Room
	// if _, err := (wkbox).GetPara(&req.Key, &room); err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	keynum := -1
	if len(room.CellStatus) == 9 && room.Round == 9 {
		log.Println("the game should be end")
		return nil, errors.New("GameEnd")
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
	// if _, err := wkbox.UpdatePara(&req.Key, &room); err != nil {
	// 	log.Fatalln(err)
	// 	return nil, err
	// }
	// (wkbox).Preserve(false)
	log.Println("b.RoomList", b.Roomlist)
	// log.Println("r", r)
	return room.CellStatus[len(room.CellStatus)], nil
}
