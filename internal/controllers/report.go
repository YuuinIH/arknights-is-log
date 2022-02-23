package controllers

import (
	"log"
	"strconv"

	"github.com/YuuinIH/arknights-is-log/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// AddReport godoc
// @Summary      Upload a single report.
// @Description  Upload a single report.
// @Tags         report
// @Accept       json
// @Param        report  body  models.Roguelike_Report  true  "Report"
// @Produce      json
// @Success      200  {object}  models.JSONResult{data=models.ReportID}
// @Router       /api/v1/report [post]
//TODO:Added verification of upload report。
func AddReport(ctx *fiber.Ctx) error {
	p := new(models.Roguelike_Report)
	code := 500
	uuid, err := uuid.Parse(ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["uuid"].(string))
	if err != nil {
		code = models.ERROR_AUTH_CHECK_TOKEN_FAIL
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": nil,
		})
	}
	if err := ctx.BodyParser(p); err != nil {
		log.Println(err)
		code = models.INVALID_PARAMS
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": nil,
		})
	}
	if result, err := models.AddReport(*p, uuid); err != nil && err.Error() == "report already exists" {
		code = models.ERROR_EXIST_REPORT
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": nil,
		})
	} else if err != nil {
		code = models.ERROR
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": nil,
		})
	} else {
		log.Println(result)
		code = models.CREATED
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": fiber.Map{
				"id": result},
		})
	}
}

// AddReports godoc
// @Summary      Upload multiple reports.
// @Description  Upload multiple reports.
// @Tags         report
// @Param        report  body  []models.Roguelike_Report  true  "Reports"
// @Produce      json
// @Success      200  {object}  models.JSONResult{data=models.ReportIDs}
// @Router       /api/v1/reports [post]
func AddReports(ctx *fiber.Ctx) error {
	p := new([]models.Roguelike_Report)
	code := 500
	var err error
	uuid, err := uuid.Parse(ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["uuid"].(string))
	if err != nil {
		code = models.ERROR_AUTH_CHECK_TOKEN_FAIL
	}
	if err = ctx.BodyParser(p); err != nil {
		log.Println(err)
		code = models.INVALID_PARAMS
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": nil,
		})
	}
	if result, err := models.AddReports(*p, uuid); err == nil {
		code = models.CREATED
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": nil,
		})
	} else if err.Error() == "all reports already exists" {
		code = models.ERROR_EXIST_REPORT
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": nil,
		})
	} else {
		code = models.ERROR
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": fiber.Map{
				"id": result,
			},
		})
	}
}

// GetReportListByAccount godoc
// @Summary      Get reports under this account.
// @Description  Get reports under this account.
// @Tags         report
// @Param        order     query  string  false  "order"
// @Param        page      query  int     true   "page"
// @Param        pagesize  query  int     true   "pagesize"
// @Produce      json
// @Success      200  {object}  models.JSONResult{data=[]models.Roguelike_Report_With_ID}
// @Router       /api/v1/report [get]
func GetReportListByAccount(ctx *fiber.Ctx) error {
	type input struct {
		order    string
		page     int
		pagesize int
	}
	p := new(input)
	code := 500
	var err error
	uuid, err := uuid.Parse(ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["uuid"].(string))
	if err != nil {
		code = models.ERROR_AUTH_CHECK_TOKEN_FAIL
	}
	var result []models.Roguelike_Report_With_ID
	{
		p.order = ctx.Query("order", "")
		p.page, _ = strconv.Atoi(ctx.Query("page"))
		p.pagesize, _ = strconv.Atoi(ctx.Query("pagesize"))
	}
	if err = models.GetReportListByAccount(uuid.String(), p.page, p.pagesize, &result); err != nil {
		log.Println(err)
		if err.Error() == "No reports was found." {
			return ctx.JSON(fiber.Map{
				"code": models.ERROR_NOT_EXIST_REPORT,
				"msg":  models.GetMsg(models.ERROR_NOT_EXIST_REPORT),
				"data": nil,
			})
		}
		code = models.ERROR
	} else {
		code = models.SUCCESS
	}
	return ctx.JSON(fiber.Map{
		"code": code,
		"msg":  models.GetMsg(code),
		"data": result,
	})
}

// GetReportByID godoc
// @Summary      Get report corresponding to ID.
// @Description  Get report corresponding to ID.
// @Tags         report
// @Accept       json
// @Param        id  path  string  true  "id"
// @Produce      json
// @Success      200  {object}  models.JSONResult{data=[]models.Roguelike_Report_With_ID}
// @Router       /api/v1/report/{id} [get]
func GetReportByID(ctx *fiber.Ctx) error {
	p := ctx.Params("id")
	var code int
	var result models.Roguelike_Report_With_ID
	if err := models.GetReportByID(p, &result); err != nil {
		log.Println(err)
		if err.Error() == "mongo: no documents in result" {
			code = models.ERROR_NOT_EXIST_REPORT
			return ctx.JSON(fiber.Map{
				"code": code,
				"msg":  models.GetMsg(code),
				"data": nil,
			})
		}
		code = models.ERROR
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": nil,
		})
	} else {
		code = models.SUCCESS
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": result,
		})
	}
}

// DeleteReportByID godoc
// @Summary      Delete report corresponding to ID.
// @Description  Delete report corresponding to ID.
// @Tags         report
// @Accept       json
// @Param        id  path  string  true  "id"
// @Produce      json
// @Success      200  {object}  models.JSONResult{data=models.ReportID}
// @Router       /api/v1/report/{id} [delete]
func DeleteReportByID(ctx *fiber.Ctx) error {
	p := ctx.Params("id")
	var code int
	if err := models.DeleteReport(p); err != nil {
		log.Println(err)
		code = models.ERROR
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": "",
		})
	} else {
		code = models.SUCCESS
		return ctx.JSON(fiber.Map{
			"code": code,
			"msg":  models.GetMsg(code),
			"data": fiber.Map{
				"id": p,
			},
		})
	}
}
