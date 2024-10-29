package public

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	SetTradesroutes(app)
}
