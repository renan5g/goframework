package facade

import (
	"github.com/renan5g/goframework/contracts/route"
)

func Route() route.Route {
	return App().MakeRoute()
}
