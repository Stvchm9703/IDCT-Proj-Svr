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
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

type Backend struct {
	pb.UnimplementedRoomStatusServer
	mu *sync.Mutex
	// channel
	redChannel []chan *rd.RdsCliBox
	// Roomlist []*pb.Room
	redhdlr []*rd.RdsCliBox
	CoreKey string
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
	if _, err := wkbox.SetPara(&f, rmTmp); err != nil {
		log.Fatalln(err)
		return &rmTmp, err
	}
	return &rmTmp, nil
}

// GetRoomList :
func (b *Backend) GetRoomList(ctx context.Context, req *pb.RoomListRequest) (*pb.RoomListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomList not implemented")
}

// GetRoomCurrentInfo :
func (b *Backend) GetRoomCurrentInfo(ctx context.Context, req *pb.RoomRequest) (*pb.Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomCurrentInfo not implemented")
}

// GetRoomStream :
func (b *Backend) GetRoomStream(req *pb.RoomRequest, srv pb.RoomStatus_GetRoomStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRoomStream not implemented")
}

// UpdateRoomStatus :
func (b *Backend) UpdateRoomStatus(ctx context.Context, req *pb.CellStatus) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoomStatus not implemented")
}

// DeleteRoom :
func (b *Backend) DeleteRoom(ctx context.Context, req *pb.RoomRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}

type workerResultBox struct {
	Result interface{}
	Error  error
}

func workerCreateRoom(procID int, reqJob <-chan *pb.RoomCreateRequest, room chan<- workerResultBox) {

}
