package route

import (
	"github.com/YuuinIH/arknights-is-log/internal/config"
	"github.com/YuuinIH/arknights-is-log/internal/controllers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func ApiRoute(app *fiber.App) {
	v1 := app.Group("/api/v1").Use(jwtware.New(jwtware.Config{
		SigningKey:    config.PrivateKey.Public(),
		SigningMethod: "RS256",
	}))
	{
		{
			v1.Post("/report", controllers.AddReport)
			v1.Post("/reports", controllers.AddReports)
			v1.Get("/report", controllers.GetReportListByAccount)
			v1.Get("/report/:id", controllers.GetReportByID)
			v1.Delete("/report/:id", controllers.DeleteReportByID)
		}
	}
}
