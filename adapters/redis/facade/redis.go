package facades

import (
	"log"

	"github.com/renan5g/goframework/contracts/cache"

	"github.com/renan5g/goframework/adapters/redis"
)

func Redis(store string) cache.Driver {
	if store == "" {
		log.Fatalln("store is required")
		return nil
	}

	instance, err := redis.App.MakeWith(redis.Binding, map[string]any{"store": store})
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(*redis.Redis)
}
