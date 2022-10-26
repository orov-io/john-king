package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(setContextLogger(e.Logger))
	e.Use(BodyDumpOnHeader())
	e.GET("/ping", func(c echo.Context) error {
		/* 		c.Logger().Debug("A debug")
		   		c.Logger().Info("A info")
		   		c.Logger().Warn("A warn")
		   		c.Logger().Print("A print")
		   		c.Logger().Error("A error") */
		return c.JSON(http.StatusOK, echo.Map{"message": "pong"})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
