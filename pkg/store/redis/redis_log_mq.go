// Package redis ://redis_log_mq.go
// Remark : for future usage , Redis MQ framework
// https://godoc.org/github.com/go-redis/redis#PubSub
package redis

import (
	"RoomStatus/config"
	"log"

	"github.com/go-redis/redis"
	fb "github.com/google/flatbuffers/go"
)

type MQProducer struct {
	ReferConf  *config.ConfTmp
	conn       *redis.Client
	ChannelKey map[string]MQFeedback
}

type MQProducerInf interface {
	Init()
	Send(channel_key *string, msg interface{}, option ...string) (bool, error)
	sendJson(channel_key *string, msg interface{}) (bool, error)
	// FlatBuffer
	sendFB(channelKey *string, msg interface{}) (bool, error)
}

type MQFeedback struct {
	OnError *func()
	OnAck   *func()
	OnSucc  *func()
	OnComp  *func()
}

func (mqp *MQProducer) sendFB(channelKey *string, msg interface{}, feedback ...MQFeedback) (bool, error) {
	Fbc := fb.FlatbuffersCodec{}
	msgbf, err := Fbc.Marshal(msg)
	log.Println(msgbf)
	mqp.ChannelKey[*channelKey] = feedback[0]

	return false, err
}
