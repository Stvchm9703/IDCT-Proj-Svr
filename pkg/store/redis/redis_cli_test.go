package redis_test

import (
	cf "RoomStatus/config"
	pb "RoomStatus/proto"
	rd "RoomStatus/pkg/store/redis"
	"encoding/json"
	"log"
	"testing"
	"time"
)

var test_conf = cf.ConfTmp{
	cf.CfTemplServer{},
	cf.CfAPIServer{},
	cf.CfTDatabase{
		Connector:  "redis",
		WorkerNode: 12,
		Host:       "192.168.0.110",
		Port:       6379,
		Username:   "",
		Password:   "",
		Database:   "redis",
		Filepath:   "",
	},
}

func TestConnect(t *testing.T) {
	// test := rd.RdsCliBox{
	// 	CoreKey: "testingUnit",
	// 	Key:     "workerU",
	// }
	test := rd.New("testingUnit", "workerU")
	if _, err := test.Connect(&test_conf); err != nil {
		t.Log(err)
		log.Println(err)
	}
	t.Log(test)
	time.Sleep(5000)
	if _, err := test.Disconn(); err != nil {
		t.Log(err)
	}
}

var testObj = pb.Room{
	Key:        "Rm0122212",
	HostId:     "192.180",
	DuelerId:   "123.125",
	Status:     0,
	Round:      1,
	Cell:       1,
	CellStatus: nil,
}

func TestParaSet(t *testing.T) {
	test := rd.New("testingUnit", "workerU")
	if _, err := test.Connect(&test_conf); err != nil {
		t.Log(err)
	}

	time.Sleep(120000)

	res, err := test.SetPara(&testObj.Key, testObj)
	if err != nil {
		t.Log("err:", err)
	}
	t.Log("res:", res)
	log.Println(res)

	time.Sleep(120000)

	if _, err := test.Disconn(); err != nil {
		t.Log(err)
	}
}

func TestParaGet(t *testing.T) {
	t.Log("running ", t.Name())
	test := rd.New("testingUnit", "workerU")
	if _, err := test.Connect(&test_conf); err != nil {
		t.Log(err)
	}

	time.Sleep(120000)
	var fres pb.Room
	res, err := test.GetPara(&testObj.Key, &fres)
	if err != nil {
		t.Log("err:", err)
	}
	log.Println("res:", *res)
	log.Println("fres:", fres)
	jsoon, _ := json.Marshal(fres)
	log.Println(string(jsoon))
	time.Sleep(120000)

	if _, err := test.Disconn(); err != nil {
		t.Log(err)
	}
}
func TestParaRemove(t *testing.T) {
	test := rd.New("testingUnit", "workerU")
	if _, err := test.Connect(&test_conf); err != nil {
		t.Log(err)
	}

	time.Sleep(120000)
	res, err := test.RemovePara(&testObj.Key)
	if err != nil {
		t.Log("err:", err)
	}
	t.Log("res:", res)

	time.Sleep(120000)

	if _, err := test.Disconn(); err != nil {
		t.Log(err)
	}
}
func TestListKey(t *testing.T) {
	log.Println("running ", t.Name())
	t.Log("running ", t.Name())
	test := rd.New("testingUnit", "workerU")
	if _, err := test.Connect(&test_conf); err != nil {
		t.Fatal(err) //expected
	}
	ls, _ := test.ListRem()
	log.Print(*ls)
	time.Sleep(5000)
	if _, err := test.Disconn(); err != nil {
		t.Log(err)
	}
	log.Println("end ", t.Name())
	t.Log("end ", t.Name())
}

func TestCleanRem(t *testing.T) {
	log.Println("running ", t.Name())
	t.Log("running ", t.Name())
	// test := rd.New("testingUnit", "workerU")
	test := rd.New("RSCore8a3d11e", "wKU")
	if _, err := test.Connect(&test_conf); err != nil {
		t.Fatal(err) //expected
	}

	// // insert testing data
	// for i := 0; i < 50; i++ {
	// 	var testObjLoop = pb.Room{
	// 		Key:        "Rm" + cm.HashText("num"+strconv.Itoa(i)+"test"),
	// 		HostId:     "192.180.0." + strconv.Itoa(i),
	// 		Status:     0,
	// 		Round:      1,
	// 		Cell:       1,
	// 		CellStatus: nil,
	// 	}
	// 	if _, err := test.SetPara(&testObjLoop.Key, testObjLoop); err != nil {
	// 		t.Log("err:", err)
	// 	}
	// 	time.Sleep(5000)
	// }
	// time.Sleep(500000)
	// // show the list
	// ls, _ := test.ListRem()
	// log.Print(*ls)
	// time.Sleep(500000)

	// testing clean all
	if _, err := test.CleanRem(); err != nil {
		t.Fatal(err)
		log.Fatal(err)
	}

	if _, err := test.Disconn(); err != nil {
		t.Log(err)
	}
	log.Println("end ", t.Name())
	t.Log("end ", t.Name())
}

func TestForceClear(t *testing.T) {
	log.Println("running ", t.Name())
	t.Log("running ", t.Name())
	// test := rd.New("testingUnit", "workerU")
	test := rd.New("RSCore8a3d11e", "wKU")
	if _, err := test.Connect(&test_conf); err != nil {
		t.Fatal(err) //expected
	}

	// testing clean all
	if _, err := test.CleanRem(); err != nil {
		t.Fatal(err)
		log.Fatal(err)
	}
	if _, err := test.ForceClear(); err != nil {
		t.Fatal(err)
		log.Fatal(err)
	}

	if _, err := test.Disconn(); err != nil {
		t.Log(err)
	}
	log.Println("end ", t.Name())
	t.Log("end ", t.Name())
}
