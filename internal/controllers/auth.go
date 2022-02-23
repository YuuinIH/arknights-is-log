package controllers

import (
	"log"

	"github.com/YuuinIH/arknights-is-log/internal/models"
	"github.com/YuuinIH/arknights-is-log/internal/utils"
	"github.com/gofiber/fiber/v2"

	u "github.com/google/uuid"
)

// GetNewUUID godoc
// @Summary      Get a new uuid.
// @Description  Generate a new UUID.
// @Tags         auth
// @Produce      json
// @Success      200  {object}  models.JSONResult{data=models.Login}
// @Router       /auth/uuid [get]
func GetNewUUID(ctx *fiber.Ctx) error {
	uuid, _ := models.NewUUID()
	return ctx.JSON(fiber.Map{
		"code": 200,
		"msg":  models.GetMsg(200),
		"data": fiber.Map{
			"uuid": uuid.String(),
		},
	})
}

// Login godoc
// @Summary      Login with UUID.
// @Description  Login with UUID,it return a token.
// @Tags         auth
// @Accept       multipart/form-data
// @Accept       json
// @Param        uuid  body  models.Login  true  "uuid"
// @Produce      json
// @Success      200  {object}  models.JSONResult{data=models.NewToken}
// @Router       /auth/login [post]
func Login(ctx *fiber.Ctx) error {
	var a models.Login
	if err := ctx.BodyParser(&a); err != nil {
		log.Print(err)
		return ctx.Status(400).JSON(fiber.Map{
			"code": 400,
			"msg":  models.GetMsg(400),
			"data": nil,
		})
	}
	uuid, err := u.Parse(a.UUID)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"code": 400,
			"msg":  models.GetMsg(400),
			"data": nil,
		})
	}
	isuuidexist, err := models.CheckUUID(uuid)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"code": 500,
			"msg":  models.GetMsg(500),
			"data": nil,
		})
	}
	if !isuuidexist {
		return ctx.Status(404).JSON(fiber.Map{
			"code": 20001,
			"msg":  models.GetMsg(20001),
			"data": nil,
		})
	}
	ss, err := utils.GenerateToken(uuid)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"code": 500,
			"msg":  models.GetMsg(500) + err.Error(),
			"data": nil,
		})
	}
	return ctx.JSON(fiber.Map{
		"code": 200,
		"msg":  models.GetMsg(200),
		"data": fiber.Map{
			"token": ss,
		},
	})
}
