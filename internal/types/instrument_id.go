package types

import (
	"fmt"
	"strings"
	"unique"

	"github.com/yourusername/Market-Data/internal/types/exchange"
)

// InstrumentId combines an ExchangeId and Symbol
type InstrumentId struct {
	Exchange exchange.ExchangeId
	Symbol   Symbol
}


func NewInstrumentId(exchangeName string, base, quote string) (InstrumentId, error) {
	exc := exchange.NewExchangeId(exchangeName)

	if !exc.Validate() {
		return InstrumentId{}, fmt.Errorf("invalid exchange: %s", exchangeName)
	}

	return InstrumentId{
		Exchange: exc,
		Symbol:   NewSymbol(base, quote),
	}, nil
}

func (id InstrumentId) String() string {
	return fmt.Sprintf("%s@%s", id.Symbol.String(), id.Exchange.String())
}




type Symbol struct {
	base  unique.Handle[string]
	quote unique.Handle[string]
}

// NewSymbol creates a new Symbol with lowercase currencies
func NewSymbol(base, quote string) Symbol {
	return Symbol{
		base:  unique.Make(strings.ToLower(base)),
		quote: unique.Make(strings.ToLower(quote)),
	}
}

// String returns the symbol in "base/quote" format
func (self Symbol) String() string {
	return strings.ToUpper(self.Base() + "/" + self.Quote())
}

func (self Symbol) Base() string {
	return self.base.Value()
}

func (self Symbol) Quote() string {
	return self.quote.Value()
}
