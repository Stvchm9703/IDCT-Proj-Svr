package authServer

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	cf "RoomStatus/config"
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

func ServerMainProcess(testing_config *cf.ConfTmp) {
	log.Println("start run")
	addr := testing_config.AuthServer.IP + ":" + strconv.Itoa(testing_config.AuthServer.Port)
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

	RMServer := New(testing_config)

	pb.RegisterCreditsAuthServer(
		s, RMServer)
	log.Println("Serving gRPC on https://", addr)
	go func() {
		panic(s.Serve(lis))
	}()
	beforeGracefulStop(s, RMServer)

	// call your cleanup method with this channel as a routine
}
func beforeGracefulStop(ss *grpc.Server, rms *CreditsAuthBackend) {
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
