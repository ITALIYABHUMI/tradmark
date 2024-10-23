package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tradmark/config"
	"github.com/tradmark/pkg"
)

func SetTradesroutes(app *fiber.App) {
	trads := app.Group("/trads")
	trads.Get("/indics", CreateIndics)
	trads.Get("/serialnumber/:serialnumber", FetchTradsBySerialNumber)
	trads.Get("/:data", SearchTradMarks)

}

func FetchTradsBySerialNumber(c *fiber.Ctx) error {
	serialNumber := c.Params("serialnumber")
	caseFile, err := pkg.SearchRepository.FetchTradsBySerialNumber(config.GetDB(), serialNumber)
	if err != nil {
		return err
	}
	return c.JSON(caseFile)
}

func CreateIndics(c *fiber.Ctx) error {
	resp, err := pkg.SearchRepository.Create(config.GetDB())
	if err != nil || resp != nil {
		return fmt.Errorf("index not created")
	} else {
		return fmt.Errorf("index created")
	}
}

func SearchTradMarks(c *fiber.Ctx) error {
	data := c.Params("data")
	resp, err := pkg.SearchRepository.Search(config.GetDB(), data)
	if err != nil {
		return err
	}
	return c.JSON(resp)
}
