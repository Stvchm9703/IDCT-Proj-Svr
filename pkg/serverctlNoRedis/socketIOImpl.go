package serverctlNoRedis

import (
	pb "RoomStatus/proto"
	"encoding/json"
	"fmt"
	"log"
	"strings"
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
	fmt.Printf("\nserver-option\n\t%#v\n", server)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("[SocketIO][Connect]")
		fmt.Println("\t- NameSpace : /")
		fmt.Println("\t- connected:", s.ID())
		fmt.Printf("\t- Room Request : \n\t\t%#v \n\n", s.URL())
		return nil
	})

	server.OnEvent("/", "join_room", func(s socketio.Conn, msg string) string {
		fmt.Printf("\n[Socket.IO][JoinRoom]\n\t- req-room:\t%v\n\n", msg)
		for _, v := range this.Roomlist {
			if msg == v.Key {
				s.Join(v.Key)
				return proto.MarshalTextString(v)
			}
		}
		return "Not_Found"
	})

	// chatroom
	server.OnEvent("/", "chat_msg", func(s socketio.Conn, msg string) {
		s.SetContext(msg)
		fmt.Println("chat_msg:", msg)
		room := s.Rooms()
		fmt.Println("rooms_:\t", room)
		ind := -1
		for k := range room {
			if strings.Contains(room[k], "Rm") {
				ind = k
			}
		}
		if ind == -1 {
			return
		}
		if server.BroadcastToRoom("/", room[ind], "chat_msg_recv", msg) == false {
			fmt.Printf("fall out? %v\n", room[ind])
			s.Emit("chat_msg_recv", msg)
		}
	})

	server.OnEvent("/", "close", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Printf("\n[Socket.IO][Error]\n\t error:%#v \n\tSocket Client %#v", e, s.ID())
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		s.Close()
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

// func (this *RoomStatusBackend) BroadCast(msg *pb.CellStatusResp) error {
// 	fmt.Println("Broadcast-msg")
// 	if this.castServer == nil {
// 		fmt.Println("Broadcast Not Inited")
// 		return status.Error(codes.Internal, "Broadcast Not Inited")
// 	}
// 	msgpt, _ := proto.Marshal(msg)
// 	fmt.Println(this.castServer.BroadcastToRoom("/", msg.Key, "syst_msg", (msgpt)))
// 	return nil
// }

func (this *RoomStatusBackend) BroadCastRaw(msg *pb.CellStatusResp) error {
	if this.castServer == nil {
		return status.Error(codes.Internal, "Broadcast Not Inited")
	}
	msgpt, _ := json.Marshal(msg)
	this.castServer.BroadcastToRoom("/", msg.Key, "syst_msg", msgpt)
	return nil
}

func (this *RoomStatusBackend) BroadCastShutdown() error {
	// this.castServer.LeaveAllRooms("/room", )
	roomlist := this.castServer.Rooms("/")

	wg := sync.WaitGroup{}
	wg.Add(len(roomlist))
	for _, v := range roomlist {
		go func(v string) {
			this.castServer.BroadcastToRoom("/", v, "syst_msg", "system_shutdown")
			this.castServer.ClearRoom("/", v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	return this.castServer.Close()
}
