package main

import (
	"net/http"
  "fmt"
	"github.com/tkanos/gonfig"
	"github.com/labstack/echo"
)

type Configuration struct {
    Port              int
    Name              string
    ContextPath       string
}



func main() {
  config := Configuration{}
  err := gonfig.GetConf("conf/docker.config.json", &config)  
  fmt.Println(err)
	e := echo.New()
  frontend := e.Group(config.ContextPath)
  frontend.Static("/static", "./public/static")
	frontend.GET("/", func(c echo.Context) error {
    content := createJson()
		return c.HTML(http.StatusOK, content)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
