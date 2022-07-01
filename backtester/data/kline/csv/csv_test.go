package csv

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/DollarCoinman/gocryptotrader/backtester/common"
	"github.com/DollarCoinman/gocryptotrader/currency"
	"github.com/DollarCoinman/gocryptotrader/exchanges/asset"
	gctkline "github.com/DollarCoinman/gocryptotrader/exchanges/kline"
)

const testExchange = "binance"

func TestLoadDataCandles(t *testing.T) {
	exch := testExchange
	a := asset.Spot
	p := currency.NewPair(currency.BTC, currency.USDT)
	_, err := LoadData(
		common.DataCandle,
		filepath.Join("..", "..", "..", "..", "testdata", "binance_BTCUSDT_24h_2019_01_01_2020_01_01.csv"),
		exch,
		gctkline.FifteenMin.Duration(),
		p,
		a,
		false)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadDataTrades(t *testing.T) {
	exch := testExchange
	a := asset.Spot
	p := currency.NewPair(currency.BTC, currency.USDT)
	_, err := LoadData(
		common.DataTrade,
		filepath.Join("..", "..", "..", "..", "testdata", "binance_BTCUSDT_24h-trades_2020_11_16.csv"),
		exch,
		gctkline.FifteenMin.Duration(),
		p,
		a,
		false)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadDataInvalid(t *testing.T) {
	exch := testExchange
	a := asset.Spot
	p := currency.NewPair(currency.BTC, currency.USDT)
	_, err := LoadData(
		-1,
		filepath.Join("..", "..", "..", "..", "testdata", "binance_BTCUSDT_24h-trades_2020_11_16.csv"),
		exch,
		gctkline.FifteenMin.Duration(),
		p,
		a,
		false)
	if !errors.Is(err, common.ErrInvalidDataType) {
		t.Errorf("received: %v, expected: %v", err, common.ErrInvalidDataType)
	}

	_, err = LoadData(
		-1,
		filepath.Join("..", "..", "..", "..", "testdata", "binance_BTCUSDT_24h-trades_2020_11_16.csv"),
		exch,
		gctkline.FifteenMin.Duration(),
		p,
		a,
		true)
	if !errors.Is(err, errNoUSDData) {
		t.Errorf("received: %v, expected: %v", err, errNoUSDData)
	}
}
