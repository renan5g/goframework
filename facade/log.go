package facade

import (
	"github.com/renan5g/goframework/contracts/log"
)

func Log() log.Log {
	return App().MakeLog()
}
