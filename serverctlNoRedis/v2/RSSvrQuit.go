package v2

import (
	pb "RoomStatus/proto/v2"
	"context"
	"errors"

	types "github.com/gogo/protobuf/types"
)

func (b *RoomStatusBackend) QuitRoom(ctx context.Context, req *pb.RoomCreateReq) (*types.Empty, error) {
	var tmpRoom *RoomMgr
	for _, v := range b.Roomlist {
		if v.Room.HostId == req.UserId {
			v.conn_pool.BroadCast("RoomSvrMgr",
				&pb.CellStatusResp{
					ResponseMsg: &pb.CellStatusResp_ErrorMsg{
						ErrorMsg: &pb.ErrorMsg{
							MsgInfo: "RoomHostQuit",
							MsgDesp: "RoomSvr:Host Player<" + req.UserId + "> is going to quit, this Room may close connect",
						}},
				})
			tmpRoom = v
			break

		} else if v.Room.DuelerId == req.UserId {
			v.conn_pool.BroadCast("RoomSvrMgr",
				&pb.CellStatusResp{
					ResponseMsg: &pb.CellStatusResp_ErrorMsg{
						ErrorMsg: &pb.ErrorMsg{
							MsgInfo: "RoomDuelQuit",
							MsgDesp: "RoomSvr:Duel Player<" + req.UserId + "> is going to quit, this Room may ready for open",
						}},
				})
			tmpRoom = v
			break
		}

		if k := v.conn_pool.Get(req.UserId); k != nil {
			v.conn_pool.BroadCast("RoomSvrMgr",
				&pb.CellStatusResp{
					ResponseMsg: &pb.CellStatusResp_ErrorMsg{
						ErrorMsg: &pb.ErrorMsg{
							MsgInfo: "RoomWatcherQuit",
							MsgDesp: "RoomSvr:Watcher<" + req.UserId + "> is going to quit",
						}},
				})
			tmpRoom = v
			break
		}
	}

	if tmpRoom == nil {
		return nil, errors.New("NoRoomPlayerJoined")
	}
	tmpRoom.conn_pool.Del(req.UserId)
	return nil, nil
}
