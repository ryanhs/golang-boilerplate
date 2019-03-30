package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("# ECHO - HELLO WORLD!")

	e := echo.New()

	e.GET("/foo", func(c echo.Context) error {
		return c.String(http.StatusOK, "bar")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.Start(":3000")
}
