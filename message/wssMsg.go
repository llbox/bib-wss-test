package message

import (
	"bib-wss-test/client"
	"bib-wss-test/common"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

type order struct {
	//主键
	ID int64 `json:"id"`
	//客户端订单标识
	ClientID string `json:"clientId"`
	//用户id
	UID int64 `json:"uid"`
	//持仓类型(1 全仓，2 仓逐)
	PositionType int `json:"positionType"`
	//开平仓方向(open 开仓，close 平仓)
	Open string `json:"open"`
	//买卖方向（buy 买入，sell 卖出）
	Side string `json:"side"`
	//订单类型【LIMIT(1, "限价委托"), MARKET(2, "市价委托"), IOC(3, "无 法立即成交的部分就撤销"), FOK(4, "无法全部立即成交就撤销"), POST_ONLY(5, "无法成为被动单就撤销");】
	Type int `json:"type"`
	//杠杆倍数
	LeverageLevel int `json:"leverageLevel"`
	//下单价格
	Price float64 `json:"price"`
	//下单数量
	Volume float64 `json:"volume"`
	//剩余数量(开仓市价单：剩余金额)
	RemainingVolume float64 `json:"remainingVolume"`
	//开仓taker手续费
	OpenTakerFeeRate float64 `json:"openTakerFeeRate"`
	//开仓maker手续费
	OpenMakerFeeRate float64 `json:"openMakerFeeRate"`
	//平仓taker手续费
	CloseTakerFeeRate float64 `json:"closeTakerFeeRate"`
	//平仓maker手续费
	CloseMakerFeeRate float64 `json:"closeMakerFeeRate"`
	//订单累计盈亏
	RealizedAmount float64 `json:"realizedAmount"`
	//已成交的数量（张）
	DealVolume int `json:"dealVolume"`
	//已成交的金额
	DealMoney float64 `json:"dealMoney"`
	//成交均价
	AvgPrice float64 `json:"avgPrice"`
	//交易手续费
	TradeFee float64 `json:"tradeFee"`
	//订单状态（订单状态：0 init 初始订单，1 new 新订单，2 filled 完全成交，3 part_filled 部分成交，4 canceled 已撤单，5 pending_cancel 待撤单，6 expired 异常订单）
	Status int `json:"status"`
	//订单来源（订单来源：1web，2app，3api，4其它）
	Source int `json:"source"`
	//商户id
	BrokerID int `json:"brokerId"`
	//平仓仓位id
	PositionID int `json:"positionId"`
	Offset     int `json:"offset"`
}

type tradeOrder struct {
	BuyOrder  interface{} `json:"buyOrder"`
	SellOrder interface{} `json:"sellOrder"`
	//take方标识(1 buyer为taker、 2 seller为taker)
	ActiveSide int     `json:"activeSide"`
	TradeSize  int     `json:"tradeSize"`
	TradePrice float64 `json:"tradePrice"`
	TradeTime  int64   `json:"tradeTime"`
	//消息类型
	Type string `json:"type"`
	//开平仓标记 （1 双方都是开仓、 2 买单为开仓单，卖单为平仓单、3 卖单为开仓单，买单为平仓单、4 双方都是平仓单
	OpenType int `json:"openType"`
	Offset   int `json:"offset"`
}

// 获取毫秒时间戳
func getTime() int64 {
	return time.Now().UnixMilli()
}

// GetSelfTradeMsg 自成交消息
func GetSelfTradeMsg(symbol string) *primitive.Message {
	timeStamp := getTime()
	msgBody := &tradeOrder{
		BuyOrder:   nil,
		SellOrder:  nil,
		ActiveSide: 2,
		TradeSize:  42,
		TradePrice: 111.24,
		TradeTime:  timeStamp,
		Type:       "trade",
		OpenType:   4,
		Offset:     0,
	}
	data, _ := json.Marshal(msgBody)
	return getMsg(data, symbol)
}

func getMsg(data []byte, symbol string) *primitive.Message {
	return &primitive.Message{
		Topic: common.Topic + symbol,
		Body:  data,
	}
}

func WssSendPump(h *client.Hub) {
	symbol1 := "ETHUSDT"
	msg := GetSelfTradeMsg(symbol1)
	h.SendChan <- msg
}
