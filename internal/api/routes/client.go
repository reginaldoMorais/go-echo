package routes

import (
	"go-echo/internal/api/controllers"

	"github.com/labstack/echo"
)

func UserRoute(e *echo.Echo) {
	controllers := controllers.NewClientController()

	e.POST("/client", controllers.CreateClient)
	e.GET("/client/:id", controllers.GetClient)
}
