package serverctlNoRedis

// // ======================================================================================================
// // RoomMgr : Room Manager
// type RoomMgr struct {
// 	pb.Room
// 	get_only_stream map[string]*pb.RoomStatus_GetRoomStreamServer
// 	// close_link      *sync.Map
// }

// // ----------------------------------------------------------------------------------------------------
// // roommgr.get_only_stream

// func (rm *RoomMgr) GetGS(user_id string) *pb.RoomStatus_GetRoomStreamServer {
// 	a, ok := rm.get_only_stream[user_id]
// 	if ok {
// 		return a
// 	}
// 	return nil
// }

// func (rm *RoomMgr) AddGS(user_id string, stream *pb.RoomStatus_GetRoomStreamServer) (bool, error) {
// 	_, ok := rm.get_only_stream[user_id]
// 	if ok {
// 		return false, errors.New("StreamExist")
// 	}

// 	rm.get_only_stream[user_id] = stream
// 	return true, nil
// }

// func (rm *RoomMgr) DelGS(user_id string) (bool, error) {
// 	log.Println("Del Stream:", user_id)
// 	if rm.get_only_stream[user_id] != nil {
// 		*(rm.get_only_stream[user_id]) = nil
// 		delete(rm.get_only_stream, user_id)
// 		return true, nil
// 	}
// 	return false, errors.New("StreamNotExist")
// }
// func (rm *RoomMgr) BroadCastGS(from string, message *pb.CellStatusResp) {
// 	log.Println(rm.get_only_stream)

// 	for k, v := range rm.get_only_stream {
// 		if k != from {
// 			tmpv := *v
// 			tmpv.Send(message)
// 		}
// 	}

// }

// // ---------------------------------------------------------------------------------------------

// // RoomMgr
// func (rm *RoomMgr) BroadCast(from string, message *pb.CellStatusResp) {
// 	log.Println("BS!", message)
// 	rm.BroadCastGS(from, message)
// }
// func (rm *RoomMgr) ClearAll() {
// 	log.Println("ClearAll Proc")
// 	for k := range rm.get_only_stream {
// 		fmt.Println(k)
// 		rm.DelGS(k)
// 	}
// }
