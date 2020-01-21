package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"

	cf "RoomStatus/config"
	pb "RoomStatus/proto"
	server "RoomStatus/serverctlNoRedis"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"
	// Static files
	_ "RoomStatus/statik"
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
		Port:         11000,
		MaxPoolSize:  20,
		APIReferType: "proto",
		APITablePath: "{root}/thrid_party/OpenAPI",
		APIOutpath:   "./",
	},
	cf.CfTDatabase{
		Connector:  "redis",
		WorkerNode: 12,
		Host:       "192.168.0.110",
		Port:       6379,
		Username:   "",
		Password:   "",
		Database:   "redis",
		Filepath:   "",
	},
}

func main() {
	log.Println("start run")
	addr := testing_config.APIServer.IP + ":" + strconv.Itoa(testing_config.APIServer.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic("Failed to listen:\t" + err.Error())
	}
	s := grpc.NewServer(
		// grpc.Creds(credentials.NewServerTLSFromCert(insecure.Cert)),
		grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
	)

	RMServer := server.New(&testing_config)
	// s.GracefulStop()
	pb.RegisterRoomStatusServer(
		s, RMServer)

	// Serve gRPC Server
	log.Println("Serving gRPC on https://", addr)
	go func() {
		panic(s.Serve(lis))
	}()
	BeforeGracefulStop(s, RMServer)

	// call your cleanup method with this channel as a routine
}
func BeforeGracefulStop(ss *grpc.Server, rms *server.RoomStatusBackend) {
	log.Println("BeforeGracefulStop")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	aa := <-c
	log.Println("OS.signal", aa)
	log.Println(ss.GetServiceInfo())
	// ss.Shutdown()
	rms.Shutdown()
	ss.GracefulStop()
	log.Println("os GracefulStop")
	os.Exit(0)
}
