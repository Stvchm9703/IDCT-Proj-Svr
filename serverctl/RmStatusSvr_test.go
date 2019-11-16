package serverctl_test

// import (
// 	cf "RoomStatus/config"
// 	pb "RoomStatus/proto"
// 	server "RoomStatus/serverGameCtl"
// 	"encoding/json"
// 	"log"
// 	"strconv"
// 	"testing"
// )

// var testing_config = cf.ConfTmp{
// 	cf.CfTemplServer{},
// 	cf.CfAPIServer{
// 		ConnType:     "TCP",
// 		IP:           "127.0.0.1",
// 		Port:         11000,
// 		MaxPoolSize:  20,
// 		APIReferType: "proto",
// 		APITablePath: "{root}/thrid_party/OpenAPI",
// 		APIOutpath:   "./",
// 	},
// 	cf.CfTDatabase{
// 		Connector:  "redis",
// 		WorkerNode: 12,
// 		Host:       "192.168.0.110",
// 		Port:       6379,
// 		Username:   "",
// 		Password:   "",
// 		Database:   "redis",
// 		Filepath:   "",
// 	},
// }

// func TestInitServer(t *testing.T) {
// 	b := server.New(&testing_config)
// 	t.Log("Server Opened")
// 	t.Log(b)
// 	log.Println("Server Opened")
// 	log.Println(b)
// 	j, _ := json.Marshal(b)
// 	t.Log(string(j))
// 	log.Println(string(j))
// }

// func TestCreateWkTask(t *testing.T) {
// 	t.Log("Init Server")
// 	b := server.New(&testing_config)
// 	pl := server.WkTask{
// 		In: &pb.RoomCreateRequest{
// 			HostId: "123.451.1AA",
// 		},
// 		Out: make(chan interface{})}
// 	Tfunc, err := b.TestCreateWkTask(pl)
// 	log.Println("func copy")
// 	log.Println(Tfunc)
// 	log.Println("RoomList", b.Roomlist)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	b.Shutdown()
// }

// func BenchmarkCreateWkTask(t *testing.B) {

// 	b := server.New(&testing_config)
// 	t.ResetTimer()
// 	for i := 0; i < t.N; i++ {
// 		pl := server.WkTask{
// 			In: &pb.RoomCreateRequest{
// 				HostId: "123.451.1AA" + strconv.Itoa(i),
// 			},
// 			Out: make(chan interface{})}
// 		_, _ = b.TestCreateWkTask(pl)

// 	}
// 	t.StopTimer()
// 	b.Shutdown()
// }

// func TestDeleteWkTask(t *testing.T) {
// 	t.Log("Init Server")
// 	b := server.New(&testing_config)

// 	// Create Room
// 	pl := server.WkTask{
// 		In: &pb.RoomCreateRequest{
// 			HostId: "123.451.1AA",
// 		},
// 		Out: make(chan interface{})}
// 	log.Println("Create req")
// 	Tfunc, err := b.TestCreateWkTask(pl)
// 	log.Println(Tfunc)
// 	log.Println("RoomList", b.Roomlist)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	// Detele Room
// 	l := server.WkTask{
// 		In: &pb.RoomRequest{
// 			Key: b.Roomlist[0].Key,
// 		},
// 		Out: make(chan interface{}),
// 	}

// 	ff, er := b.TestDeteleWkTask(l)
// 	log.Println("Delete req")
// 	log.Println(ff)
// 	log.Println("Room List")
// 	log.Println(b.Roomlist)
// 	if er != nil {
// 		log.Println(er)
// 		b.Shutdown()
// 	}
// 	b.Shutdown()
// }

// func TestGetWkTask(t *testing.T) {
// 	t.Log("Init Server")
// 	b := server.New(&testing_config)

// 	// Create Room
// 	pl := server.WkTask{
// 		In: &pb.RoomCreateRequest{
// 			HostId: "123.451.1AA",
// 		},
// 		Out: make(chan interface{})}
// 	log.Println("Create req")
// 	Tfunc, err := b.TestCreateWkTask(pl)
// 	log.Println(Tfunc)
// 	log.Println("RoomList", b.Roomlist)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	// Get Room
// 	g := server.WkTask{
// 		In: &pb.RoomRequest{
// 			Key: b.Roomlist[0].Key,
// 		},
// 		Out: make(chan interface{}),
// 	}
// 	gg, err := b.TestGetInfoWkTask(g)
// 	log.Println("Get req")
// 	log.Println(gg)
// 	log.Println("Room List")
// 	log.Println(b.Roomlist)
// 	if err != nil {
// 		log.Println(err)
// 		b.Shutdown()
// 	}

