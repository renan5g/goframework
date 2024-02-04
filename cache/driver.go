package cache

import (
	"context"
	"fmt"

	"github.com/renan5g/goframework/contracts/cache"
	"github.com/renan5g/goframework/contracts/config"
)

//go:generate mockery --name=Driver
type Driver interface {
	New(store string) (cache.Driver, error)
}

type DriverImpl struct {
	config config.Config
}

func NewDriverImpl(config config.Config) *DriverImpl {
	return &DriverImpl{
		config: config,
	}
}

func (d *DriverImpl) New(store string) (cache.Driver, error) {
	driver := CacheDriver(d.config.GetString(fmt.Sprintf("cache.stores.%s.driver", store)))
	switch driver {
	case DriverMemory:
		return d.memory()
	case DriverRedis:
		return d.custom(store)
	case DriverCustom:
		return d.custom(store)
	default:
		return nil, fmt.Errorf("invalid driver: %s, only support memory, custom\n", driver)
	}
}

func (d *DriverImpl) memory() (cache.Driver, error) {
	memory, err := NewMemory(d.config)
	if err != nil {
		return nil, fmt.Errorf("init memory driver error: %v", err)
	}

	return memory, nil
}

func (d *DriverImpl) redis(store string) (cache.Driver, error) {
	memory, err := NewRedis(context.Background(), d.config, store)
	if err != nil {
		return nil, fmt.Errorf("init redis driver error: %v", err)
	}

	return memory, nil
}

func (d *DriverImpl) custom(store string) (cache.Driver, error) {
	if custom, ok := d.config.Get(fmt.Sprintf("cache.stores.%s.via", store)).(cache.Driver); ok {
		return custom, nil
	}
	if custom, ok := d.config.Get(fmt.Sprintf("cache.stores.%s.via", store)).(func() (cache.Driver, error)); ok {
		return custom()
	}

	return nil, fmt.Errorf("%s doesn't implement contracts/cache/store\n", store)
}
