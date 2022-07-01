package size

import (
	"errors"

	"github.com/DollarCoinman/gocryptotrader/backtester/eventhandlers/exchange"
)

var (
	errNoFunds         = errors.New("no funds available")
	errLessThanMinimum = errors.New("sized amount less than minimum")
	errCannotAllocate  = errors.New("portfolio manager cannot allocate funds for an order")
)

// Size contains buy and sell side rules
type Size struct {
	BuySide  exchange.MinMax
	SellSide exchange.MinMax
}
