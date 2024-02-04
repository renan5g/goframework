package cache

import (
	"github.com/renan5g/goframework/contracts/cache"
	"github.com/renan5g/goframework/contracts/config"
	"github.com/renan5g/goframework/contracts/log"
)

type CacheDriver string

const (
	DriverMemory CacheDriver = "memory"
	DriverRedis  CacheDriver = "redis"
	DriverCustom CacheDriver = "custom"
)

type Cacher struct {
	cache.Driver
	config config.Config
	driver Driver
	log    log.Log
	stores map[string]cache.Driver
}

func NewCacher(config config.Config, log log.Log, store string) (*Cacher, error) {
	driver := NewDriverImpl(config)
	instance, err := driver.New(store)
	if err != nil {
		return nil, err
	}

	return &Cacher{
		Driver: instance,
		config: config,
		driver: driver,
		log:    log,
		stores: map[string]cache.Driver{
			store: instance,
		},
	}, nil
}

func (app *Cacher) Store(name string) cache.Driver {
	if driver, exist := app.stores[name]; exist {
		return driver
	}

	instance, err := app.driver.New(name)
	if err != nil {
		app.log.Error(err)

		return nil
	}

	app.stores[name] = instance

	return instance
}
