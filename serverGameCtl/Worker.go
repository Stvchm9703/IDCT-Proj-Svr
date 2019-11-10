package serverGameCtrl

// https://talks.golang.org/2015/gotham-grpc.slide#21
import (
	cm "RoomStatus/common"
	pb "RoomStatus/proto"
	rd "RoomStatus/store/redis"
	"errors"
	"log"
	"sync"
	"time"
)

type task struct {
	id  int
	do  func() (interface{}, error)
	arg interface{}
}

type GCExecutor struct {
	workerId     string
	maxWorker    int
	ActiveWorker int
	rdsbox       *rd.RdsCliBox
	RoomInTask   chan task
	RoomInWatch  []*pb.Room
	RoomWatch    sync.Mutex
	is_running   bool
}

func CreateWorker(workerId *string, maxWorker *int, assignRds *rd.RdsCliBox) *GCExecutor {
	return &GCExecutor{
		workerId:     *workerId,
		maxWorker:    *maxWorker,
		ActiveWorker: 0,
		rdsbox:       assignRds,
		is_running:   false,
	}
}

func (gce *GCExecutor) StartOperation() <-chan int {
	for gce.is_running {
		h := <-gce.RoomInTask
		cas, err := h.do()
	}
}

func (gce *GCExecutor) PalseOperation() {}

func (gce *GCExecutor) StopOperation() {}

// func (gce *GCExecutor) CreateTask(
// 	task_id *string,
// 	callfunc func() (interface{}, error),
// 	args ...interface{},
// ) (err error) {
// 	select {
// 	case gce.RoomInTask <- gce.rdsbox.SetPara(args[0], args[1]):
// 		gce.ActiveWorker++
// 		break
// 	default:
// 		err = errors.New("buffer full")
// 	}
// }

// ctCreateRoom : internal function for create task
func ctCreateRoom(taskId *int, exwk *GCExecutor, req *pb.RoomCreateRequest) *task {
	return &task{
		id:  *taskId,
		arg: req,
		do: func() (interface{}, error) {

			tmptime := time.Now().String() + req.HostId
			// for loop it
			var f = ""
			for {
				f = cm.HashText(tmptime)
				l, err := exwk.rdsbox.ListRem(&f)
				if err != nil {
					log.Fatal(err)
					return nil, err
				}
				if len(*l) == 0 {
					break
				}
			}

			rmTmp := pb.Room{
				Key:        "Rm" + f,
				HostId:     req.HostId,
				DuelerId:   "",
				Status:     0,
				Round:      0,
				Cell:       -1,
				CellStatus: nil,
			}
			if _, err := exwk.rdsbox.SetPara(&rmTmp.Key, rmTmp); err != nil {
				log.Fatalln(err)
				return &rmTmp, err
			}

			// create room success
			exwk.RoomInWatch = append(exwk.RoomInWatch, &rmTmp)
			return &rmTmp, nil
		},
	}
}

// CreateRoomTask : export function for implement GCExecutor
func (gce *GCExecutor) CreateRoomTask(
	task_id *int,
	args *pb.RoomCreateRequest,
) (err error) {
	tk := ctCreateRoom(task_id, gce, args)
	select {
	case gce.RoomInTask <- *tk:
		gce.ActiveWorker++
		break
	default:
		err = errors.New("buffer full")
	}
	return
}

// ctGetRoomList : internal function for create task
func ctGetRoomList(taskId *int, exwk *GCExecutor, req *pb.RoomListRequest) *task {
	return &task{
		id:  *taskId,
		arg: req,
		do: func() (interface{}, error) {
			var res pb.RoomListResponse
			var tmp pb.Room
			var RmList []pb.Room

			if _, err2 := exwk.rdsbox.GetParaList(&req.Requirement, &RmList, tmp); err2 != nil {
				log.Fatalln(err2)
				return nil, err2
			}
			for _, v := range RmList {
				res.Result = append(res.Result, &v)
			}
			return res, nil
		},
	}
}

// GetRoomListTask : exported function for GCE implement
func (gce *GCExecutor) GetRoomListTask(
	task_id *int,
	args *pb.RoomListRequest,
) (err error) {
	tk := ctGetRoomList(task_id, gce, args)
	select {
	case gce.RoomInTask <- *tk:
		gce.ActiveWorker++
		break
	default:
		err = errors.New("buffer full")
	}
	return
}

// ct
