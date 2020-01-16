package serverctlNoRedis

// pb.GetRoomStream
// func (*UnimplementedRoomStatusServer) GetRoomStream(req *RoomRequest, srv RoomStatus_GetRoomStreamServer) error {
// 	return status.Errorf(codes.Unimplemented, "method GetRoomStream not implemented")
// }
// GetRoomStream(*RoomRequest, RoomStatus_GetRoomStreamServer) error

// func (b *RoomStatusBackend) RoomStream(stream pb.RoomStatus_RoomStreamServer) (err error) {
// 	// printReqLog(ctx, req)

// 	var RmList []*pb.Room

// 	peer, _ := peer.FromContext(stream.Context())
// 	log.Printf("Received new connection.  %s", peer.Addr.String())
// 	md, _ := metadata.FromIncomingContext(stream.Context())
// 	username := md["name"][0]
// 	if connect_pool.Get(username) != nil {
// 		stream.Send(&pb.CellStatusResp{
// 			Status: 500,
// 		})
// 		return nil
// 	} else {
// 		connect_pool.Add(username, stream)
// 		stream.Send(&pb.HelloReply{
// 			Message:     fmt.Sprintf("Connect success!"),
// 			MessageType: pb.HelloReply_CONNECT_SUCCESS,
// 		})
// 	}
// 	go func() {
// 		<-stream.Context().Done()
// 		connect_pool.Del(username)
// 		connect_pool.BroadCast(username, fmt.Sprintf("%s leval room", username))
// 	}()
// 	connect_pool.BroadCast(username, fmt.Sprintf("Welcome %s!", username))
// 	for {
// 		req, err := stream.Recv()
// 		if err != nil {
// 			return err
// 		}
// 		connect_pool.BroadCast(username, fmt.Sprintf("%s: %s", username, req.Message))
// 	}
// 	return nil
// 	log.Println("list:", RmList)
// 	log.Println("typeof:", reflect.TypeOf(RmList))

// 	return
// }

// pb.UpdateRoomStream
