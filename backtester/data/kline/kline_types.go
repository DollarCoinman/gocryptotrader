package kline

import (
	"errors"

	"github.com/DollarCoinman/gocryptotrader/backtester/data"
	gctkline "github.com/DollarCoinman/gocryptotrader/exchanges/kline"
)

var errNoCandleData = errors.New("no candle data provided")

// DataFromKline is a struct which implements the data.Streamer interface
// It holds candle data for a specified range with helper functions
type DataFromKline struct {
	data.Base
	addedTimes  map[int64]bool
	Item        gctkline.Item
	RangeHolder *gctkline.IntervalRangeHolder
}
