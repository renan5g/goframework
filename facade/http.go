package facade

import (
	"github.com/renan5g/goframework/contracts/http"
)

func RateLimiter() http.RateLimiter {
	return App().MakeRateLimiter()
}

func View() http.View {
	return App().MakeView()
}
