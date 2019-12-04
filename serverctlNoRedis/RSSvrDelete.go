package serverctlNoRedisl

import (
	pb "RoomStatus/proto"
	"context"
	"errors"
	"log"

	types "github.com/gogo/protobuf/types"
)

func (b *RoomStatusBackend) deleteWkTask(payload interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()

	req, ok := payload.(WkTask).In.(*pb.RoomRequest)
	if !ok {
		return
	}

	wkbox := b.searchAliveClient()

	if _, err := (wkbox).RemovePara(&req.Key); err != nil {
		log.Fatalln(err)
		return
	}
	(wkbox).Preserve(false)
	payload.(WkTask).Out <- &req.Key
	return
}

// TestCreateWkTask : Test Unit
// func (b *RoomStatusBackend) TestDeteleWkTask(pl interface{}) (rmTmp *pb.Room, err error) {
// 	if err := b.deleteWk.Invoke(pl.(WkTask)); err != nil {
// 		log.Println("err in create Wk", err)
// 		return nil, err
// 	}
// 	// ====== Worker End =======
// 	plc := <-(pl.(WkTask)).Out
// 	for k, v := range b.Roomlist {
// 		if v.Key == *plc.(*string) {
// 			rmTmp = b.Roomlist[k]
// 			b.Roomlist = append(b.Roomlist[:k], b.Roomlist[k+1:]...)
// 		}
// 	}
// 	return
// }

// DeleteRoom :
func (b *RoomStatusBackend) DeleteRoom(ctx context.Context, req *pb.RoomRequest) (*types.Empty, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	b.mu.Lock()
	defer b.mu.Unlock()

	// wkbox := b.searchAliveClient()

	// if _, err := (wkbox).RemovePara(&req.Key); err != nil {
	// 	log.Fatalln(err)
	// 	return nil, err
	// }
	// (wkbox).Preserve(false)
	done := false
	for k, v := range b.Roomlist {
		if v.Key == req.Key {
			// rmTmp = b.Roomlist[k]
			log.Println(b.Roomlist[k])
			b.Roomlist = append(b.Roomlist[:k], b.Roomlist[k+1:]...)
			done = true
		}
	}
	if !done {
		return nil, errors.New("RoomNotExist")
	}
	log.Println("b.RoomList", b.Roomlist)
	return nil, nil
}

//
