package serverctlNoRedis

import (
	"RoomStatus/common"
	pb "RoomStatus/proto"
	"context"
	"errors"
	"time"

	types "github.com/gogo/protobuf/types"
)

func (b *RoomStatusBackend) QuitRoom(ctx context.Context, req *pb.RoomCreateReq) (*types.Empty, error) {
	common.PrintReqLog(ctx, "quit-room", req)
	var tmpRoom *pb.Room
	delRoom := false
	for k := range b.Roomlist {
		if (*b.Roomlist[k]).HostId == req.UserId || b.Roomlist[k].DuelerId == req.UserId {
			tmpRoom = b.Roomlist[k]
		}
	}
	if tmpRoom == nil {
		return nil, errors.New("NoRoomPlayerJoined")
	}
	// !Broadcast
	go b.BroadCast(&pb.CellStatusResp{
		UserId:    "RoomSvrMgr",
		Key:       tmpRoom.Key,
		Timestamp: time.Now().String(),
		ResponseMsg: &pb.CellStatusResp_ErrorMsg{
			ErrorMsg: &pb.ErrorMsg{
				MsgInfo: "RoomWatcherQuit",
				MsgDesp: "RoomSvr:Watcher<" + req.UserId + "> is going to quit",
			}},
	})
	if tmpRoom.HostId == req.UserId || tmpRoom.DuelerId == req.UserId {
		// !Broadcast
		go b.BroadCast(&pb.CellStatusResp{
			UserId:    "RoomSvrMgr",
			Key:       tmpRoom.Key,
			Timestamp: time.Now().String(),
			ResponseMsg: &pb.CellStatusResp_ErrorMsg{
				ErrorMsg: &pb.ErrorMsg{
					MsgInfo: "RoomHostQuit",
					MsgDesp: "RoomSvr:Host Player<" + req.UserId + "> is going to quit, this Room may close connect",
				}},
		})
		delRoom = true
	}
	// tmpRoom.DelGS(req.UserId)
	if delRoom {
		_, err := b.DeleteRoom(ctx, &pb.RoomReq{
			Key: tmpRoom.Key,
		})
		return nil, err
	}

	return &types.Empty{}, nil
}
