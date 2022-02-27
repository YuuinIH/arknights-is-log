package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(
		cors.New(),
		compress.New(),
		recover.New(),
	)
}
