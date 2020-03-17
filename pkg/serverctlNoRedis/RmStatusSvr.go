package serverctlNoRedis

import (
	// _ "RoomStatus"
	cm "RoomStatus/common"
	cf "RoomStatus/config"
	pb "RoomStatus/proto"
	"log"
	"sync"

	socketio "github.com/googollee/go-socket.io"
	"github.com/micro/go-micro/errors"
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomStatusServer = (*RoomStatusBackend)(nil)

// Remark: the framework make consider "instant" request
//
type RoomStatusBackend struct {
	// pb.RoomStatusServer
	mu         *sync.Mutex
	CoreKey    string
	Roomlist   []*pb.Room
	castServer *socketio.Server
}

// New : Create new backend
func New(conf *cf.ConfTmp) *RoomStatusBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)

	g := RoomStatusBackend{
		CoreKey: ck,
		mu:      &sync.Mutex{},
	}
	serv, err := g.InitSocketServer()
	if err != nil {
		panic(err)
	}
	g.castServer = serv
	return &g
}

func (this *RoomStatusBackend) SetCastServer(s *socketio.Server) error {
	if s == nil {
		return errors.New("001", "Empty SocketIO server", 0001)
	}
	this.castServer = s
	return nil
}

func (this *RoomStatusBackend) Shutdown() {
	log.Println("in shtdown proc")
	this.BroadCastShutdown()
	this.CloseDB()
	log.Println("endof shutdown proc:", this.CoreKey)
}
