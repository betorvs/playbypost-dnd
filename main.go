package main

import (
	"github.com/betorvs/playbypost-dnd-api/config"
	"github.com/betorvs/playbypost-dnd-api/controller"
	_ "github.com/betorvs/playbypost-dnd-api/gateway/customlog"
	_ "github.com/betorvs/playbypost-dnd-api/gateway/mongodb"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.MapRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Values.Port))
}
