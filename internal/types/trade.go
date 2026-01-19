package types

import ccxtpro "github.com/ccxt/ccxt/go/v4/pro"


type TradeData struct {
	Timestamp int64
	Price     float64
	Cost      float64
}

func NewTradeData(trade ccxtpro.Trade) TradeData {
	return TradeData{
		Timestamp: *trade.Timestamp,
		Price:     *trade.Price,
		Cost:      *trade.Cost,
	}
}