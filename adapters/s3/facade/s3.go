package facades

import (
	"log"

	"github.com/renan5g/goframework/contracts/filesystem"

	"github.com/renan5g/goframework/adapters/s3"
)

func S3(disk string) filesystem.Driver {
	instance, err := s3.App.MakeWith(s3.Binding, map[string]any{"disk": disk})
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(*s3.S3)
}
