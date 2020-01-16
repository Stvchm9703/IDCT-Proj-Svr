package v2

import (
	pb "RoomStatus/proto/v2"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/peer"
)

// RoomStream :
func (b *RoomStatusBackend) RoomStream(stream pb.RoomStatus_RoomStreamServer) error {
	peer, _ := peer.FromContext(stream.Context())
	log.Printf("Received new connection.  %s", peer.Addr.String())
	// md, _ := metadata.FromIncomingContext(stream.Context())
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	_, err = new_conn_reg(b, stream, req)
	if err != nil {
		return err
	}

	return nil
}

func new_conn_reg(b *RoomStatusBackend, stream pb.RoomStatus_RoomStreamServer, req *pb.CellStatusReq) (bool, error) {
	for _, v := range b.Roomlist {
		if v.conn_pool.Get(req.UserId) != nil {
			// Conn exist
			stream.Send(&pb.CellStatusResp{
				UserId:    "RmSvrMgr",
				Key:       v.Key,
				Status:    300,
				Timestamp: time.Now().String(),
				ResponseMsg: &pb.CellStatusResp_ErrorMsg{
					ErrorMsg: &pb.ErrorMsg{
						MsgInfo: "ConnIsExist",
						MsgDesp: "Connect is exist in Room<" + v.Key + ">",
					},
				},
			})
			return false, errors.New("ConnIsExist")
		}
		if v.Room.Key == req.Key {
			fmt.Println("new conn in Rm<" + v.Key + ">,UserId<" + req.UserId + ">")
			v.conn_pool.Add(req.UserId, stream)
			stream.Send(&pb.CellStatusResp{
				UserId:    "RmSvrMgr",
				Key:       v.Key,
				Status:    200,
				Timestamp: time.Now().String(),
				ResponseMsg: &pb.CellStatusResp_ErrorMsg{
					ErrorMsg: &pb.ErrorMsg{
						MsgInfo: "ConnSuccess",
						MsgDesp: "Connect to Room<" + v.Key + "> Success",
					},
				},
			})
			go func() {
				<-stream.Context().Done()
				v.conn_pool.Del(req.UserId)
				v.conn_pool.BroadCast(req.UserId,
					&pb.CellStatusResp{
						UserId:    "RmSvrMgr",
						Key:       v.Key,
						Status:    201,
						Timestamp: time.Now().String(),
						ResponseMsg: &pb.CellStatusResp_ErrorMsg{
							ErrorMsg: &pb.ErrorMsg{
								MsgInfo: "ConnEnd",
								MsgDesp: fmt.Sprintf(
									"User<%v> End to Room<%v>",
									req.UserId, v.Key),
							},
						},
					})
			}()
			for {
				req, err := stream.Recv()
				if err != nil {
					return false, err
				}
				ghj, err := update_proc(v, req)
				if err != nil {
					return false, err
				}
				v.conn_pool.BroadCast(req.UserId, ghj)
			}
			return true, nil
		}
	}
	return false, errors.New("RoomNotExist")
}

func update_proc(rmg *RoomMgr, req *pb.CellStatusReq) (*pb.CellStatusResp, error) {

	keynum := -1
	if len(rmg.Room.CellStatus) == 9 && rmg.Room.Round == 9 {
		log.Println("the game should be end")
		return nil, errors.New("GameEnd")
	}

	reqRoom := req.GetCellStatus()
	if reqRoom == nil {
		return nil, errors.New("UnknownCellStatus")
	}

	for k, v := range rmg.Room.CellStatus {
		if v.Turn == reqRoom.Turn {
			rmg.Room.Cell = int32(k)
			v.CellNum = reqRoom.CellNum
			keynum = k
			break
		}
	}

	if keynum == -1 {
		rmg.CellStatus = append(rmg.CellStatus, req.GetCellStatus())
		rmg.Cell = int32(len(rmg.CellStatus))
		rmg.Round++
	}
	return &pb.CellStatusResp{
		UserId:    req.UserId,
		Key:       rmg.Key,
		Timestamp: time.Now().String(),
		Status:    200,
		ResponseMsg: &pb.CellStatusResp_CellStatus{
			CellStatus: reqRoom,
		},
	}, nil
}
