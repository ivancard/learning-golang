package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// echo framework
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Jellou guorld")
	})
	
	e.Logger.Fatal(e.Start(":1323"))
}
