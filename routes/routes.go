package routes

import (
	"checkout-services/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/checkout", controllers.PostCheckoutController)

	return e
}
