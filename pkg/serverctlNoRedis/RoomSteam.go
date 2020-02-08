package serverctlNoRedis

import (
	cm "RoomStatus/common"
	pb "RoomStatus/proto"
	"errors"
	"fmt"
	"log"
	"time"
)

func (b *RoomStatusBackend) GetRoomStream(csq *pb.CellStatusReq, rs pb.RoomStatus_GetRoomStreamServer) error {
	cm.PrintReqLog(rs.Context(), csq)
	no_rm_flg := true
	var vf *RoomMgr

	for _, v := range b.Roomlist {
		if v.Key == csq.Key {
			if v.GetGS(csq.UserId) != nil {
				return errors.New("GetStreamIsExist")
			}
			v.AddGS(csq.UserId, &rs)
			log.Println("Add GS Stream")
			no_rm_flg = false
			vf = v
			break
		}
	}
	if no_rm_flg {
		return errors.New("NoRoomExist")
	}

	go func() {
		<-rs.Context().Done()
		log.Println("close done")
		_, err := vf.DelGS(csq.UserId)
		if err != nil {
			log.Println(err)
		}
		vf.BroadCast(csq.UserId,
			&pb.CellStatusResp{
				UserId:    csq.UserId,
				Key:       vf.Key,
				Status:    201,
				Timestamp: time.Now().String(),
				ResponseMsg: &pb.CellStatusResp_ErrorMsg{
					ErrorMsg: &pb.ErrorMsg{MsgInfo: "ConnEnd", MsgDesp: fmt.Sprintf("User<%v> End to Room<%v>", csq.UserId, vf.Key)},
				},
			})
	}()
	for {
	}
}

// RoomStream : Skipp the service
func (b *RoomStatusBackend) RoomStream(stream pb.RoomStatus_RoomStreamServer) error {
	return nil
}
