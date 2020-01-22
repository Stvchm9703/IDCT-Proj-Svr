package serverctlNoRedis

import (
	pb "RoomStatus/proto"
	"context"
	"errors"
	"time"

	types "github.com/gogo/protobuf/types"
)

func (b *RoomStatusBackend) QuitRoom(ctx context.Context, req *pb.RoomCreateReq) (*types.Empty, error) {
	var tmpRoom *RoomMgr
	delRoom := false
	for _, v := range b.Roomlist {
		if k := v.GetGS(req.UserId); k != nil {
			tmpRoom = v
			break
		}
	}
	if tmpRoom == nil {
		return nil, errors.New("NoRoomPlayerJoined")
	}
	tmpRoom.BroadCast("RoomSvrMgr",
		&pb.CellStatusResp{
			UserId:    "RoomSvrMgr",
			Key:       tmpRoom.Key,
			Timestamp: time.Now().String(),
			ResponseMsg: &pb.CellStatusResp_ErrorMsg{
				ErrorMsg: &pb.ErrorMsg{
					MsgInfo: "RoomWatcherQuit",
					MsgDesp: "RoomSvr:Watcher<" + req.UserId + "> is going to quit",
				}},
		})
	if tmpRoom.Room.HostId == req.UserId || tmpRoom.Room.DuelerId == req.UserId {
		tmpRoom.BroadCast("RoomSvrMgr",
			&pb.CellStatusResp{
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
	// tmpRoom.DelBS(req.UserId)
	tmpRoom.DelGS(req.UserId)
	if delRoom {
		_, err := b.DeleteRoom(ctx, &pb.RoomReq{
			Key: tmpRoom.Key,
		})
		return nil, err
	}

	return nil, nil
}
