package facade

import (
	corecontract "github.com/renan5g/goframework/contracts/core"
	"github.com/renan5g/goframework/core"
)

func App() corecontract.Application {
	return core.App
}
