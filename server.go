package main

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	User struct {
		Name  string `query:"name" validate:"required"`
		Email string `query:"email" validate:"required,email"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Validator = &CustomValidator{validator: validator.New()}
	e.GET("/", func(c echo.Context) error {
		return errors.New("omg, an error")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
