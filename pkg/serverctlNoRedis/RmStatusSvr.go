package serverctlNoRedis

import (
	// _ "RoomStatus"
	cm "RoomStatus/common"
	cf "RoomStatus/config"
	pb "RoomStatus/proto"
	"log"
	"sync"

	socketio "github.com/googollee/go-socket.io"
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomStatusServer = (*RoomStatusBackend)(nil)

// Remark: the framework make consider "instant" request
//
type RoomStatusBackend struct {
	// pb.RoomStatusServer
	mu          *sync.Mutex
	CoreKey     string
	Roomlist    []*pb.Room
	castServer  *socketio.Server
	castServer1 *SocketHub
}

// New : Create new backend
func New(conf *cf.ConfTmp) *RoomStatusBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)

	g := RoomStatusBackend{
		CoreKey: ck,
		mu:      &sync.Mutex{},
	}
	// serv, err := g.InitSocketServer()
	// if err != nil {
	// 	panic(err)
	// }
	// g.castServer = serv
	return &g
}

func (this *RoomStatusBackend) Shutdown() {
	log.Println("in shtdown proc")
	// this.BroadCastShutdown()
	this.CloseDB()
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

// PrintReqLog
