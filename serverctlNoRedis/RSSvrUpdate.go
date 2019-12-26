package serverctlNoRedis

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
	var room pb.Room

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
	log.Println("b.RoomList", b.Roomlist)
	return room.CellStatus[len(room.CellStatus)], nil
}
