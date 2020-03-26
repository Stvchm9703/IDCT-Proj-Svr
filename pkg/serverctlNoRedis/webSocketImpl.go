package serverctlNoRedis

import (
	"fmt"
	"log"
	"net/http"
	"time"

	pb "RoomStatus/proto"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type SocketConn struct {
	// The websocket SocketConn.
	ws *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte
}

// SocketClient is a middleman between the websocket SocketConn and the hub.
type SocketClient struct {
	hub     *SocketHub
	conn    *websocket.Conn
	roomKey string
	send    chan []byte
}

// readPump pumps messages from the websocket SocketConn to the hub.
//
// The application runs readPump in a per-SocketConn goroutine. The application
// ensures that there is at most one reader on a SocketConn by executing all
// reads from this goroutine.
func (c *SocketClient) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		// Note: Igore client send out
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket SocketConn.
//
// A goroutine running writePump is started for each SocketConn. The
// application ensures that there is at most one writer to a SocketConn by
// executing all writes from this goroutine.
func (c *SocketClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(rsb *RoomStatusBackend, hub *SocketHub, c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(c.Param("roomId"))
	found := false
	for _, rm := range rsb.Roomlist {
		if rm.Key == c.Param("roomId") {
			found = true
		}
	}
	if found {
		client := &SocketClient{
			roomKey: c.Param("roomId"),
			hub:     hub,
			conn:    conn,
			send:    make(chan []byte, 256),
		}
		client.hub.register <- client
		go client.writePump()
		go client.readPump()
	} else {
		c.AbortWithStatus(412)
	}
}

type SocketHub struct {
	// Registered clients.
	clients map[*SocketClient]bool
	// Inbound messages from the clients.
	broadcast chan []byte
	// server-side cell-update
	cellCast chan *pb.CellStatusResp
	// Register requests from the clients.
	register chan *SocketClient
	// Unregister requests from clients.
	unregister chan *SocketClient
}

func newHub() *SocketHub {
	return &SocketHub{
		broadcast:  make(chan []byte),
		register:   make(chan *SocketClient),
		unregister: make(chan *SocketClient),
		cellCast:   make(chan *pb.CellStatusResp),
		clients:    make(map[*SocketClient]bool),
	}
}

// wraper to gin handler
func wrapfunc(rsb *RoomStatusBackend, hub *SocketHub) gin.HandlerFunc {
	return func(c *gin.Context) {
		serveWs(rsb, hub, c)
	}
}

func (h *SocketHub) run() {
	for {
		select {
		case client := <-h.register:
			fmt.Printf("\nclient regi :%#v \n", client)
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}

		case cellResp := <-h.cellCast:
			fmt.Printf("[WS:CellCast] Receive Cast\n\t- Message : %#v\n", cellResp)
			msgpt, _ := proto.Marshal(cellResp)
			for cli := range h.clients {
				if cli.roomKey == cellResp.Key {
					fmt.Println("have one :", cli.roomKey)
					select {
					case cli.send <- msgpt:
					default:
						close(cli.send)
						delete(h.clients, cli)
					}
				}
			}
		}
	}
}

func (rsb *RoomStatusBackend) RunWebSocketServer() error {
	hub := newHub()
	go hub.run()
	router := gin.New()
	router.GET("/:roomId", wrapfunc(rsb, hub))
	rsb.castServer1 = hub
	return router.Run(":8000")
}

func (rsb *RoomStatusBackend) BroadCast(cp *pb.CellStatusResp) error {
	fmt.Printf("[rsb:Broadcast]\n\t%#v\n", cp)
	rsb.castServer1.cellCast <- cp
	return nil
}
