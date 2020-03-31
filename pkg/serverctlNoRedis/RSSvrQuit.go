package serverctlNoRedis

import (
	"RoomStatus/pkg/common"
	pb "RoomStatus/proto"
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	types "github.com/gogo/protobuf/types"
)

func (b *RoomStatusBackend) QuitRoom(ctx context.Context, req *pb.RoomCreateReq) (*types.Empty, error) {
	common.PrintReqLog(ctx, "quit-room", req)
	var tmpRoom *pb.Room
	delRoom := false
	wg := sync.WaitGroup{}
	for k := range b.Roomlist {
		if (*b.Roomlist[k]).HostId == req.UserId || b.Roomlist[k].DuelerId == req.UserId {
			tmpRoom = b.Roomlist[k]
		}
	}
	if tmpRoom == nil {
		fmt.Println("NoRoomPlayerJoined")
		return nil, errors.New("NoRoomPlayerJoined")
	}
	// !Broadcast
	wg.Add(1)
	fmt.Println("broadcast RoomWatcherQuit")
	go func() {
		b.BroadCast(&pb.CellStatusResp{
			UserId:    "RoomSvrMgr",
			Key:       tmpRoom.Key,
			Timestamp: time.Now().String(),
			ResponseMsg: &pb.CellStatusResp_ErrorMsg{
				ErrorMsg: &pb.ErrorMsg{
					MsgInfo: "RoomWatcherQuit",
					MsgDesp: "RoomSvr:Watcher<" + req.UserId + "> is going to quit",
				}},
		})
		wg.Done()
	}()
	fmt.Printf("\nSomeone quit %s \n", req.UserId)
	if tmpRoom.HostId == req.UserId || tmpRoom.DuelerId == req.UserId {
		// !Broadcast
		wg.Add(1)
		log.Println("Player quit")

		go func() {
			b.BroadCast(&pb.CellStatusResp{
				UserId:    "RoomSvrMgr",
				Key:       tmpRoom.Key,
				Timestamp: time.Now().String(),
				ResponseMsg: &pb.CellStatusResp_ErrorMsg{
					ErrorMsg: &pb.ErrorMsg{
						MsgInfo: "RoomPlayerQuit",
						MsgDesp: fmt.Sprintf("RoomSvr:Player<%s> is going to quit, this Room may close connect", req.UserId),
					}},
			})
			wg.Done()
		}()
		delRoom = true
	}
	wg.Wait()
	// tmpRoom.DelGS(req.UserId)
	if delRoom {
		err := b.RemoveRoom(&pb.RoomReq{
			Key: tmpRoom.Key,
		})
		return &types.Empty{}, err
	}
	return &types.Empty{}, nil
}
