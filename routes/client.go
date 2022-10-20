package routes

import (
	"go-echo/controllers"

	"github.com/labstack/echo"
)

func UserRoute(e *echo.Echo) {
	e.POST("/client", controllers.CreateClient)
	e.GET("/client/:id", controllers.GetClient)
}
