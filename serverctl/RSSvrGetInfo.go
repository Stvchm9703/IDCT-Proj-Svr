package serverctl

import (
	pb "RoomStatus/proto"
	"context"
	"log"
)

func (b *RoomStatusBackend) getInfoWkTask(payload interface{}) {
	req, ok := payload.(WkTask).In.(*pb.RoomRequest)
	if !ok {
		return
	}

	wkbox := b.searchAliveClient()
	var tmp pb.Room
	if _, err := wkbox.GetPara(&req.Key, &tmp); err != nil {
		log.Fatalln(err)
		// return nil, err
	}
	(wkbox).Preserve(false)
	payload.(WkTask).Out <- &tmp
	return
}

// func (b *RoomStatusBackend) TestGetInfoWkTask(pl interface{}) (rmTmp *pb.Room, err error) {
// 	if err = b.getWk.Invoke(pl.(WkTask)); err != nil {
// 		log.Println("err in create Wk", err)
// 		return
// 	}
// 	// ====== Worker End =======
// 	rm := <-(pl.(WkTask)).Out
// 	rmTmp = rm.(*pb.Room)
// 	return
// }

// GetRoomInfo :
func (b *RoomStatusBackend) GetRoomInfo(ctx context.Context, req *pb.RoomRequest) (*pb.Room, error) {

	printReqLog(ctx, req)
	wkbox := b.searchAliveClient()
	var tmp pb.Room
	if _, err := wkbox.GetPara(&req.Key, &tmp); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	(wkbox).Preserve(false)
	return &tmp, nil
}
