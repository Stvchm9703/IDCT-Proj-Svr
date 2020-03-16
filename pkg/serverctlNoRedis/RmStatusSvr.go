package serverctlNoRedis

import (
	// _ "RoomStatus"
	cm "RoomStatus/common"
	cf "RoomStatus/config"
	pb "RoomStatus/proto"
	"log"
	"sync"
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomStatusServer = (*RoomStatusBackend)(nil)

// Remark: the framework make consider "instant" request
//
type RoomStatusBackend struct {
	// pb.RoomStatusServer
	mu       *sync.Mutex
	CoreKey  string
	Roomlist []*pb.Room
}

// New : Create new backend
func New(conf *cf.ConfTmp) *RoomStatusBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)

	g := RoomStatusBackend{
		CoreKey: ck,
		mu:      &sync.Mutex{},
	}
	g.InitDB(&conf.Database)
	return &g
}

func (this *RoomStatusBackend) Shutdown() {
	log.Println("in shtdown proc")
	/// TODO: send closing msg to all client
	// for _, v := range this.Roomlist {
	// 	log.Println("Server OS.sigKill")

	// !Broadcast
	// 	v.BroadCast("RmSvrMgr",
	// 		&pb.CellStatusResp{
	// 			UserId:    "RmSvrMgr",
	// 			Key:       v.Key,
	// 			Status:    201,
	// 			Timestamp: time.Now().String(),
	// 			ResponseMsg: &pb.CellStatusResp_ErrorMsg{
	// 				ErrorMsg: &pb.ErrorMsg{
	// 					MsgInfo: "ConnEnd",
	// 					MsgDesp: "Server OS.sigKill",
	// 				},
	// 			},
	// 		})
	// 	v.ClearAll()
	// }
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

// ======================================================================================================
// RoomMgr : Room Manager
