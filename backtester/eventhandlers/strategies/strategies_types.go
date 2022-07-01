package strategies

import (
	"github.com/DollarCoinman/gocryptotrader/backtester/data"
	"github.com/DollarCoinman/gocryptotrader/backtester/eventhandlers/portfolio"
	"github.com/DollarCoinman/gocryptotrader/backtester/eventtypes/signal"
	"github.com/DollarCoinman/gocryptotrader/backtester/funding"
)

// Handler defines all functions required to run strategies against data events
type Handler interface {
	Name() string
	Description() string
	OnSignal(data.Handler, funding.IFundingTransferer, portfolio.Handler) (signal.Event, error)
	OnSimultaneousSignals([]data.Handler, funding.IFundingTransferer, portfolio.Handler) ([]signal.Event, error)
	UsingSimultaneousProcessing() bool
	SupportsSimultaneousProcessing() bool
	SetSimultaneousProcessing(bool)
	SetCustomSettings(map[string]interface{}) error
	SetDefaults()
}
