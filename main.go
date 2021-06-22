package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kazix-mrs/factory"
	"github.com/labstack/echo"
)

func main() {
	appEnv := flag.String("app-env", os.Getenv("APP_HOME"), "app env")
	flag.Parse()
	e := echo.New()

	if *appEnv == "" {
		*appEnv = `./`
	}

	fac := factory.Factory{JsonConfigPath: *appEnv}

	fac.Initialize()

	e.GET("/say", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})
	e.POST("/echo", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "success")
	})

	e.Logger.Fatal(e.Start(":2020"))
	fmt.Print("sdf")

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		for {
			signal := <-signalChan

			switch signal {

			case syscall.SIGHUP:
				os.Exit(0)
			case syscall.SIGINT:
				os.Exit(0)
			case syscall.SIGTERM:
				os.Exit(0)
			default:
				//				os.Exit(0)
			}
		}
	}()
}
