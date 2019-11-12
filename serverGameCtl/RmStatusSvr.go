package serverGameCtl

import (
	cm "RoomStatus/common"
	cf "RoomStatus/config"
	pb "RoomStatus/proto"
	rd "RoomStatus/store/redis"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"time"

	ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomStatusServer = (*RoomStatusBackend)(nil)

type RoomStatusBackend struct {
	// pb.RoomStatusServer
	mu *sync.Mutex
	// channels / workers
	deleteWk  *ants.PoolWithFunc
	createWk  *ants.PoolWithFunc
	getWk     *ants.PoolWithFunc
	getListWk *ants.PoolWithFunc
	updateWk  *ants.PoolWithFunc
	steamWk   *ants.PoolWithFunc

	redhdlr  []*rd.RdsCliBox
	sub      []*rd.RdsPubSub
	CoreKey  string
	Roomlist []*pb.Room
}

// Remark: the framework make consider "instant" request
//

// New : Create new backend
func New(conf *cf.ConfTmp) *RoomStatusBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)
	rdfl := []*rd.RdsCliBox{}
	for i := 0; i < conf.Database.WorkerNode; i++ {
		rdf := rd.New(ck, "wKU"+cm.HashText("num"+strconv.Itoa(i)))
		if _, err := rdf.Connect(conf); err == nil {
			rdfl = append(rdfl, rdf)
		}
	}

	g := RoomStatusBackend{
		CoreKey:   ck,
		mu:        &sync.Mutex{},
		redhdlr:   rdfl,
		deleteWk:  nil,
		createWk:  nil,
		updateWk:  nil,
		getWk:     nil,
		getListWk: nil,
	}
	// @refer: RSSvrCreate.go
	g.createWk, _ = ants.NewPoolWithFunc(
		conf.APIServer.MaxPoolSize/4,
		g.createWkTask)

	// @refer: RSSvrDelete.go
	g.deleteWk, _ = ants.NewPoolWithFunc(
		conf.APIServer.MaxPoolSize/4,
		g.deleteWkTask)

	// @refer: RSSvrGetInfo.go
	g.getWk, _ = ants.NewPoolWithFunc(
		conf.APIServer.MaxPoolSize/4,
		g.getInfoWkTask)

	// @refer: RSSvrGetList.go
	g.getListWk, _ = ants.NewPoolWithFunc(
		conf.APIServer.MaxPoolSize/4,
		g.getLsWkTask)
	g.updateWk, _ = ants.NewPoolWithFunc(
		conf.APIServer.MaxPoolSize/4,
		g.getLsWkTask)
	// @refer: RSSvrGetStream.go
	// g.steamWk, _ = ants.NewPoolWithFunc(
	// 	conf.APIServer.MaxPoolSize,
	// 	g.roomStreamWkTask)
	return &g
}

func (this *RoomStatusBackend) Shutdown() {
	log.Println("in shtdown proc")
	for _, v := range this.redhdlr {
		if _, err := v.CleanRem(); err != nil {
			log.Println(err)
		}
		if _, e := v.Disconn(); e != nil {
			log.Println(e)
		}
	}
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

func (b *RoomStatusBackend) searchAliveClient() *rd.RdsCliBox {
	for {
		wk := b.checkAliveClient()
		if wk == nil {
			// log.Println("busy at " + time.Now().String())
			time.Sleep(500)
		} else {
			wk.Preserve(true)
			return wk
		}
	}
}

// checkAliveClient
func (b *RoomStatusBackend) checkAliveClient() *rd.RdsCliBox {
	for _, v := range b.redhdlr {
		if !*v.IsRunning() {
			return v
		}
	}
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
