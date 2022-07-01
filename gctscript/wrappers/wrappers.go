package wrappers

import (
	"github.com/DollarCoinman/gocryptotrader/gctscript/modules"
	"github.com/DollarCoinman/gocryptotrader/gctscript/wrappers/validator"
)

// GetWrapper returns the instance of each wrapper to use
func GetWrapper() modules.GCTExchange {
	if validator.IsTestExecution.Load() == true {
		return validator.Wrapper{}
	}
	return modules.Wrapper
}
