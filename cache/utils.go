package cache

import (
	"github.com/renan5g/goframework/contracts/config"
)

func prefix(config config.Config) string {
	return config.GetString("cache.prefix") + ":"
}
