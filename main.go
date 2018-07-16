package main

import (
	"net/http"
	
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
  e.Static("/ripple/static", "./public/static")
	e.GET("/ripple/", func(c echo.Context) error {
    content := create_json()
		return c.HTML(http.StatusOK, content)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
