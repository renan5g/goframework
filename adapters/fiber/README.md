# Fiber

Fiber http driver for Goframework.

## Install

1. Register service provider

```
// config/app.go
import "github.com/goravel/fiber"

"providers": []foundation.ServiceProvider{
    ...
    &fiber.ServiceProvider{},
}
```

3. Add fiber config to `config/http.go` file

```
// config/http.go
import (
    fiberfacades "github.com/renan5g/goframework/adapters/fiber/facade"
    "github.com/gofiber/template/html/v2"
    "github.com/gofiber/fiber/v2"
)

"default": "fiber",

"drivers": map[string]any{
    "fiber": map[string]any{
        // prefork mode, see https://docs.gofiber.io/api/fiber/#config
        "prefork": false,
        // Optional, default is 4096 KB
        "body_limit": 4096,
        "route": func() (route.Route, error) {
            return fiberfacades.Route("fiber"), nil
        },
        // Optional, default is "html/template"
        "template": func() (fiber.Views, error) {
            return html.New("./resources/views", ".tmpl"), nil
        },
    },
},
```

## Testing

Run command below to run test:

```
go test ./...
```
