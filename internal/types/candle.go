package types

import (

	config "github.com/yourusername/Market-Data/internal"
)

// Candle represents an aggregated candle with OHLCV data
type Candle struct {
	Timestamp  int64
	Open       float32
	High       float32
	Low        float32
	Close      float32
	Volume     float32
	Vwap       float32
	TradeCount uint32
}

// CandleBuilder aggregates trades into candles
type CandleBuilder struct {
	timestamp  int64
	open       float32
	high       float32
	low        float32
	close      float32
	volume     float32
	vw         float64
	tradeCount uint32
}

// NewCandleBuilder creates a new CandleBuilder from the first trade
func NewCandleBuilder(trade TradeData) CandleBuilder {
	price := float32(trade.Price)
	return CandleBuilder{
		timestamp:  getTimeKey(trade),
		open:       price,
		high:       price,
		low:        price,
		close:      price,
		volume:     float32(trade.Cost),
		vw:         trade.Price * trade.Cost,
		tradeCount: 1,
	}
}

// AddTrade adds a trade to the candle builder
// Returns a completed Candle if the trade belongs to a new time window, nil otherwise
func (self *CandleBuilder) AddTrade(trade TradeData) (*Candle,bool) {
	
	
	tradeTimeKey := getTimeKey(trade)

	// If trade is in a new time window, return the completed candle
	if tradeTimeKey != self.timestamp {
		if tradeTimeKey < self.timestamp {
			return nil,true
		}

		candle := self.ToCandle()
		*self = NewCandleBuilder(trade)
		return &candle,false
	}
	
	// Update candle state with the new trade
	price := float32(trade.Price)
	self.close = price
	if price > self.high {
		self.high = price
	}
	if price < self.low {
		self.low = price
	}
	self.volume += float32(trade.Cost)
	self.vw += trade.Price * trade.Cost
	self.tradeCount++

	return nil,false
}

// ToCandle converts the builder state to a Candle
func (self *CandleBuilder) ToCandle() Candle {
	var vwap float32
	if self.volume > 0 {
		vwap = float32(self.vw / float64(self.volume))
	}

	return Candle{
		Timestamp:  self.timestamp,
		Open:       self.open,
		High:       self.high,
		Low:        self.low,
		Close:      self.close,
		Volume:     self.volume,
		Vwap:       vwap,
		TradeCount: self.tradeCount,
	}
}

func getTimeKey(trade TradeData) int64 {
	return (trade.Timestamp / config.CandleDurationMs) * config.CandleDurationMs
}
