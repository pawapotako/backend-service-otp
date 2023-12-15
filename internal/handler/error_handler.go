package handler

import (
	"net/http"

	"m2ex-otp-service/internal/util"

	"github.com/labstack/echo/v4"
)

func errorHandler(c echo.Context, err error) error {
	switch e := err.(type) {
	case util.AppErrors:
		c.JSON(e.Errors[0].Status, e)
	case error:
		c.JSON(http.StatusInternalServerError, e)
	}

	return nil
}
