package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"mime"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gogo/gateway"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	cm "RoomStatus/common"
	cf "RoomStatus/config"
	"RoomStatus/insecure"
	pb "RoomStatus/proto"
	server "RoomStatus/serverGameCtl"

	// Static files
	_ "RoomStatus/statik"
)

var (
	gRPCPort    = flag.Int("grpc-port", 10000, "The gRPC server port")
	gatewayPort = flag.Int("gateway-port", 11000, "The gRPC-Gateway server port")
)

var log grpclog.LoggerV2

func init() {
	f, _ := os.Getwd()
	dt := time.Now()
	w := cm.SetLog(filepath.Join(f, "tmp", "log", dt.Format("01-02-2006")+".log"))
	log = grpclog.NewLoggerV2(w, w, w)
	grpclog.SetLoggerV2(log)
}

// serveOpenAPI serves an OpenAPI UI on /openapi-ui/
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func serveOpenAPI(mux *http.ServeMux) error {
	mime.AddExtensionType(".svg", "image/svg+xml")

	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	// Expose files in static on <host>/openapi-ui
	fileServer := http.FileServer(statikFS)
	prefix := "/openapi-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	return nil
}

var testing_config = cf.ConfTmp{
	cf.CfTemplServer{},
	cf.CfAPIServer{
		ConnType:     "TCP",
		IP:           "127.0.0.1",
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
	flag.Parse()
	// addr := fmt.Sprintf("127.0.0.1:%d", *gRPCPort)
	addr := "127.0.0.1:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic("Failed to listen:\t"+ err.Error())
	}
	s := grpc.NewServer(
		grpc.Creds( credentials.NewServerTLSFromCert(insecure.Cert)),
		grpc.UnaryInterceptor( grpc_validator.UnaryServerInterceptor()),
		grpc.StreamInterceptor( grpc_validator.StreamServerInterceptor()),
	)
	// s.GracefulStop()
	pb.RegisterRoomStatusServer(
		s,
		server.New(&testing_config))

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		panic(s.Serve(lis))
	}()

	// See https://github.com/grpc/grpc/blob/master/doc/naming.md
	// for gRPC naming standard information.
	dialAddr := fmt.Sprintf("passthrough://127.0.0.1/%s", addr)
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithTransportCredentials(
			credentials.NewClientTLSFromCert(
				insecure.CertPool, "",
			),
		),
		grpc.WithBlock(),
	)
	if err != nil {
		panic("Failed to dial server:\t"+err.Error())
	}

	mux := http.NewServeMux()

	jsonpb := &gateway.JSONPb{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
		// This is necessary to get error details properly
		// marshalled in unary requests.
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)
	err = pb.RegisterRoomStatusHandler(
		context.Background(), gwmux, conn)
	if err != nil {
		panic("Failed to register gateway:\t"+err.Error())
	}

	mux.Handle("/", gwmux)
	err = serveOpenAPI(mux)
	if err != nil {
		panic("Failed to serve OpenAPI UI")
	}

	gatewayAddr := fmt.Sprintf("localhost:%d", *gatewayPort)
	log.Info("Serving gRPC-Gateway on https://", gatewayAddr)
	log.Info("Serving OpenAPI Documentation on https://", gatewayAddr, "/openapi-ui/")
	gwServer := http.Server{
		Addr: gatewayAddr,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*insecure.Cert},
		},
		Handler: mux,
	}
	panic(gwServer.ListenAndServeTLS("", ""))
}
