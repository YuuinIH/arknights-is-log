package route

import (
	"github.com/YuuinIH/is-log/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	app.Get("/auth/uuid", controllers.GetNewUUID)
	app.Post("/auth/login", controllers.Login)
}
