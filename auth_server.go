package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	cf "RoomStatus/config"
	server "RoomStatus/pkg/authServer"
	pb "RoomStatus/proto"

	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// Static files
	// _ "RoomStatus/statik"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)
var testing_config = cf.ConfTmp{
	cf.CfTemplServer{
		IP:               "0.0.0.0",
		Port:             9000,
		RootFilePath:     "",
		MainPath:         "",
		StaticFilepath:   "",
		StaticOutpath:    "",
		TemplateFilepath: "",
		TemplateOutpath:  "",
	},
	cf.CfAPIServer{
		ConnType:     "TCP",
		IP:           "0.0.0.0",
		Port:         12000,
		MaxPoolSize:  20,
		APIReferType: "grpc",
		APITablePath: "{root}/thrid_party/OpenAPI",
		APIOutpath:   "./",
	},

	cf.CfTDatabase{
		Connector:  "postgres",
		WorkerNode: 1,
		Host:       "127.0.0.1",
		Port:       5432,
		Username:   "",
		Password:   "",
		Database:   "idct_db",
		Filepath:   "",
	},
	cf.CfTDatabase{
		Connector:  "redis",
		WorkerNode: 12,
		Host:       "192.168.0.110",
		Port:       6379,
		Username:   "test",
		Password:   "",
		Database:   "redis",
		Filepath:   "",
	},
	cf.CfAPIServer{}
}

func main() {
	log.Println("start run")
	addr := testing_config.APIServer.IP + ":" + strconv.Itoa(testing_config.APIServer.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic("Failed to listen:\t" + err.Error())
	}
	// d := insecure.Cert
	// log.Println(d)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
	)

	RMServer := server.New(&testing_config)

	pb.RegisterCreditsAuthServer(
		s, RMServer)
	log.Println("Serving gRPC on https://", addr)
	go func() {
		panic(s.Serve(lis))
	}()
	beforeGracefulStop(s, RMServer)

	// call your cleanup method with this channel as a routine
}
func beforeGracefulStop(ss *grpc.Server, rms *server.CreditsAuthBackend) {
	log.Println("BeforeGracefulStop")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	aa := <-c
	log.Println("OS.signal", aa)
	log.Println(ss.GetServiceInfo())
	// ss.Shutdown()
	rms.Shutdown()
	ss.Stop()
	log.Println("os GracefulStop")
	os.Exit(0)
}
