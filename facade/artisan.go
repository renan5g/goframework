package facade

import (
	"github.com/renan5g/goframework/contracts/console"
)

func Artisan() console.Artisan {
	return App().MakeArtisan()
}
