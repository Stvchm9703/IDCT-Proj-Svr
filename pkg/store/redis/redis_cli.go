package redis

import (
	"RoomStatus/config"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	rd "github.com/go-redis/redis"
	"github.com/micro/go-micro/errors"
	// rejson "github.com/nitishm/go-rejson"
)

// var t :=

// 	consider the go-redis-client :
// 	key : <core-master-key>/_<redis-worker-hash-key>
// 	value : [] Room  {
// 		Room-Obj-props
//	}

// Remark : Export Main function is need to add
// 			Lock / Unlock for sync

// RdsCliBox : Redis client box custom interface
type RdsCliBox struct {
	conn      *rd.Client
	CoreKey   string
	Key       string // redis-worker-cli
	isRunning bool
	mu        *sync.Mutex
}

const (
	redisCliPoolName = "grpc-redis-cli-pool"
	redisCliSetTime  = 0
)

func New(coreKey string, key string) *RdsCliBox {
	v := sync.Mutex{}
	return &RdsCliBox{
		CoreKey: coreKey,
		Key:     key,
		mu:      &v,
	}
}

func (rc *RdsCliBox) IsRunning() *bool { return &rc.isRunning }

func (rc *RdsCliBox) lock() {
	if rc.mu != nil {
		rc.mu.Lock()
	}
	// rc.isRunning = true
}

func (rc *RdsCliBox) unlock() {
	if rc.mu != nil {
		rc.mu.Unlock()
	}
	// rc.isRunning = false
}

func (rc *RdsCliBox) Preserve(s bool) {
	rc.isRunning = s
}

func (rc *RdsCliBox) options() *rd.Options {
	return rc.conn.Options()
}

// Connect : Constructor of Redis client
func (rc *RdsCliBox) Connect(cf *config.ConfTmp) (bool, error) {
	rc.lock()
	defer rc.unlock()

	rc.conn = rd.NewClient(&rd.Options{
		Addr:     cf.Database.Host + ":" + strconv.Itoa(cf.Database.Port),
		Password: cf.Database.Password,
		PoolSize: cf.Database.WorkerNode,
	})
	// try ping conn
	_, err := rc.conn.Ping().Result()
	if err != nil {
		return false, err
	}

	if _, err = rc.register(); err != nil {
		log.Println("hi form outside of register")
		return false, err
	}

	return true, nil
}

// Disconn : notice redis server to kill process, Gratefully;;
func (rc *RdsCliBox) Disconn() (bool, error) {
	// 	Note: Clean up , it is suggested to clean Rem manually
	// if _, err := rc.CleanRem(); err != nil {
	// 	return false, err
	// }
	rc.lock()
	defer rc.unlock()
	// unregister
	if _, err := rc.unregister(); err != nil {
		return false, err
	}

	if err := rc.conn.Close(); err != nil {
		return false, err
	}
	return true, nil
}

// Recover :
func (rc *RdsCliBox) Recover() (*RdsCliBox, error) {
	rc.lock()
	defer rc.unlock()
	optionBu := rc.conn.Options()
	if err := rc.conn.Close(); err != nil {
		return nil, err
	}
	time.Sleep(50000)
	log.Println("re-create the redis client")
	newConn := rd.NewClient(optionBu)
	log.Println("try ping")
	_, err := newConn.Ping().Result()
	if err != nil {
		return nil, err
	}
	rc.conn = newConn
	return rc, nil
}

// register : push self working id into temp pool
func (rc *RdsCliBox) register() (bool, error) {
	str := rc.CoreKey + "/_" + rc.Key
	ind, err := rc.conn.LRange(redisCliPoolName, 0, -1).Result()
	if err != nil {
		log.Println("error search")
		log.Println("ind:", ind)
		log.Println(err)
		keyexist, err := rc.conn.Exists(redisCliPoolName).Result()
		if err != nil {
			return false, err
		} else if keyexist == 0 {
			// pass
		}
	} else {
		// log.Println("ind:", ind)
		for _, v := range ind {
			if v == str {
				log.Println("key exist")
				return false, nil
			}
		}
		// not exist in list
		// pass
	}
	res, err := rc.conn.RPush(redisCliPoolName, str).Result()
	if err != nil {
		return false, err
	}
	log.Println("register-proc:", res, ":", str)
	return true, nil
}

// unregister
func (rc *RdsCliBox) unregister() (bool, error) {
	str := rc.CoreKey + "/_" + rc.Key
	ind, err := rc.conn.LRange(redisCliPoolName, 0, -1).Result()
	if err != nil {
		log.Println("error search")
		log.Println("ind:", ind)
		log.Println(err)
		keyexist, err := rc.conn.Exists(redisCliPoolName).Result()
		if err != nil {
			return false, err
		} else if keyexist == 0 {
			// pass
		}
	} else {
		// log.Println("ind:", ind)
		cd := len(ind)
		for _, v := range ind {
			if v == str {
				break
			} else {
				cd--
			}
		}
		if cd == 0 {
			return false, nil
		}
	}

	_, err = rc.conn.LRem(redisCliPoolName, -1, str).Result()
	if err != nil {
		return false, err
	}
	log.Println("unreg-proc:", str)
	return true, nil
}

// alive :

