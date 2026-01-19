package exchange

import (
	ccxtpro "github.com/ccxt/ccxt/go/v4/pro"
)

type ExchangeManager struct {
	exchanges map[ExchangeId]ccxtpro.IExchange
}

func NewExchangeManager() ExchangeManager {
	return ExchangeManager{
		exchanges: make(map[ExchangeId]ccxtpro.IExchange),
	}
}

func (self *ExchangeManager) Get(exchangeId ExchangeId) ccxtpro.IExchange {
	cachedExchange, exists := self.exchanges[exchangeId]
	if exists {
		return cachedExchange
	}

	markets, marketsExist := getCachedMarkets(exchangeId)
	if !marketsExist {
		exchange := ccxtpro.CreateExchange(exchangeId.String(), nil)
		self.exchanges[exchangeId] = exchange
		return exchange
	}

	exchange := ccxtpro.CreateExchange(exchangeId.String(), nil)

	settable, ok := exchange.(SettableMarket)
	if !ok {
		panic("exchange doesn't implement SettableMarket")
	}
	settable.SetMarkets(markets)

	self.exchanges[exchangeId] = exchange
	return exchange
}

type SettableMarket interface {
	SetMarkets(markets any, optionalArgs ...any) any
}
