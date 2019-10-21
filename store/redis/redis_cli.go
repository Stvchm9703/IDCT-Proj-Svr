package redis

import (
	"RoomStatus/config"
	"encoding/json"
	"log"
	"strconv"

	"github.com/go-redis/redis"
	// rejson "github.com/nitishm/go-rejson"
)

// 	consider the go-redis-client :
// 	key : <core-master-key>/_<redis-worker-hash-key>
// 	value : [] Room  {
// 		Room-Obj-props
//	}

// RdsCliBox : Redis client box custom interface
type RdsCliBox struct {
	conn    *redis.Client
	CoreKey string
	Key     string // redis-worker-cli
}

const (
	redisCliPoolName = "grpc-redis-cli-pool"
	redisCliSetTime  = 0
)

// Connect : Constructor of Redis client
func (rc *RdsCliBox) Connect(cf *config.ConfTmp) (bool, error) {
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
	// unregister
	if _, err := rc.unregister(); err != nil {
		return false, err
	}

	if err := rc.conn.Close(); err != nil {
		return false, err
	}
	return true, nil
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
	// return nil, nil
}

// SetPara : set the key-value
func (rc *RdsCliBox) SetPara(key *string, value interface{}) (bool, error) {
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
	res, err := rc.conn.Del(rc.CoreKey + "/_" + rc.Key + "." + *key).Result()
	if err != nil {
		return false, err
	}
	log.Println("res", res)
	return true, nil
}

// CleanRem : clear all this redis-cli rem
func (rc *RdsCliBox) CleanRem() (bool, error) {
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
func (rc *RdsCliBox) ListRem() (*[]string, error) {
	list, err := rc.conn.Keys(rc.CoreKey + "/_" + rc.Key + ".*").Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}
