package client

import (
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"log"
)

type Hub struct {
	//消息发送通道
	SendChan chan *primitive.Message
	//消息返回通道
	RevChan chan string
}

func NewHub() *Hub {
	return &Hub{
		SendChan: make(chan *primitive.Message, 256),
		RevChan:  make(chan string, 256),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case msg := <-h.RevChan:
			log.Println(msg)
		}
	}
}
