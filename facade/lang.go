package facade

import (
	"context"

	"github.com/renan5g/goframework/contracts/translation"
)

func Lang(ctx context.Context) translation.Translator {
	return App().MakeLang(ctx)
}
