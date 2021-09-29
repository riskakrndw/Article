package main

import (
	"fmt"
	"project/pasarwarga/config"
	"project/pasarwarga/routes"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	config.InitDb()
	config.InitPort()
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
