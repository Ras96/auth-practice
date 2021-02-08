package model

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Me struct {
	Username string `json:"username,omitempty"  db:"username"`
}

func GetWhoAmIHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Me{
		Username: c.Get("userName").(string),
	})
}
