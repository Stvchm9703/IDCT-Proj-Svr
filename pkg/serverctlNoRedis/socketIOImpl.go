package serverctlNoRedis

import (
	pb "RoomStatus/proto"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	socketio "github.com/googollee/go-socket.io"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// InitSocketServer :
// since the broadcast issue case [Shutdown ungradefully]
// this implement socket.io websocket server is used for client receive only
func (this *RoomStatusBackend) InitSocketServer() (*socketio.Server, error) {

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("NameSpace : /")
		fmt.Println("connected:", s.ID())
		fmt.Printf("Room Request : %#v \n", s.URL())
		return nil
	})

	// room server support
	server.OnEvent("/room", "room_list", func(s socketio.Conn, msg string) []string {
		fmt.Println("notice:", msg)
		ls := []string{}
		for _, v := range this.Roomlist {
			ls = append(ls, proto.MarshalTextString(v))
		}
		// s.Emit("reply", ls)
		return ls
	})

	server.OnEvent("/room", "join_room", func(s socketio.Conn, msg string) string {
		for _, v := range this.Roomlist {
			if msg == v.Key {
				s.Join(v.Key)
				return proto.MarshalTextString(v)
			}
		}
		// s.Emit("reply", ls)
		return "Not_Found"
	})

	server.OnEvent("/room", "update_state", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	// chatroom
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) {
		s.SetContext(msg)
		s.Emit("recv", msg)
		// return "recv " + msg
	})
	server.OnEvent("/", "close", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	return server, err
}

func (this *RoomStatusBackend) RunSocketServer() error {
	log.Println("Serve Socket Server")
	router := gin.New()

	go this.castServer.Serve()

	router.GET("/socket.io/*any", gin.WrapH(this.castServer))
	router.POST("/socket.io/*any", gin.WrapH(this.castServer))
	return router.Run(":8000")
	// return nil
}

func (this *RoomStatusBackend) BroadCast(msg *pb.CellStatusResp) error {
	if this.castServer == nil {
		return status.Error(codes.Internal, "Broadcast Not Inited")
	}
	msgpt := proto.MarshalTextString(msg)
	this.castServer.BroadcastToRoom("/room", msg.Key, "SystMsg", msgpt)
	return nil
}

func (this *RoomStatusBackend) BroadCastRaw(msg *pb.CellStatusResp) error {
	if this.castServer == nil {
		return status.Error(codes.Internal, "Broadcast Not Inited")
	}
	msgpt, _ := json.Marshal(msg)
	this.castServer.BroadcastToRoom("/room", msg.Key, "SystMsg", msgpt)
	return nil
}

func (this *RoomStatusBackend) BroadCastShutdown() error {
	// this.castServer.LeaveAllRooms("/room", )
	roomlist := this.castServer.Rooms("/room")

	wg := sync.WaitGroup{}
	wg.Add(len(roomlist))
	for _, v := range roomlist {
		go func(v string) {
			this.castServer.BroadcastToRoom("/room", v, "Syst_Msg", "system_shutdown")
			this.castServer.ClearRoom("/room", v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	return this.castServer.Close()
}
