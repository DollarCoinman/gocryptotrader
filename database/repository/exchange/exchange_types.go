package exchange

import (
	"errors"

	"github.com/DollarCoinman/gocryptotrader/common/cache"
	"github.com/gofrs/uuid"
)

var (
	exchangeCache = cache.New(30)
	// ErrNoExchangeFound is a basic predefined error
	ErrNoExchangeFound = errors.New("exchange not found")
)

// Details holds exchange information such as Name
type Details struct {
	UUID uuid.UUID
	Name string
}