// ForceClear
func (rc *RdsCliBox) ForceClear() (bool, error) {
	// str := rc.CoreKey + "/_*"
	ind, err := rc.conn.LRange(redisCliPoolName, 0, -1).Result()
	if err != nil {
		log.Println("error search")
		log.Println("ind:", ind)
		log.Println(err)
		keyexist, err := rc.conn.Exists(redisCliPoolName).Result()
		if err != nil {
			return false, err
		} else if keyexist == 0 {
			// pass
		}
	} else {
		log.Println("ind:", ind)

	}
	for i := 0; i < len(ind); i++ {
		if ind[i] != "_index_content_" {
			res, err := rc.conn.LRem(redisCliPoolName, -1, ind[i]).Result()
			if err != nil {
				return false, err
			}
			log.Println("unreg-proc:", res)
		}
	}
	// res, err := rc.conn.LRem(redisCliPoolName, -1, rc.CoreKey+"/_*").Result()
	// if err != nil {
	// 	return false, err
	// }
	// log.Println("unreg-proc:", res)
	return true, nil
}

// GetPara : get the value by key
func (rc *RdsCliBox) GetPara(key *string, target interface{}) (*interface{}, error) {
	rc.lock()
	defer rc.unlock()
	keystr := rc.CoreKey + "/_" + rc.Key + "." + *key
	res, err := rc.conn.Get(keystr).Result()
	if err != nil {
		return nil, err
	}
	resstr, err := strconv.Unquote(res)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(resstr), &target); err != nil {
		return nil, err
	}
	return &target, nil
}

// SetPara : set the key-value
func (rc *RdsCliBox) SetPara(key *string, value interface{}) (bool, error) {
	rc.lock()
	defer rc.unlock()
	keystr := rc.CoreKey + "/_" + rc.Key + "." + *key
	jsonFormat, err := json.Marshal(value)
	if err != nil {
		return false, err
	}
	strr := strconv.Quote(string(jsonFormat))

	if _, err := rc.conn.Set(keystr, strr, redisCliSetTime).Result(); err != nil {
		return false, err
	}
	return true, nil
}

// RemovePara : remove the k-v
func (rc *RdsCliBox) RemovePara(key *string) (bool, error) {
	rc.lock()
	defer rc.unlock()
	res, err := rc.conn.Del(rc.CoreKey + "/_" + rc.Key + "." + *key).Result()
	if err != nil {
		return false, err
	}
	log.Println("res", res)
	return true, nil
}

// UpdatePara : set the key-value
func (rc *RdsCliBox) UpdatePara(key *string, value interface{}) (bool, error) {
	rc.lock()
	defer rc.unlock()
	keystr := "*/_*." + *key
	keys, err := rc.conn.Keys(keystr).Result()
	if err != nil {
		log.Println("get-para-process,conn-keys,err", err)
		return false, err
	}
	if len(keys) != 1 {
		// not ok
		return false, errors.New("UPDATE_ERROR", "more than one keys:["+strings.Join(keys, ",")+"]", 1)
	}

	// OK case
	jsonFormat, err := json.Marshal(value)
	if err != nil {
		return false, err
	}
	strr := strconv.Quote(string(jsonFormat))

	if _, err := rc.conn.Set(keys[0], strr, redisCliSetTime).Result(); err != nil {
		return false, err
	}
	return true, nil
}

// CleanRem : clear all this redis-cli rem
func (rc *RdsCliBox) CleanRem() (bool, error) {
	// rc.lock()
	// defer rc.unlock()
	list, err := rc.ListRem()
	if err != nil {
		return false, nil
	}
	for _, v := range *list {
		if _, err := rc.conn.Del(v).Result(); err != nil {
			return false, err
		}
	}
	return true, nil
}

// ListRem : check the ha key
func (rc *RdsCliBox) ListRem(optionKey ...*string) (*[]string, error) {
	rc.lock()
	defer rc.unlock()
	var list []string
	var err error

	list, err = rc.conn.Keys(rc.CoreKey + "/_" + rc.Key + ".*").Result()
	if err != nil {
		return nil, err
	}
	if len(optionKey) > 0 {
		var listy []string
		for _, v := range list {
			for _, lv := range optionKey {
				if strings.Contains(v, *lv) {
					listy = append(listy, v)
				}
			}
		}
		list = listy
	}
	return &list, nil
}

/// NOTE: Need add testing

// GetParaList : get a list of feature Para
func (rc *RdsCliBox) GetParaList(key *string) (*[]byte, error) {
	rc.lock()
	defer rc.unlock()
	keystr := rc.CoreKey + "/_" + rc.Key + ".*" + *key + "*"
	log.Println("Key-in:", keystr)
	keys, err := rc.conn.Keys(keystr).Result()
	if err != nil {
		log.Println("get-para-process,conn-keys,err", err)
		return nil, err
	}
	log.Println("key-find:", keys)
	var res []interface{}
	if len(keys) > 0 {
		res, err = rc.conn.MGet(keys...).Result()
		if err != nil {
			log.Println("get-para-process-err", err)
			return nil, err
		}
	}

	resStr := []byte(`[`)
	for k, v := range res {
		if v != nil {
			tmp, err := strconv.Unquote(v.(string))
			if err != nil {
				return nil, err
			}
			resStr = append(resStr, []byte(tmp)...)
			if k != len(res)-1 {
				resStr = append(resStr, []byte(",")...)
			}
		}
	}
	// refPointer = nil
	resStr = append(resStr, []byte(`]`)...)
	return &resStr, nil
}
