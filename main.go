package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stobita/go-echo-vue/model"
)

func main() {
	e := echo.New()
	// ログ出力
	e.Use(middleware.Logger())
	// Access-Control-Allow-Originの対策
	e.Use(middleware.CORS())

	e.GET("/hello/", model.Hello)
	e.POST("/users/", model.InsertUser)

	e.Logger.Fatal(e.Start(":1323"))
}
