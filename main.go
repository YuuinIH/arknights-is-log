package main

import (
	"crypto/rsa"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/YuuinIH/arknights-is-log/internal/config"
	"github.com/YuuinIH/arknights-is-log/internal/middleware"
	"github.com/YuuinIH/arknights-is-log/internal/route"
)

var (
	privateKey *rsa.PrivateKey
)

// @title        IS-LOG Api
// @version      1.0
// @description  The api of the IS-LOG.

// @contact.name   YuuinIH
// @contact.email  zsf821797423@gmail.com

// @license.name  MIT License
// @license.url   https://opensource.org/licenses/MIT

// @securityDefinitions.apiKey  JWT
// @in                          header
// @name                        Authorization

// @host  localhost:8080
func main() {
	app := fiber.New()

	middleware.FiberMiddleware(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	route.SwaggerRoute(app)
	route.AuthRoute(app)
	route.ApiRoute(app)

	log.Fatal(app.Listen(":" + fmt.Sprintf("%d", config.SERVER.HTTP_PORT)))
}
