package serverctlNoRedis

import (
	"RoomStatus/common"
	pb "RoomStatus/proto"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

// UpdateRoom :
func (b *RoomStatusBackend) UpdateRoom(ctx context.Context, req *pb.CellStatusReq) (*pb.CellStatusResp, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	common.PrintReqLog(ctx, "update-room", req)
	var rmg *pb.Room
	for k := range b.Roomlist {
		if (*b.Roomlist[k]).Key == req.Key {
			rmg = b.Roomlist[k]
		}
	}
	if (rmg) == nil {
		log.Println("RoomNotExistInUpdate")
		return nil, errors.New("RoomNotExistInUpdate")
	}

	// remark!!
	// -1 is initial msg (testing)

	reqRoom := req.GetCellStatus()
	if reqRoom == nil {
		log.Println("UnknownCellStatus")
		return nil, errors.New("UnknownCellStatus")
	}
	if reqRoom.CellNum == -2 {
		log.Println("Player Give Up")
		msgp := &pb.CellStatusResp{
			UserId:    req.UserId,
			Key:       (*rmg).Key,
			Timestamp: time.Now().String(),
			Status:    200,
			ResponseMsg: &pb.CellStatusResp_ErrorMsg{
				ErrorMsg: &pb.ErrorMsg{
					MsgInfo: "PlayerGiveUp",
					MsgDesp: fmt.Sprintf("Player %s GiveUp", req.UserId),
				},
			},
		}
		go b.BroadCast(msgp)
		return msgp, nil
	}

	if len((*rmg).CellStatus) == 10 {
		log.Println("the game should be end")
		return nil, errors.New("GameEnd")
	}

	fmt.Println(len(rmg.CellStatus))
	keynum := len((*rmg).CellStatus)
	if keynum > 0 {
		cs := (*rmg).CellStatus[keynum-1]
		log.Println(cs)
		if cs.Turn == reqRoom.Turn {
			log.Println("GameRuleNotPlyrTurn")
			return nil, errors.New("GameRuleNotPlyrTurn")
		}
		for _, v := range (*rmg).CellStatus {
			if v.CellNum == reqRoom.CellNum {
				log.Println("GameRuleCellOcc")
				return nil, errors.New("GameRuleCellOcc")
			}
		}
	}

	(*rmg).CellStatus = append((*rmg).CellStatus, req.GetCellStatus())
	(*rmg).Cell = int32(len((*rmg).CellStatus))
	(*rmg).Round++

	log.Println(rmg.CellStatus)

	log.Println("b.RoomList", b.Roomlist)

	// send update BroadCast
	msgp := &pb.CellStatusResp{
		UserId:    req.UserId,
		Key:       (*rmg).Key,
		Timestamp: time.Now().String(),
		Status:    200,
		ResponseMsg: &pb.CellStatusResp_CellStatus{
			CellStatus: reqRoom,
		},
	}

	// !Broadcast
	go b.BroadCast(msgp)
	return msgp, nil
}
