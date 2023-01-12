package service

import (
	"bib-wss-test/client"
	"bib-wss-test/common"
	"bib-wss-test/message"
	"bib-wss-test/myUtils"
	"log"
	"strconv"
	"time"
)

func OrderPush(h *client.Hub) {
	clomns := myUtils.GetColumns("./bib_jean_apiKey_demo.csv")
	uids := clomns.Column0[1:]

	i := 0
	for i < len(uids) {
		//构建参数
		uidStr := uids[i]
		uid, err := strconv.Atoi(uidStr)
		if err != nil {
			log.Println("uid str to int fail:", err)
		}
		topics := []string{common.OrdersTopic, common.TradeTopic, common.AccountsTopic}

		//循环发送消息
		for _, topicType := range topics {
			msgData := message.GetOrderPushMsg(topicType, uid)
			h.SendChan <- msgData
		}

		//循环控制
		i++
		if i == len(uids) {
			i = 0
		}

		time.Sleep(1 * time.Second)
	}
}
