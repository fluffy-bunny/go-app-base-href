package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve the WASM file from the web folder
	e.GET("/web/app.wasm", func(c echo.Context) error {
		return c.File("web/app.wasm")
	})

	// Serve static files from static_output/base_line
	e.Static("/", "static_output/base_line")

	// SPA fallback for non-file routes (must come after Static)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				httpErr, ok := err.(*echo.HTTPError)
				if ok && httpErr.Code == 404 {
					// Return index.html for 404s to support SPA routing
					return c.File("static_output/base_line/index.html")
				}
			}
			return err
		}
	})

	e.Logger.Fatal(e.Start(":3557"))
}
