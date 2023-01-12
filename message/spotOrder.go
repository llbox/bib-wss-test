package message

import (
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

// 1、orders_ws_api
type orderMsg struct {
	AccountID       int    `json:"accountId"`
	Aggressor       bool   `json:"aggressor"`
	ClientOrderID   string `json:"clientOrderId"`
	EventType       string `json:"eventType"`
	OrderCreateTime int64  `json:"orderCreateTime"`
	OrderID         int64  `json:"orderId"`
	OrderPrice      string `json:"orderPrice"`
	OrderSide       string `json:"orderSide"`
	OrderSize       string `json:"orderSize"`
	OrderStatus     string `json:"orderStatus"`
	OrderType       string `json:"orderType"`
	Source          string `json:"source"`
	Symbol          string `json:"symbol"`
}

// 2、trade_ws_api
type tradeMsg struct {
	AccountID       int    `json:"accountId"`
	Aggressor       bool   `json:"aggressor"`
	ClientOrderID   string `json:"clientOrderId"`
	EventType       string `json:"eventType"`
	OrderCreateTime int64  `json:"orderCreateTime"`
	OrderID         int64  `json:"orderId"`
	OrderPrice      string `json:"orderPrice"`
	OrderSide       string `json:"orderSide"`
	OrderSize       string `json:"orderSize"`
	OrderStatus     string `json:"orderStatus"`
	OrderType       string `json:"orderType"`
	Source          string `json:"source"`
	Symbol          string `json:"symbol"`
}

// 3、accounts_ws_api
type accountsChangeMsg struct {
	TransferAccounts TransferAccounts `json:"transferAccounts"`
	AccountDos       []accountDo      `json:"accountDos"`
}

type accountDo struct {
	ID      int     `json:"id"`
	UID     int     `json:"uid"`
	Type    int     `json:"type"`
	Balance float64 `json:"balance"`
}

type TransferAccounts struct {
	UID          int           `json:"uid"`
	ExtParams    interface{}   `json:"extParams"`
	Transactions []transaction `json:"transactions"`
}

type transaction struct {
	AccountType   int    `json:"accountType"`
	OppositeUID   int    `json:"oppositeUid"`
	OppositeType  int    `json:"oppositeType"`
	Symbol        string `json:"symbol"`
	Amount        int    `json:"amount"`
	Direction     string `json:"direction"`
	Scene         string `json:"scene"`
	RefID         int64  `json:"refId"`
	RefType       string `json:"refType"`
	OldTransferID int64  `json:"oldTransferId"`
}

func getOrderPushMsg() {

}

// NewMsg 生成一个mq消息
func NewMsg(topic string, data []byte) *primitive.Message {
	return &primitive.Message{
		Topic: topic,
		Body:  data,
	}
}

// orders消息
func buildOrderMsg(uid int) *primitive.Message {
	msgStruct := &orderMsg{
		AccountID:       uid,
		Aggressor:       false,
		ClientOrderID:   "1572041365652844545",
		EventType:       "orders",
		OrderCreateTime: time.Now().Add(-1 * time.Second).UnixMilli(),
		OrderID:         1604106271466196992,
		OrderPrice:      "2996.54",
		OrderSide:       "sell",
		OrderSize:       "0.4118",
		OrderStatus:     "submitted",
		OrderType:       "limit",
		Source:          "3",
		Symbol:          "btcusdt",
	}
	data, _ := json.Marshal(msgStruct)
	return NewMsg("", data)
}

//
