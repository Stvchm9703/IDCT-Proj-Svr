package authServer

import (
	// _ "RoomStatus"
	cm "RoomStatus/common"
	cf "RoomStatus/config"
	pb "RoomStatus/proto"
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
)

var _ pb.CreditsAuthServer = (*CreditsAuthBackend)(nil)

// Remark: the framework make consider "instant" request
//
type CreditsAuthBackend struct {
	mu      *sync.Mutex
	CoreKey string
	DBconn  *gorm.DB
}

// New : Create new backend
func New(conf *cf.ConfTmp) *CreditsAuthBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)

	g := CreditsAuthBackend{
		CoreKey: ck,
		mu:      &sync.Mutex{},
	}

	g.InitDB(&conf.Database)
	
	return &g
}

func (this *CreditsAuthBackend) Shutdown() {
	log.Println("in shtdown proc")
	this.CloseDB()
	log.Println("endof shutdown proc:", this.CoreKey)
}

// 	Impletement from GameCtl.pb.go(auto-gen file)
//		CheckCred(context.Context, *CredReq) (*CheckCredResp, error)
// 		GetCred(context.Context, *CredReq) (*CreateCredResp, error)
// 		CreateCred(*CredReq, CreditsAuth_CreateCredServer) error

// printReqLog
func printReqLog(ctx context.Context, req interface{}) {
	jsoon, _ := json.Marshal(ctx)
	log.Println(string(jsoon))

	jsoon, _ = json.Marshal(req)
	log.Println(string(jsoon))
}
