package exchange

import (
	"errors"

	"github.com/DollarCoinman/gocryptotrader/backtester/data"
	"github.com/DollarCoinman/gocryptotrader/backtester/eventtypes/fill"
	"github.com/DollarCoinman/gocryptotrader/backtester/eventtypes/order"
	"github.com/DollarCoinman/gocryptotrader/backtester/funding"
	"github.com/DollarCoinman/gocryptotrader/currency"
	"github.com/DollarCoinman/gocryptotrader/engine"
	exchange "github.com/DollarCoinman/gocryptotrader/exchanges"
	"github.com/DollarCoinman/gocryptotrader/exchanges/asset"
	gctorder "github.com/DollarCoinman/gocryptotrader/exchanges/order"
	"github.com/shopspring/decimal"
)

var (
	errDataMayBeIncorrect      = errors.New("data may be incorrect")
	errExceededPortfolioLimit  = errors.New("exceeded portfolio limit")
	errNilCurrencySettings     = errors.New("received nil currency settings")
	errInvalidDirection        = errors.New("received invalid order direction")
	errNoCurrencySettingsFound = errors.New("no currency settings found")
)

// ExecutionHandler interface dictates what functions are required to submit an order
type ExecutionHandler interface {
	SetExchangeAssetCurrencySettings(asset.Item, currency.Pair, *Settings)
	GetCurrencySettings(string, asset.Item, currency.Pair) (Settings, error)
	ExecuteOrder(order.Event, data.Handler, *engine.OrderManager, funding.IFundReleaser) (fill.Event, error)
	Reset()
}

// Exchange contains all the currency settings
type Exchange struct {
	CurrencySettings []Settings
}

// Settings allow the eventhandler to size an order within the limitations set by the config file
type Settings struct {
	Exchange      exchange.IBotExchange
	UseRealOrders bool

	Pair  currency.Pair
	Asset asset.Item

	MakerFee decimal.Decimal
	TakerFee decimal.Decimal

	BuySide  MinMax
	SellSide MinMax

	Leverage Leverage

	MinimumSlippageRate decimal.Decimal
	MaximumSlippageRate decimal.Decimal

	Limits                  gctorder.MinMaxLevel
	CanUseExchangeLimits    bool
	SkipCandleVolumeFitting bool

	UseExchangePNLCalculation bool
}

// MinMax are the rules which limit the placement of orders.
type MinMax struct {
	MinimumSize  decimal.Decimal
	MaximumSize  decimal.Decimal
	MaximumTotal decimal.Decimal
}

// Leverage rules are used to allow or limit the use of leverage in orders
// when supported
type Leverage struct {
	CanUseLeverage                 bool
	MaximumOrdersWithLeverageRatio decimal.Decimal
	MaximumLeverageRate            decimal.Decimal
}
