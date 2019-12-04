package serverctl

import (
	pb "RoomStatus/proto"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gogo/status"
)

// GetRoomList :
func (b *RoomStatusBackend) GetRoomList(ctx context.Context, req *pb.RoomListRequest) (*pb.RoomListResponse, error) {

	printReqLog(ctx, req)
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()
	wkbox := b.searchAliveClient()
	// var tmp pb.Room
	var RmList []*pb.Room
	strl, err2 := wkbox.GetParaList(&req.Requirement)
	if err2 != nil {
		log.Fatalln(err2)
		// ignore err ()
	}
	// log.Println("strl:", string(*strl))
	err2 = json.Unmarshal(*strl, &RmList)
	if err2 != nil {
		log.Fatalln(err2)
		return nil, status.Errorf(500, err2.Error())
	}
	(wkbox).Preserve(false)
	// log.Println("list:", RmList)
	// log.Println("typeof:", reflect.TypeOf(RmList))
	res := &pb.RoomListResponse{
		Result: RmList,
	}
	return res, nil
}
