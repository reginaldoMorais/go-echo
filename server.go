package main

import (
	"go-echo/configs"
	"go-echo/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	configs.ConnectDB()

	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":1323"))
}
