package client

import (
	"bib-wss-test/common"
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"log"
	"os"
	"time"
)

func WritePump(h *Hub) {
	p, _ := rocketmq.NewProducer(
		// 设置  nameSrvAddr
		// nameSrvAddr 是 Topic 路由注册中心
		producer.WithNameServer(common.NameServerAddr),
		// 指定发送失败时的重试时间
		producer.WithRetry(2),
		// 设置 Group
		producer.WithGroupName(common.GroupName),
	)

	// 开始连接
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	defer func() {
		// 关闭生产者
		err = p.Shutdown()
		if err != nil {
			fmt.Printf("shutdown producer error:%s", err.Error())
		}
		close(h.SendChan)
		close(h.RevChan)
	}()

	// 循坏发送信息 (同步发送)
	stop := false
	for {
		select {
		//发送消息
		case msg := <-h.SendChan:
			result, err := p.SendSync(context.Background(), msg)
			if err != nil {
				log.Println("send msg err:", err)
			} else {
				h.RevChan <- result.String()
			}

		//超时控制
		case <-time.After(60 * time.Second):
			log.Println("timeout writePump stop!")
			stop = true
			break
		}
		if stop {
			break
		}
	}

}