// 	// Detele Room
// 	l := server.WkTask{
// 		In: &pb.RoomRequest{
// 			Key: b.Roomlist[0].Key,
// 		},
// 		Out: make(chan interface{}),
// 	}

// 	ff, er := b.TestDeteleWkTask(l)
// 	log.Println("Delete req")
// 	log.Println(ff)
// 	log.Println("Room List")
// 	log.Println(b.Roomlist)
// 	if er != nil {
// 		log.Println(er)
// 		b.Shutdown()
// 	}
// 	b.Shutdown()
// }

// func TestGetListPWkTask(t *testing.T) {
// 	t.Log("Init Server")
// 	b := server.New(&testing_config)

// 	// Create Room
// 	pl := server.WkTask{
// 		In: &pb.RoomCreateRequest{
// 			HostId: "123.451.1AA",
// 		},
// 		Out: make(chan interface{})}
// 	log.Println("Create req")
// 	for i := 0; i < 10; i++ {
// 		_, err := b.TestCreateWkTask(pl)
// 		// log.Println(Tfunc)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
// 	log.Println("RoomList", b.Roomlist)
// 	// Get Room
// 	g := server.WkTask{
// 		In: &pb.RoomListRequest{
// 			Requirement: "*",
// 		},
// 		Out: make(chan interface{}),
// 	}
// 	gg, err := b.TestGetLsWkTask(g)
// 	log.Println("Get req")
// 	log.Println(gg)
// 	log.Println("Room List")
// 	log.Println(b.Roomlist)
// 	log.Println("err", err)

// 	if err != nil {
// 		log.Println(err)
// 		b.Shutdown()
// 	}

// 	// Detele Room
// 	log.Println("Delete Rm")
// 	for i := 0; i < 10; i++ {
// 		log.Println(b.Roomlist[0].Key)
// 		l := server.WkTask{
// 			In: &pb.RoomRequest{
// 				Key: b.Roomlist[0].Key,
// 			},
// 			Out: make(chan interface{}),
// 		}
// 		_, er := b.TestDeteleWkTask(l)
// 		if er != nil {
// 			log.Println(er)
// 			b.Shutdown()
// 		}
// 	}
// 	log.Println(b.Roomlist)
// 	b.Shutdown()
// }

// func TestGetListWkTask(t *testing.T) {
// 	t.Log("Init Server")
// 	b := server.New(&testing_config)

// 	// Create Room
// 	pl := server.WkTask{
// 		In: &pb.RoomCreateRequest{
// 			HostId: "123.451.1AA",
// 		},
// 		Out: make(chan interface{})}
// 	log.Println("Create req")
// 	for i := 0; i < 10; i++ {
// 		_, err := b.TestCreateWkTask(pl)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
// 	log.Println("RoomList", b.Roomlist)
// 	// Get Room
// 	g := server.WkTask{
// 		In: &pb.RoomListRequest{
// 			Requirement: "*",
// 		},
// 		Out: make(chan interface{}),
// 	}
// 	gg, err := b.TestGetLsWkTask(g)
// 	log.Println("Get req")
// 	log.Println(gg)
// 	log.Println("Room List")
// 	log.Println(b.Roomlist)
// 	log.Println("err", err)

// 	if err != nil {
// 		log.Println(err)
// 		b.Shutdown()
// 	}

// 	// Detele Room
// 	log.Println("Delete Rm")
// 	for i := 0; i < 10; i++ {
// 		log.Println(b.Roomlist[0].Key)
// 		l := server.WkTask{
// 			In: &pb.RoomRequest{
// 				Key: b.Roomlist[0].Key,
// 			},
// 			Out: make(chan interface{}),
// 		}
// 		_, er := b.TestDeteleWkTask(l)
// 		if er != nil {
// 			log.Println(er)
// 			b.Shutdown()
// 		}
// 	}
// 	log.Println(b.Roomlist)
// 	b.Shutdown()
// }
