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

	types "github.com/gogo/protobuf/types"
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

	redhdlr  []*rd.RdsCliBox
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
		getWk:     nil,
		getListWk: nil,
	}
	g.createWk, _ = ants.NewPoolWithFunc(
		conf.APIServer.MaxPoolSize/4,
		g.createWkTask)

	g.deleteWk, _ = ants.NewPoolWithFunc(
		conf.APIServer.MaxPoolSize/4,
		g.deleteWkTask)

	g.getWk, _ = ants.NewPoolWithFunc(
		conf.APIServer.MaxPoolSize/4,
		g.getInfoWkTask)

	g.getListWk, _ = ants.NewPoolWithFunc(
		conf.APIServer.MaxPoolSize/4,
		g.getLsWkTask)

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
	In  interface{}
	Out chan interface{}
}

/// ===>>> Worker Goroutine function
// createWkTask:
func (b *RoomStatusBackend) createWkTask(payload interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()

	req, ok := payload.(WkTask).In.(*pb.RoomCreateRequest)
	if !ok {
		return
	}

	wkbox := b.searchAliveClient()

	// for loop it
	tmptime := time.Now().String() + req.HostId
	var f = ""

	for {
		f = cm.HashText(tmptime)
		l, err := (wkbox).ListRem(&f)
		if err != nil {
			log.Println(err)
			return
		}
		if len(*l) == 0 {
			break
		}
	}
	rmTmp := pb.Room{
		Key:        "Rm" + f,
		HostId:     req.HostId,
		DuelerId:   "",
		Status:     0,
		Round:      0,
		Cell:       -1,
		CellStatus: nil,
	}
	if _, err := wkbox.SetPara(&rmTmp.Key, rmTmp); err != nil {
		log.Println(err)
		return
	}
	wkbox.Preserve(false)

	payload.(WkTask).Out <- rmTmp
}

// TestCreateWkTask : Test Unit
func (b *RoomStatusBackend) TestCreateWkTask(pl interface{}) (rmTmp *pb.Room, err error) {
	if err := b.createWk.Invoke(pl.(WkTask)); err != nil {
		log.Println("err in create Wk", err)
		return nil, err
	}
	// ====== Worker End =======
	plc := <-(pl.(WkTask)).Out
	rmTmpa := plc.(pb.Room)
	rmTmp = &rmTmpa
	// create room success
	b.Roomlist = append(b.Roomlist, rmTmp)
	return
}

func (b *RoomStatusBackend) deleteWkTask(payload interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()

	req, ok := payload.(WkTask).In.(*pb.RoomRequest)
	if !ok {
		return
	}

	wkbox := b.searchAliveClient()

	if _, err := (wkbox).RemovePara(&req.Key); err != nil {
		log.Fatalln(err)
		return
	}
	(wkbox).Preserve(false)
	payload.(WkTask).Out <- &req.Key
}

// TestCreateWkTask : Test Unit
func (b *RoomStatusBackend) TestDeteleWkTask(pl interface{}) (rmTmp *pb.Room, err error) {
	if err := b.deleteWk.Invoke(pl.(WkTask)); err != nil {
		log.Println("err in create Wk", err)
		return nil, err
	}
	// ====== Worker End =======
	plc := <-(pl.(WkTask)).Out
	for k, v := range b.Roomlist {
		if v.Key == *plc.(*string) {
			rmTmp = b.Roomlist[k]
			b.Roomlist = append(b.Roomlist[:k], b.Roomlist[k+1:]...)
		}
	}
	return
}

func (b *RoomStatusBackend) getInfoWkTask(payload interface{}) {
	req, ok := payload.(WkTask).In.(*pb.RoomRequest)
	if !ok {
		return
	}

	wkbox := b.searchAliveClient()
	var tmp pb.Room
	if _, err := wkbox.GetPara(&req.Key, &tmp); err != nil {
		log.Fatalln(err)
		// return nil, err
	}
	(wkbox).Preserve(false)
	payload.(WkTask).Out <- &tmp
}

func (b *RoomStatusBackend) TestGetInfoWkTask(pl interface{}) (rmTmp *pb.Room, err error) {
	if err = b.getWk.Invoke(pl.(WkTask)); err != nil {
		log.Println("err in create Wk", err)
		return
	}
	// ====== Worker End =======
	rm := <-(pl.(WkTask)).Out
	rmTmp = rm.(*pb.Room)
	return
}

