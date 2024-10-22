package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tradmark/config"
	"github.com/tradmark/pkg"
)

func SetTradesroutes(app *fiber.App) {
	trads := app.Group("/trads")
	trads.Get("/", FetchTrads)
	trads.Get("/:serialnumber", FetchTradsBySerialNumber)
}

func FetchTrads(c *fiber.Ctx) error {

	caseFiles, err := pkg.TradesRepository.FetchTrads(config.GetDB())
	if err != nil {
		return err
	}
	return c.JSON(caseFiles)
}

func FetchTradsBySerialNumber(c *fiber.Ctx) error {
	serialNumber := c.Params("serialnumber")
	caseFile, err := pkg.TradesRepository.FetchTradsBySerialNumber(config.GetDB(), serialNumber)
	if err != nil {
		return err
	}
	return c.JSON(caseFile)
}
