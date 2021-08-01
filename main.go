package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.GET("/hello/v1/sayHello", func(c echo.Context) error {
		log.Println("my hello log")
		fmt.Println("log backup")
		return c.String(http.StatusOK, "Hello, World from v1! version 0.0.4")
	})
	e.GET("/hello/v1/sayHello/toMe", func(c echo.Context) error {
		log.Println("my hello log")
		fmt.Println("log backup")
		return c.String(http.StatusOK, "Hello, to me!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}