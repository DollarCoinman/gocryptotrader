package kline

import (
	"github.com/DollarCoinman/gocryptotrader/backtester/eventtypes/event"
	"github.com/shopspring/decimal"
)

// Kline holds kline data and an event to be processed as
// a common.DataEventHandler type
type Kline struct {
	*event.Base
	Open             decimal.Decimal
	Close            decimal.Decimal
	Low              decimal.Decimal
	High             decimal.Decimal
	Volume           decimal.Decimal
	ValidationIssues string
}
