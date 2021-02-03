package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.Secure(),
		middleware.RequestID(),
		middleware.BodyLimit("1M"),
		middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}),
	)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "I'm offset.")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"code": "OK", "message": "I'm offset."})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
