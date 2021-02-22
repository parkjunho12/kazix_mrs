package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.GET("/say", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hi")
	})
	e.Logger.Fatal(e.Start(":2021"))
	fmt.Print("sdf")
}