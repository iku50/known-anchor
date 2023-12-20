package mq

import "encoding/json"

type MessageType uint

type Message struct {
	From      string `json:"from"`
	Func      string `json:"func"`
	Content   string `json:"content"`
	TimeStamp int64  `json:"timeStamp"`
}

func toMessage(msg []byte) (message Message, err error) {
	err = json.Unmarshal(msg, &message)
	return
}

func (m *Message) toBytes() (msg []byte, err error) {
	msg, err = json.Marshal(m)
	return
}
