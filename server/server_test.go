package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	RoomStatus "RoomStatus/pkg/protos"

	"github.com/lileio/lile"
)

var s = RoomStatusServer{}
var cli RoomStatus.RoomStatusClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		RoomStatus.RegisterRoomStatusServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = RoomStatus.NewRoomStatusClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
