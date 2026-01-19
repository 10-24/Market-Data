package exchange

import (
	"slices"
	"strings"
	"unique"
	ccxtpro "github.com/ccxt/ccxt/go/v4/pro"
)

// ExchangeId represents a unique identifier for an exchange
type ExchangeId struct {
	name unique.Handle[string]
}

func NewExchangeId(name string) ExchangeId {
	return ExchangeId{
		name: unique.Make(strings.ToLower(name)),
	}
}

// String returns the lowercase string representation
func (self ExchangeId) String() string {
	return self.name.Value()
}

func (self ExchangeId) Validate() bool {
	_, found := slices.BinarySearch(ccxtpro.Exchanges, self.String())
	return found
}
