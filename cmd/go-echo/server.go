package main

import (
	"go-echo/internal/api/routes"
	"go-echo/internal/core"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	core.ConnectDB()

	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":1323"))
}
