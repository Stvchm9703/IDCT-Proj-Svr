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
)

type Backend struct {
	pb.UnimplementedRoomStatusServer
	mu *sync.Mutex
	// channel
	redChannel []chan *rd.RdsCliBox
	Roomlist   []*pb.Room
	redhdlr    []*rd.RdsCliBox
	CoreKey    string
}

// https://www.cnblogs.com/jkko123/p/6885389.html

// https://gobyexample.com/worker-pools
// https://www.atatus.com/blog/goroutines-error-handling/
// https://michaelchen.tech/golang-programming/concurrency/
// https://mgleon08.github.io/blog/2018/05/17/golang-goroutine-channel-worker-pool-select-mutex/
// https://eli.thegreenplace.net/2019/on-concurrency-in-go-http-servers/

// Remark: the framework make consider "instant" request
//

// New : Create new backend
func New(conf *cf.ConfTmp) *Backend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)
	rdfl := []*rd.RdsCliBox{}
	for i := 0; i < conf.Database.WorkerNode; i++ {
		rdf := rd.RdsCliBox{
			CoreKey: ck,
			Key:     "wKU" + cm.HashText("num"+strconv.Itoa(i)),
		}
		if _, err := rdf.Connect(conf); err == nil {
			rdfl = append(rdfl, &rdf)
		}
	}

	return &Backend{
		mu:      &sync.Mutex{},
		redhdlr: rdfl,
	}
}

// Start : Start the server loop logic
func (b *Backend) Start() error {
	for {
		for _, v := range b.Roomlist {
			log.Println("room:", v)

		}
	}
}

// 	Impletement from GameCtl.pb.go(auto-gen file)
// 		CreateRoom(context.Context, *types.Empty) (*Room, error)
// 		GetRoomList(context.Context, *RoomListRequest) (*RoomListResponse, error)
// 		GetRoomCurrentInfo(context.Context, *RoomRequest) (*Room, error)
// 		GetRoomStream(*RoomRequest, RoomStatus_GetRoomStreamServer) error
// 		UpdateRoomStatus(context.Context, *CellStatus) (*types.Empty, error)
// 		DeleteRoom(context.Context, *RoomRequest) (*types.Empty, error)

// checkAliveClient
func (b *Backend) checkAliveClient() *rd.RdsCliBox {
	for _, v := range b.redhdlr {
		if !v.IsRunning {
			return v
		}
	}
	return nil
}

// printReqLog
func printReqLog(ctx context.Context, req interface{}) {
	jsoon, _ := json.Marshal(ctx)
	log.Println(string(jsoon))

	jsoon, _ = json.Marshal(req)
	log.Println(string(jsoon))
}

// CreateRoom :
func (b *Backend) CreateRoom(ctx context.Context, req *pb.RoomCreateRequest) (*pb.Room, error) {
	printReqLog(ctx, req)
	b.mu.Lock()
	defer b.mu.Unlock()
	// search free box
	wkbox := b.checkAliveClient()
	if wkbox == nil {
		// busy
		log.Println("busy at " + time.Now().String())
		return nil, nil
	}
	tmptime := time.Now().String() + req.HostId

	// for loop it
	var f = ""
	for {
		f = cm.HashText(tmptime)
		l, err := wkbox.ListRem(&f)
		if err != nil {
			log.Fatal(err)
			return nil, err
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
		log.Fatalln(err)
		return &rmTmp, err
	}

	// create room success
	b.Roomlist = append(b.Roomlist, rmTmp)
	return &rmTmp, nil
}

// GetRoomList :
func (b *Backend) GetRoomList(ctx context.Context, req *pb.RoomListRequest) (res *pb.RoomListResponse, err error) {
	printReqLog(ctx, req)
	b.mu.Lock()
	defer b.mu.Unlock()
	// search free box
	wkbox := b.checkAliveClient()
	if wkbox == nil {
		// busy
		log.Println("busy at " + time.Now().String())
		return nil, nil
		// TODO : use chan to push task ?
	}
	var tmp pb.Room
	var RmList []pb.Room

	if _, err2 := wkbox.GetParaList(&req.Requirement, &RmList, tmp); err2 != nil {
		log.Fatalln(err)
		err = err2
	}
	for _, v := range RmList {
		res.Result = append(res.Result, &v)
	}
	err = nil
	return res, err
}

// GetRoomCurrentInfo :
func (b *Backend) GetRoomCurrentInfo(ctx context.Context, req *pb.RoomRequest) (*pb.Room, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method GetRoomCurrentInfo not implemented")
	printReqLog(ctx, req)
	b.mu.Lock()
	defer b.mu.Unlock()
	// search free box
	wkbox := b.checkAliveClient()
	if wkbox == nil {
		// busy
		log.Println("busy at " + time.Now().String())
		return nil, nil
		// TODO : use chan to push task ?
	}
	var tmp pb.Room

	if _, err := wkbox.GetPara(&req.Key, &tmp); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &tmp, nil
}

// GetRoomStream :
func (b *Backend) GetRoomStream(req *pb.RoomRequest, srv pb.RoomStatus_GetRoomStreamServer) error {
	// return status.Errorf(codes.Unimplemented, "method GetRoomStream not implemented")
	// if
	return nil
}

// UpdateRoomStatus :
func (b *Backend) UpdateRoomStatus(ctx context.Context, req *pb.CellStatus) (*types.Empty, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method UpdateRoomStatus not implemented")

	return nil, nil
}

// DeleteRoom :
func (b *Backend) DeleteRoom(ctx context.Context, req *pb.RoomRequest) (*types.Empty, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	printReqLog(ctx, req)
	b.mu.Lock()
	defer b.mu.Unlock()
	// search free box
	wkbox := b.checkAliveClient()
	if wkbox == nil {
		// busy
		log.Println("busy at " + time.Now().String())
		return nil, nil
		// TODO : use chan to push task ?
	}

	if _, err := wkbox.RemovePara(&req.Key); err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for k, v := range b.Roomlist {
		if v.Key == req.Key {
			b.Roomlist = append(b.Roomlist[:k], b.Roomlist[k+1:]...)
		}
	}
	log.Println("b.RoomList", b.Roomlist)
	return nil, nil
}

type workerResultBox struct {
	Result interface{}
	Error  error
}

func workerCreateRoom(procID int, reqJob <-chan *pb.RoomCreateRequest, room chan<- workerResultBox) {

}
