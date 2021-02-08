package model

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Country struct {
	Code sql.NullString `json:"code,omitempty"  db:"Code"`
	Name sql.NullString `json:"name,omitempty" db:"Name"`
}

func GetCountriesHandler(c echo.Context) error {
	countries := []Country{}

	err := db.Select(&countries, "SELECT Name FROM country")
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}

	result := []string{}
	for _, v := range countries {
		if v.Name.Valid {
			result = append(result, v.Name.String)
		}
	}
	return c.JSON(http.StatusOK, result)
}

func GetCountryInfoHandler(c echo.Context) error {
	countryName := c.Param("countryName")

	country := Country{}
	cities := []City{}

	err := db.Get(&country, "SELECT Code FROM country WHERE Name = ? LIMIT 1", countryName)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}

	err = db.Select(&cities, "SELECT Name FROM city WHERE CountryCode = ?", country.Code)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}

	result := []string{}
	for _, v := range cities {
		if v.Name.Valid {
			result = append(result, v.Name.String)
		}
	}
	return c.JSON(http.StatusOK, result)
}
