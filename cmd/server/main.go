package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Root selector page
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
<!DOCTYPE html>
<html>
<head>
	<title>Select Demo</title>
	<style>
		body { font-family: sans-serif; padding: 2rem; }
		a { display: inline-block; padding: 10px 20px; background: #007bff; color: white; text-decoration: none; border-radius: 5px; }
		a:hover { background: #0056b3; }
	</style>
</head>
<body>
	<h1>Select Demo</h1>
	<p><a href="/demo1/">Go to Demo 1</a></p>
</body>
</html>
`)
	})

	// Serve the WASM file from the web folder for demo1
	e.GET("/demo1/web/app.wasm", func(c echo.Context) error {
		return c.File("web/app.wasm")
	})

	// Serve static files from static_output/demo1 at /demo1
	e.Static("/demo1", "static_output/demo1")

	// SPA fallback for demo1 routes (must come after Static)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				httpErr, ok := err.(*echo.HTTPError)
				if ok && httpErr.Code == 404 {
					path := c.Request().URL.Path
					// Only handle 404s under /demo1/
					if len(path) >= 7 && path[:7] == "/demo1/" {
						return c.File("static_output/demo1/index.html")
					}
				}
			}
			return err
		}
	})

	e.Logger.Fatal(e.Start(":3556"))
}
