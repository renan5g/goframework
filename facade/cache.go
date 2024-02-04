package facade

import (
	"github.com/renan5g/goframework/contracts/cache"
)

func Cache() cache.Cache {
	return App().MakeCache()
}
