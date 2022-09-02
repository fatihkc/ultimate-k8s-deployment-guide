package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {

	user := os.Getenv("USER")
	secret, err := os.ReadFile("/tmp/secret-file.txt")
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		hello := "Hello, " + user + "!" + " Your secret is: " + string(secret)
		return c.String(http.StatusOK, hello)
	})
	e.GET("/health", healthCheck)
	e.Logger.Fatal(e.Start(":8080"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Service is healthy!")
}
