package main

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const logLevelHeader = "X-Loglevel"

var defaultLevel = log.DEBUG

func setContextLogger(logger echo.Logger) echo.MiddlewareFunc {
	logger.SetLevel(log.DEBUG)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.SetLogger(log.New("context"))
			c.Logger().SetOutput(logger.Output())

			// Force the auth middleware to set a header in the request (X-LoggedUserID) and use it
			// to provide a Format in the middleware.Logger to have the default format with previous header.

			c.Logger().SetHeader(`{"id":"${header:X-Request-ID}",,"time":"${time_rfc3339_nano},"level":"${level}","id":"${id}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}"`)
			c.Logger().SetLevel(getLogLevel(c))

			// We wan't know if the logger is different in each request:
			echoLogger := c.Echo().Logger
			contextLogger := c.Logger()

			echoLogger.Print("echo")
			contextLogger.Print("context")

			return next(c)
		}
	}
}

func getLogLevel(c echo.Context) log.Lvl {
	logLevelHeader := c.Request().Header.Get(logLevelHeader)
	if logLevelHeader == "" {
		return c.Logger().Level()
	}

	u64, err := strconv.ParseUint(logLevelHeader, 10, 32)
	if err != nil || u64 > 7 {
		c.Logger().Errorf("Invalid log level in header %s. Please, provide a int between 0 and 7", logLevelHeader)
		return c.Logger().Level()
	}

	return log.Lvl(uint8(u64))
}