func (b *RoomStatusBackend) getLsWkTask(payload interface{}) {
	req, ok := payload.(WkTask).In.(*pb.RoomListRequest)
	if !ok {
		return
	}

	wkbox := b.searchAliveClient()
	// var tmp pb.Room
	var RmList []*pb.Room
	strl, err2 := wkbox.GetParaList(&req.Requirement)
	if err2 != nil {
		log.Fatalln(err2)
	}
	// log.Println("strl:", string(*strl))
	err := json.Unmarshal(*strl, &RmList)
	if err != nil {
		log.Fatalln(err)
	}
	// log.Println(RmList)
	(wkbox).Preserve(false)
	payload.(WkTask).Out <- RmList
}

func (b *RoomStatusBackend) TestGetLsWkTask(pl interface{}) (rmTmp *pb.RoomListResponse, err error) {
	if err = b.getListWk.Invoke(pl.(WkTask)); err != nil {
		log.Println("err in create Wk", err)
		return
	}
	// ====== Worker End =======
	rm := <-(pl.(WkTask)).Out
	fg := rm.([]*pb.Room)
	rmTmp = &pb.RoomListResponse{
		Result: fg,
	}
	return
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

// CreateRoom :
func (b *RoomStatusBackend) CreateRoom(ctx context.Context, req *pb.RoomCreateRequest) (*pb.Room, error) {
	printReqLog(ctx, req)

	// var k chan pb.Room
	// ====== Worker Start =======
	pl := &WkTask{In: req, Out: make(chan interface{})}
	if err := b.createWk.Invoke(pl); err != nil {
		log.Println(err)
		return nil, err
	}
	// ====== Worker End =======
	rmTmp := (<-pl.Out).(pb.Room)
	// create room success
	b.Roomlist = append(b.Roomlist, &rmTmp)
	return &rmTmp, nil
}

// GetRoomList :
func (b *RoomStatusBackend) GetRoomList(ctx context.Context, req *pb.RoomListRequest) (res *pb.RoomListResponse, err error) {
	printReqLog(ctx, req)
	// b.mu.Lock()
	// defer b.mu.Unlock()

	// search free box
	wkbox := b.checkAliveClient()
	if wkbox == nil {
		// busy
		log.Println("busy at " + time.Now().String())
		return nil, nil
		// TODO : use chan to push task ?
	}
	// var tmp pb.Room
	// var RmList []pb.Room

	// if _, err2 := wkbox.GetParaList(&req.Requirement, &RmList, tmp); err2 != nil {
	// 	log.Fatalln(err)
	// 	err = err2
	// }
	// for _, v := range RmList {
	// 	res.Result = append(res.Result, &v)
	// }
	err = nil
	return res, err
}

// GetRoomCurrentInfo :
func (b *RoomStatusBackend) GetRoomCurrentInfo(ctx context.Context, req *pb.RoomRequest) (*pb.Room, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method GetRoomCurrentInfo not implemented")
	printReqLog(ctx, req)
	pl := &WkTask{
		In:  req,
		Out: make(chan interface{})}
	if err := b.getWk.Invoke(pl); err != nil {
		log.Println("err in create Wk", err)
		return nil, err
	}
	// ====== Worker End =======
	rm := <-(pl).Out
	tmp := rm.(*pb.Room)
	return tmp, nil
}

// DeleteRoom :
func (b *RoomStatusBackend) DeleteRoom(ctx context.Context, req *pb.RoomRequest) (*types.Empty, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	printReqLog(ctx, req)
	// var k chan pb.Room
	// ====== Worker Start =======
	pl := &WkTask{In: req, Out: make(chan interface{})}
	if err := b.deleteWk.Invoke(pl); err != nil {
		log.Println(err)
		return nil, err
	}
	// ====== Worker End =======
	plc := <-(pl).Out
	for k, v := range b.Roomlist {
		if v.Key == *plc.(*string) {
			// rmTmp = b.Roomlist[k]
			log.Println(b.Roomlist[k])
			b.Roomlist = append(b.Roomlist[:k], b.Roomlist[k+1:]...)
		}
	}
	log.Println("b.RoomList", b.Roomlist)
	return nil, nil
}
