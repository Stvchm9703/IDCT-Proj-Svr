package serverctl

import (
	pb "RoomStatus/proto"
	"context"
	"encoding/json"
	"log"
	"reflect"
)

func (b *RoomStatusBackend) getLsWkTask(payload interface{}) {
	req, ok := payload.(WkTask).In.(*pb.RoomListRequest)
	if !ok {
		return
	}

	wkbox := b.searchAliveClient()
	// var tmp pb.Room
	var RmList []*pb.Room
	strl, err2 := wkbox.GetParaList(&req.Requirement)
	if err2 != nil {
		log.Fatalln(err2)
	}
	log.Println("strl:", string(*strl))
	err := json.Unmarshal(*strl, &RmList)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(RmList)
	(wkbox).Preserve(false)
	payload.(WkTask).Out <- RmList
	return
}

// func (b *RoomStatusBackend) TestGetLsWkTask(pl interface{}) (rmTmp *pb.RoomListResponse, err error) {
// 	if err = b.getListWk.Invoke(pl.(WkTask)); err != nil {
// 		log.Println("err in GetList Wk", err)
// 		return
// 	}
// 	// ====== Worker End =======
// 	rm := <-(pl.(WkTask)).Out
// 	fg := rm.([]*pb.Room)
// 	rmTmp = &pb.RoomListResponse{
// 		Result: fg,
// 	}
// 	return
// }

// GetRoomList :
func (b *RoomStatusBackend) GetRoomList(ctx context.Context, req *pb.RoomListRequest) (res *pb.RoomListResponse, err error) {
	printReqLog(ctx, req)
	wkbox := b.searchAliveClient()
	// var tmp pb.Room
	var RmList []*pb.Room
	strl, err2 := wkbox.GetParaList(&req.Requirement)
	if err2 != nil {
		log.Fatalln(err2)
	}
	log.Println("strl:", string(*strl))
	err = json.Unmarshal(*strl, &RmList)
	if err != nil {
		log.Fatalln(err)
		return
	}
	(wkbox).Preserve(false)
	log.Println("list:", RmList)
	log.Println("typeof:", reflect.TypeOf(RmList))
	res = &pb.RoomListResponse{
		Result: RmList,
	}
	return
}
