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
func NewMsg(topic string, msgStruct interface{}) *primitive.Message {
	data, _ := json.Marshal(msgStruct)
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
	return NewMsg("", msgStruct)
}

// trade消息
func buildTradeMsg(uid int) *primitive.Message {
	msgStruct := &tradeMsg{
		AccountID:       uid,
		Aggressor:       false,
		ClientOrderID:   "1572041365652844545",
		EventType:       "orders",
		OrderCreateTime: time.Now().Add(-1 * time.Second).UnixMilli(),
		OrderID:         1604106271466196992,
		OrderPrice:      "2996.54",
		OrderSide:       "sell",
		OrderSize:       "0.4118",
		OrderStatus:     "filled",
		OrderType:       "limit",
		Source:          "3",
		Symbol:          "btcusdt",
	}
	return NewMsg("", msgStruct)
}

// accounts
func buildAccountsMsg(uid int) *primitive.Message {
	//1、accountDos
	ad1 := accountDo{
		ID:      1266378177,
		UID:     uid,
		Type:    2011589,
		Balance: 9900.0000000000000000,
	}

	ad2 := accountDo{
		ID:      1266378178,
		UID:     uid,
		Type:    2021589,
		Balance: 100.0000000000000000,
	}
	dos := []accountDo{ad1, ad2}

	//2、transferAccounts
	t1 := transaction{
		AccountType:   2011589,
		OppositeUID:   1008424786,
		OppositeType:  2021589,
		Symbol:        "BIB",
		Amount:        100,
		Direction:     "OUT",
		Scene:         "create_order",
		RefID:         1607295604050137088,
		RefType:       "ex_order_bibusdt",
		OldTransferID: 1607295604050137089,
	}

	t2 := transaction{
		AccountType:   2021589,
		OppositeUID:   1008424786,
		OppositeType:  2011589,
		Symbol:        "BIB",
		Amount:        100,
		Direction:     "IN",
		Scene:         "create_order",
		RefID:         1607295604050137088,
		RefType:       "ex_order_bibusdt",
		OldTransferID: 1607295604050137089,
	}
	ts := []transaction{t1, t2}
	ta := TransferAccounts{
		UID:          uid,
		ExtParams:    nil,
		Transactions: ts,
	}

	//3、all
	msgStruct := accountsChangeMsg{
		TransferAccounts: ta,
		AccountDos:       dos,
	}

	return NewMsg("", msgStruct)
}
