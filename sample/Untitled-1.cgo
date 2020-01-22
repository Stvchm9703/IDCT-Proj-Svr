/*
*
*     Author        : tuxpy
*     Email         : q8886888@qq.com.com
*     Create time   : 3/7/18 9:18 AM
*     Filename      : service.go
*     Description   :
*
*  https://codertw.com/%E4%BC%BA%E6%9C%8D%E5%99%A8/152172/
 */


syntax = "proto3";
package helloword;
import "github.com/golang/protobuf/ptypes/timestamp/timestamp.proto";
service Greeter {
	rpc SayHello(stream HelloRequest) returns (stream HelloReply) {};
}
message HelloRequest {
	string message = 1;
}
message HelloReply {
	string message = 1;
	google.protobuf.Timestamp TS = 2;
	MessageType message_type = 3;
	enum MessageType{
		CONNECT_SUCCESS = 0;
		CONNECT_FAILED = 1;
		NORMAL_MESSAGE = 2;
	}
}
package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	pb "grpclb/helloword"
	"log"
	"net"
	"os"
	"sync"
	"time"
	"utils"
)

type Service struct{}
type ConnectPool struct {
	sync.Map
}

var connect_pool *ConnectPool

func (p *ConnectPool) Get(name string) pb.Greeter_SayHelloServer {
	if stream, ok := p.Load(name); ok {
		return stream.(pb.Greeter_SayHelloServer)
	} else {
		return nil
	}
}
func (p *ConnectPool) Add(name string, stream pb.Greeter_SayHelloServer) {
	p.Store(name, stream)
}
func (p *ConnectPool) Del(name string) {
	p.Delete(name)
}
func (p *ConnectPool) BroadCast(from, message string) {
	log.Printf("BroadCast from: %s, message: %s\n", from, message)
	p.Range(func(username_i, stream_i interface{}) bool {
		username := username_i.(string)
		stream := stream_i.(pb.Greeter_SayHelloServer)
		if username == from {
			return true
		} else {
			stream.Send(&pb.HelloReply{
				Message:     message,
				MessageType: pb.HelloReply_NORMAL_MESSAGE,
				TS:          &timestamp.Timestamp{Seconds: time.Now().Unix()},
			})
		}
		return true
	})
}
func (s *Service) SayHello(stream pb.Greeter_SayHelloServer) error {
	peer, _ := peer.FromContext(stream.Context())
	log.Printf("Received new connection.  %s", peer.Addr.String())
	md, _ := metadata.FromIncomingContext(stream.Context())
	username := md["name"][0]
	if connect_pool.Get(username) != nil {
		stream.Send(&pb.HelloReply{
			Message:     fmt.Sprintf("username %s already exists!", username),
			MessageType: pb.HelloReply_CONNECT_FAILED,
		})
		return nil
	} else { // 連線成功
		connect_pool.Add(username, stream)
		stream.Send(&pb.HelloReply{
			Message:     fmt.Sprintf("Connect success!"),
			MessageType: pb.HelloReply_CONNECT_SUCCESS,
		})
	}
	go func() {
		<-stream.Context().Done()
		connect_pool.Del(username)
		connect_pool.BroadCast(username, fmt.Sprintf("%s leval room", username))
	}()
	connect_pool.BroadCast(username, fmt.Sprintf("Welcome %s!", username))
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		connect_pool.BroadCast(username, fmt.Sprintf("%s: %s", username, req.Message))
	}
	return nil
}
func GetListen() string {
	if len(os.Args) < 2 {
		return ":8881"
	}
	return os.Args[1]
}
func main() {
	connect_pool = &ConnectPool{}
	lis, err := net.Listen("tcp", GetListen())
	utils.CheckErrorPanic(err)
	fmt.Println("Listen on", GetListen())
	s := grpc.NewServer(grpc.RPCCompressor(grpc.NewGZIPCompressor()),
		grpc.RPCDecompressor(grpc.NewGZIPDecompressor()))
	pb.RegisterGreeterServer(s, &Service{})
	utils.CheckErrorPanic(s.Serve(lis))
}
