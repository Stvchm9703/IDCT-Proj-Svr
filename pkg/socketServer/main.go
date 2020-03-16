package socketServer


import (
	RmSr "RoomStatus/pkg/serverctlNoRedis"
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

// InitSocketServer :
// since the broadcast issue case [Shutdown ungradefully]
// this implement socket.io websocket server is used for client receive only
func InitSocketServer( gs *RmSr.RoomStatusBackend ) ( *socketio.Server, error) {

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("NameSpace : /")
		fmt.Println("connected:", s.ID())
		fmt.Printf("Room Request : %#v \n", s.URL())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
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
	log.Println("Serve Socket Server")
	return server, nil

}

func RunSocketServer ( server *socketio.Server) error{
	router := gin.New()
	go server.Serve()
	// defer server.Close()
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	return  router.Run(":8000")
}


func SocketShutdown(s *socketio.Server) error {
	// this.castServer.LeaveAllRooms("/room", )
	roomlist := s.Rooms("/room")
	wg := sync.WaitGroup{}
	wg.Add(len(roomlist))
	for _, v := range roomlist {
		go func(v string) {
			s.BroadcastToRoom("/room", v, "Syst_Msg", "system_shutdown")
			s.ClearRoom("/room", v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	return s.Close()
}
