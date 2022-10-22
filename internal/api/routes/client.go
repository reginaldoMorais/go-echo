package routes

import (
	"go-echo/internal/api/controllers"

	"github.com/labstack/echo"
)

func UserRoute(e *echo.Echo) {
	controllers := controllers.NewClientController()

	e.GET("/client", controllers.GetClients)
	e.GET("/client/:id", controllers.GetClient)
	e.POST("/client", controllers.CreateClient)
	e.DELETE("/client/:id", controllers.DeleteClient)
}
