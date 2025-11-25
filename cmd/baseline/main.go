package main

import (
	my_app "demo-project/pkg/app"
	"flag"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {

	// set up flags.  generate_static, serve static, generate_serve
	generateStatic := flag.Bool("generate_static", false, "Generate static website")
	flag.Parse()

	my_app.SetupRoutes("demo1")
	app.RunWhenOnBrowser()
	appHandler := &app.Handler{
		Name:        "Demo1",
		Description: "A demo application",
	}
	// set up routes
	http.Handle("/", appHandler)

	if *generateStatic {
		err := app.GenerateStaticWebsite("static_output/base_line", appHandler)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Static site generated in static_output/base_line (fixed paths)")
		return
	}
	if err := http.ListenAndServe(":3556", nil); err != nil {
		log.Fatal(err)
	}
}
