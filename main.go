package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Ras96/auth-practice/model"
	sess "github.com/Ras96/auth-practice/session"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := model.EstablishConnection()
	if err != nil {
		panic(err)
	}

	store, err := sess.NewSession(db.DB)
	if err != nil {
		panic(fmt.Errorf("failed in session constructor:%v", err))
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(store))

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.POST("/login", model.PostLoginHandler)
	e.POST("/signup", model.PostSignUpHandler)

	todo := e.Group("/todo")
	todo.GET("", model.GetTodoHandler)
	todo.POST("", model.PostTodoHandler)

	withLogin := e.Group("")
	withLogin.Use(model.CheckLogin)
	withLogin.GET("/cities/:cityName", model.GetCityInfoHandler)
	withLogin.GET("/countries", model.GetCountriesHandler)
	withLogin.GET("/countries/:countryName", model.GetCountryInfoHandler)
	withLogin.GET("/whoami", model.GetWhoAmIHandler)

	e.Start(":4000")
}
