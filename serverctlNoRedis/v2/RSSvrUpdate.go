package v2

import (
	pb "RoomStatus/proto/v2"
	"context"
	"errors"
	"log"
	"time"
)

// UpdateRoom :
func (b *RoomStatusBackend) UpdateRoom(ctx context.Context, req *pb.CellStatusReq) (*pb.CellStatusResp, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	printReqLog(ctx, req)
	var room *pb.Room

	for _, v := range b.Roomlist {
		if v.Key == req.Key {
			room = &v.Room
		}
	}
	if room == nil {
		return nil, errors.New("RoomNotExistInUpdate")
	}

	keynum := -1
	if len(room.CellStatus) == 9 && room.Round == 9 {
		log.Println("the game should be end")
		return nil, errors.New("GameEnd")
	}

	reqRoom := req.GetCellStatus()
	if reqRoom == nil {
		return nil, errors.New("UnknownCellStatus")
	}

	for k, v := range room.CellStatus {
		if v.Turn == reqRoom.Turn {
			room.Cell = int32(k)
			v.CellNum = reqRoom.CellNum
			keynum = k
			break
		}
	}

	if keynum == -1 {
		room.CellStatus = append(room.CellStatus, req.GetCellStatus())
		room.Cell = int32(len(room.CellStatus))
		room.Round++
	}

	log.Println("b.RoomList", b.Roomlist)
	return &pb.CellStatusResp{
		UserId:    req.UserId,
		Key:       room.Key,
		Timestamp: time.Now().String(),
		Status:    200,
		ResponseMsg: &pb.CellStatusResp_CellStatus{
			CellStatus: reqRoom,
		},
	}, nil
}
