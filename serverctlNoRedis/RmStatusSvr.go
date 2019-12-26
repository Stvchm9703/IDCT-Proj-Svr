package serverctlNoRedis

import (
	// _ "RoomStatus"
	cm "RoomStatus/common"
	cf "RoomStatus/config"
	pb "RoomStatus/proto"
	rd "RoomStatus/store/redis"
	"context"
	"encoding/json"
	"log"
	"sync"
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomStatusServer = (*RoomStatusBackend)(nil)

type RoomStatusBackend struct {
	// pb.RoomStatusServer
	mu *sync.Mutex

	CoreKey  string
	Roomlist []*pb.Room
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
	// for _, v := range this.redhdlr {
	// 	if _, err := v.CleanRem(); err != nil {
	// 		log.Println(err)
	// 	}
	// 	if _, e := v.Disconn(); e != nil {
	// 		log.Println(e)
	// 	}
	// }
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

// func (b *RoomStatusBackend) searchAliveClient() *rd.RdsCliBox {
// 	for {
// 		wk := b.checkAliveClient()
// 		if wk == nil {
// 			// log.Println("busy at " + time.Now().String())
// 			time.Sleep(500)
// 		} else {
// 			wk.Preserve(true)
// 			return wk
// 		}
// 	}
// }

// checkAliveClient
func (b *RoomStatusBackend) checkAliveClient() *rd.RdsCliBox {
	// for _, v := range b.redhdlr {
	// 	if !*v.IsRunning() {
	// 		return v
	// 	}
	// }
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
