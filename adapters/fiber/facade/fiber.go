package facade

import (
	"log"

	"github.com/renan5g/goframework/contracts/route"

	"github.com/renan5g/goframework/adapters/fiber"
)

func Route(driver string) route.Route {
	instance, err := fiber.App.MakeWith(fiber.RouteBinding, map[string]any{
		"driver": driver,
	})
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(*fiber.Route)
}
