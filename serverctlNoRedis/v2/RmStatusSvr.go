package v2

import (
	// _ "RoomStatus"
	cm "RoomStatus/common"
	cf "RoomStatus/config"
	pb "RoomStatus/proto/v2"
	rd "RoomStatus/store/redis"
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomStatusServer = (*RoomStatusBackend)(nil)

func (rm *RoomMgr) GetGS(user_id string) pb.RoomStatus_GetRoomStreamServer {
	stream, ok := (*rm).conn_pool.Load(user_id)
	if !ok {
		return nil
	}
	return (stream).(pb.RoomStatus_GetRoomStreamServer)
}

func (rm *RoomMgr) AddGS(user_id string, stream pb.RoomStatus_GetRoomStreamServer) {
	rm.get_only_stream.Store(user_id, stream)
	_, cn := context.WithCancel(stream.Context())
	rm.close_link.Store(user_id, cn)
}

func (rm *RoomMgr) DelGS(user_id string) {
	if d, ok := rm.close_link.Load(user_id); ok {
		d.(context.CancelFunc)()
	}
	rm.get_only_stream.Delete(user_id)
}
func (rm *RoomMgr) BroadCastGS(from string, message *pb.CellStatusResp) {
	rm.conn_pool.Range(func(username_i, stream_i interface{}) bool {
		username := username_i.(string)
		stream := (stream_i).(pb.RoomStatus_GetRoomStreamServer)
		if username == from {
			return true
		} else {
			(stream).Send(message)
		}
		return true
	})
}

func (rm *RoomMgr) GetBS(user_id string) pb.RoomStatus_RoomStreamServer {
	stream, ok := (*rm).conn_pool.Load(user_id)
	if !ok {
		return nil
	}
	return (stream).(pb.RoomStatus_RoomStreamServer)
}

func (rm *RoomMgr) AddBStream(user_id string, stream pb.RoomStatus_RoomStreamServer) {
	rm.conn_pool.Store(user_id, stream)
	_, cn := context.WithCancel((stream).Context())
	rm.close_link.Store(user_id, cn)
}

func (rm *RoomMgr) DelBS(user_id string) {
	if d, ok := rm.close_link.Load(user_id); ok {
		d.(context.CancelFunc)()
	}
	rm.conn_pool.Delete(user_id)
	// rm.close_link.Delete(user_id)
}
func (rm *RoomMgr) ClearAll() {
	log.Println("ClearAll Proc")
	rm.close_link.Range(func(key interface{}, value interface{}) bool {
		// value.(context.CancelFunc)()
		user_id := key.(string)
		rm.DelBS(user_id)
		rm.DelGS(user_id)
		return true
	})

}
func (rm *RoomMgr) BroadCastBS(from string, message *pb.CellStatusResp) {
	rm.conn_pool.Range(func(username_i, stream_i interface{}) bool {
		username := username_i.(string)
		stream := (stream_i).(pb.RoomStatus_RoomStreamServer)
		if username == from {
			return true
		} else {
			(stream).Send(message)
		}
		return true
	})
}

func (rm *RoomMgr) BroadCast(from string, message *pb.CellStatusResp) {
	log.Println("BS!", message)
	rm.BroadCastBS(from, message)
	rm.BroadCastGS(from, message)
}

type RoomMgr struct {
	pb.Room
	conn_pool       sync.Map
	get_only_stream sync.Map
	close_link      sync.Map
}

type RoomStatusBackend struct {
	// pb.RoomStatusServer
	mu       *sync.Mutex
	CoreKey  string
	Roomlist []*RoomMgr
}

// Remark: the framework make consider "instant" request
//

// New : Create new backend
func New(conf *cf.ConfTmp) *RoomStatusBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)

	g := RoomStatusBackend{
		CoreKey: ck,
		mu:      &sync.Mutex{},
	}

	return &g
}

func (this *RoomStatusBackend) Shutdown() {
	log.Println("in shtdown proc")
	/// TODO: send closing msg to all client
	for _, v := range this.Roomlist {
		log.Println("Server OS.sigKill")
		v.BroadCast("RmSvrMgr",
			&pb.CellStatusResp{
				UserId:    "RmSvrMgr",
				Key:       v.Key,
				Status:    201,
				Timestamp: time.Now().String(),
				ResponseMsg: &pb.CellStatusResp_ErrorMsg{
					ErrorMsg: &pb.ErrorMsg{
						MsgInfo: "ConnEnd",
						MsgDesp: "Server OS.sigKill",
					},
				},
			})
		v.ClearAll()
	}
	log.Println("endof shutdown proc:", this.CoreKey)

}

// 	Impletement from GameCtl.pb.go(auto-gen file)
// 		CreateCred(req *pb.CreateCredReq, srv pb.RoomStatus_CreateCredServer) error
// 		CreateRoom(context.Context, *types.Empty) (*Room, error)
// 		GetRoomList(context.Context, *RoomListRequest) (*RoomListResponse, error)
// 		GetRoomCurrentInfo(context.Context, *RoomRequest) (*Room, error)
// 		GetRoomStream(*RoomRequest, RoomStatus_GetRoomStreamServer) error
// 		UpdateRoomStatus(context.Context, *CellStatus) (*types.Empty, error)
// 		DeleteRoom(context.Context, *RoomRequest) (*types.Empty, error)

type WkTask struct {
	In     interface{}
	Out    chan interface{}
	Stream interface{}
}

func (b *RoomStatusBackend) checkAliveClient() *rd.RdsCliBox {
	return nil
}

/// <<<=== Worker Goroutine function

// printReqLog
func printReqLog(ctx context.Context, req interface{}) {
	jsoon, _ := json.Marshal(ctx)
	log.Println(string(jsoon))

	jsoon, _ = json.Marshal(req)
	log.Println(string(jsoon))
}
