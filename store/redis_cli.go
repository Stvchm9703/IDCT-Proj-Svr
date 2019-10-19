package store

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
)

// 	consider the go-redis-client :
// 	key : <core-master-key>/_<redis-worker-hash-key>
// 	value : [] Room  {
// 		Room-Obj-props
//	}

// RedisCli : Redis client box custom interface
type RedisCli struct {
	conn *redis.Client
	rjh  *rejson.Handler
	key  *string // redis-worker-cli
}

// Connect : Constructor of Redis client
func (rc *RedisCli) Connect(conf) (bool, error) {
	rc.conn = redis.NewClient(&redis.Options{})
	// try ping conn
	_, err := rc.conn.Ping().Result()
	if err != nil {
		return false, err
	}
	rc.rjh = rejson.NewReJSONHandler()
	rc.rjh.SetGoRedisClient(rc.conn)
	return true, nil
}

// Disconn : notice redis server to kill process, Gratefully;;
func (rc *RedisCli) Disconn() (bool, error) {
	// Clean up
	if _, err := rc.CleanRem(); err != nil {
		return false, err
	}

	if err := rc.conn.Close(); err != nil {
		return false, err
	}
	return true, nil
}

// GetPara : get the value by key
func (rc *RedisCli) GetPara(key *string, target *interface{}) (*interface{}, error) {
	keystr := *rc.key + "/_" + *key
	res, err := rc.rjh.JSONGet(keystr, ".")
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(res.([]byte), &target); err != nil {
		return &res, err
	}
	return target, nil
	// return nil, nil
}

// SetPara : set the key-value
func (rc *RedisCli) SetPara(key *string, value *interface{}) (bool, error) {
	keystr := *rc.key + "/_" + *key
	res, err := rc.rjh.JSONSet(keystr, ".", value)
	if err != nil {
		return false, err
	}
	log.Println(res)
	return true, nil
}

// RemovePara : remove the k-v
func (rc *RedisCli) RemovePara(key *string) (bool, error) {
	keystr := *rc.key + "/_" + *key
	return true, nil
}

// CleanRem : clear all this redis-cli rem
func (rc *RedisCli) CleanRem() (bool, error) {
	res, err := rc.rjh.JSONDel(*rc.key+"/_*", ".")
	if err != nil {
		return false, err
	}
	log.Println("res", res)
	return true, nil
}
