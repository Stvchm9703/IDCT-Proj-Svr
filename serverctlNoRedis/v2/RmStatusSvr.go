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
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomStatusServer = (*RoomStatusBackend)(nil)

type ConnPool struct{ sync.Map }

func (p *ConnPool) Get(user_id string) pb.RoomStatus_RoomStreamServer {
	if stream, ok := p.Load(user_id); ok {
		return stream.(pb.RoomStatus_RoomStreamServer)
	} else {
		return nil
	}
}
func (p *ConnPool) Add(user_id string, stream pb.RoomStatus_RoomStreamServer) {
	p.Store(user_id, stream)
}
func (p *ConnPool) Del(user_id string) {
	p.Delete(user_id)
}

// BroadCast handler
func (p *ConnPool) BroadCast(from string, message *pb.CellStatusResp) {
	log.Printf("BroadCast from: %s, message: %s\n", from, message.UserId)
	p.Range(func(username_i, stream_i interface{}) bool {
		username := username_i.(string)
		stream := stream_i.(pb.RoomStatus_RoomStreamServer)
		if username == from {
			return true
		} else {
			stream.Send(message)
		}
		return true
	})
}

type RoomMgr struct {
	pb.Room
	conn_pool *ConnPool
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
	log.Println("endof shtdown proc:", this.CoreKey)

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
