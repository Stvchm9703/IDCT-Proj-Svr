package redis

import (
	"RoomStatus/config"
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
	// rejson "github.com/nitishm/go-rejson"
)

// 	consider the go-redis-client :
// 	key : <core-master-key>/_<redis-worker-hash-key>
// 	value : [] Room  {
// 		Room-Obj-props
//	}

// Remark : Export Main function is need to add
// 			Lock / Unlock for sync

// RdsCliBox : Redis client box custom interface
type RdsCliBox struct {
	conn      *redis.Client
	CoreKey   string
	Key       string // redis-worker-cli
	isRunning bool
	mu        *sync.Mutex
}

const (
	redisCliPoolName = "grpc-redis-cli-pool"
	redisCliSetTime  = 0
)

func (rc *RdsCliBox) IsRunning() *bool { return &rc.isRunning }

func (rc *RdsCliBox) lock() {
	rc.mu.Lock()
	// rc.isRunning = true
}

func (rc *RdsCliBox) unlock() {
	rc.mu.Unlock()
	// rc.isRunning = false
}

func (rc *RdsCliBox) Preserve(s bool) {
	rc.isRunning = s
}

// Connect : Constructor of Redis client
func (rc *RdsCliBox) Connect(cf *config.ConfTmp) (bool, error) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	rc.conn = redis.NewClient(&redis.Options{
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
	rc.mu.Lock()
	defer rc.mu.Unlock()
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
	rc.mu.Lock()
	defer rc.mu.Unlock()
	optionBu := rc.conn.Options()
	if err := rc.conn.Close(); err != nil {
		return nil, err
	}
	time.Sleep(50000)
	log.Println("re-create the redis client")
	newConn := redis.NewClient(optionBu)
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
		log.Println("ind:", ind)
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
	log.Println("register-proc:", res)
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
		log.Println("ind:", ind)
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

	res, err := rc.conn.LRem(redisCliPoolName, -1, str).Result()
	if err != nil {
		return false, err
	}
	log.Println("unreg-proc:", res)
	return true, nil
}

// alive :

//

// GetPara : get the value by key
func (rc *RdsCliBox) GetPara(key *string, target interface{}) (*interface{}, error) {
	rc.mu.lock()
	defer rc.mu.unlock()
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
	rc.mu.lock()
	defer rc.mu.unlock()
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
	rc.mu.lock()
	defer rc.mu.unlock()
	res, err := rc.conn.Del(rc.CoreKey + "/_" + rc.Key + "." + *key).Result()
	if err != nil {
		return false, err
	}
	log.Println("res", res)
	return true, nil
}

// CleanRem : clear all this redis-cli rem
func (rc *RdsCliBox) CleanRem() (bool, error) {
	rc.mu.lock()
	defer rc.mu.unlock()
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
	rc.mu.lock()
	defer rc.mu.unlock()
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
func (rc *RdsCliBox) GetParaList(key *string, target interface{}, refPointer interface{}) (*interface{}, error) {
	rc.mu.lock()
	defer rc.mu.unlock()
	keystr := rc.CoreKey + "/_" + rc.Key + "." + *key
	res, err := rc.conn.MGet(keystr).Result()
	if err != nil {
		return nil, err
	}
	log.Println(res)
	// arrKey := []string
	for _, v := range res {
		log.Println(v)
		log.Println(reflect.TypeOf(v))
		resstr, err := strconv.Unquote(v.(string))
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal([]byte(resstr), &refPointer); err != nil {
			return nil, err
		}
		tmp := (refPointer)
		target = append(target.([]interface{}), tmp)
	}
	refPointer = nil
	return &target, nil
}
