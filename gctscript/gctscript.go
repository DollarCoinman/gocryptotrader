package gctscript

import (
	"github.com/DollarCoinman/gocryptotrader/gctscript/modules"
	"github.com/DollarCoinman/gocryptotrader/gctscript/wrappers/gct"
)

// Setup configures the wrapper interface to use
func Setup() {
	modules.SetModuleWrapper(gct.Setup())
}
