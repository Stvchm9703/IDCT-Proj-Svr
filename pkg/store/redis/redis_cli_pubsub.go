package redis

import (
	"errors"

	"github.com/go-redis/redis"
	// rejson "github.com/nitishm/go-rejson"
)

type RdsPubSub struct {
	conn        *redis.Client
	pscli       *redis.PubSub
	CoreKey     string
	Key         string // redis-worker-cli
	mountedFunc []*(func(...interface{}))
	structMode  bool
	keyWord     string
}

func CBSCopyFromClient(rds *RdsCliBox) *RdsPubSub {

	n := RdsPubSub{
		conn:       redis.NewClient(rds.options()),
		pscli:      nil,
		CoreKey:    rds.CoreKey,
		Key:        rds.Key,
		structMode: false,
		keyWord:    "",
	}
	n.conn.ConfigSet("notify-keyspace-events", "AKE").Result()
	return &n
}
func (r *RdsPubSub) GetStruct() *bool {
	return &r.structMode
}
func (r *RdsPubSub) SetStruct(s bool) {
	r.structMode = s
}

func (sb *RdsPubSub) AddChannel(title *string) (<-chan *redis.Message, error) {
	if sb.conn == nil {
		return nil, errors.New("conn cli is not created")
	}
	if sb.pscli != nil {
		return nil, errors.New("pub/sub cli is already existed")
	}

	sb.keyWord = *title
	if sb.structMode == true {
		sb.pscli = sb.conn.Subscribe(*title)
	} else {
		sb.pscli = sb.conn.PSubscribe(*title)
	}

	if _, err := sb.pscli.Receive(); err != nil {
		return nil, err
	}

	return sb.pscli.Channel(), nil
}

func (sb *RdsPubSub) CloseChannel() error {
	return sb.pscli.Close()
}

func (sb *RdsPubSub) Disconn() (bool, error) {
	if e := sb.CloseChannel(); e != nil {
		return false, e
	}
	if e := sb.conn.Close(); e != nil {
		return false, e
	}
	return true, nil
}

// pubsub := rdb.Subscribe("mychannel1")

// Wait for confirmation that subscription is created before publishing anything.
// _, err := pubsub.Receive()
// if err != nil {
//     panic(err)
// }

// // Go channel which receives messages.
// ch := pubsub.Channel()

// // Publish a message.
// err = rdb.Publish("mychannel1", "hello").Err()
// if err != nil {
//     panic(err)
// }

// time.AfterFunc(time.Second, func() {
//     // When pubsub is closed channel is closed too.
//     _ = pubsub.Close()
// })

// // Consume messages.
// for msg := range ch {
//     fmt.Println(msg.Channel, msg.Payload)
// }
