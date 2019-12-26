package serverctlNoRedis

// pb.GetRoomStream
// func (*UnimplementedRoomStatusServer) GetRoomStream(req *RoomRequest, srv RoomStatus_GetRoomStreamServer) error {
// 	return status.Errorf(codes.Unimplemented, "method GetRoomStream not implemented")
// }

// roomStreamWkTask : wk channal to handle stream
// func (b *RoomStatusBackend) roomStreamWkTask(payload interface{}) {
// 	req, ok := payload.(WkTask).In.(*pb.RoomRequest)
// 	if !ok {
// 		return
// 	}
// 	wkbox := b.searchAliveClient()
// 	var tmp pb.Room
// 	if _, err := wkbox.GetPara(&req.Key, &tmp); err != nil {
// 		log.Fatalln(err)
// 	}
// 	if tmp.Key != "" {
// 		strm, ok := payload.(WkTask).Stream.(pb.RoomStatus_GetRoomStreamServer)
// 		if !ok {
// 			log.Println("stream server casting error")
// 			return
// 		}
// 		// create redis pubsub client
// 		pub := rds.CBSCopyFromClient(wkbox)
// 		channel, err := pub.AddChannel("*" + req.Key)
// 		if err != nil {
// 			payload.(WkTask).Out <- err
// 			return
// 		}
// 		b.sub = append(b.sub, pub)
// 		wkbox.Preserve(false)
// 		for msg := range channel {
// 			log.Println(msg)
// 			if e := strm.Send(&pb.CellStatus{}); e != nil {
// 				payload.(WkTask).Out <- e
// 				pub.Disconn()
// 				return
// 			}

// 		}
// 	} else {
// 		// error return :
// 		// !TODO add error message for Out!
// 		return
// 	}
// }

// GetRoomStream :
// func (b *RoomStatusBackend) GetRoomStream(req *pb.RoomRequest, srv pb.RoomStatus_GetRoomStreamServer) (err error) {
// 	// printReqLog(ctx, req)
// 	// ===== Worker Start ======
// 	pl := &WkTask{
// 		In:     req,
// 		Out:    make(chan interface{}),
// 		Stream: srv,
// 	}
// 	if err = b.steamWk.Invoke(pl); err != nil {
// 		log.Println("err in GetList Wk", err)
// 		return
// 	}
// 	// ====== Worker End =======
// 	rm := <-(pl).Out
// 	err = rm.(error)

// 	return
// }

// pb.UpdateRoomStream
