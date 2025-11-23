package main

import (
	"flag"
	"log"

	"demo-project/pkg/ResourceResolvers"
	my_app "demo-project/pkg/app"

	"github.com/rs/xid"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	my_app.SetupRoutes("demo1", "")
	version := xid.New().String()

	// set up flags.  generate_static, serve static, generate_serve
	generateStatic := flag.Bool("generate_static", false, "Generate static website")
	flag.Parse()

	if *generateStatic {
		resourceResolver := ResourceResolvers.ResourceResolverWithBaseHRefResolverOptions(
			ResourceResolvers.BaseHRefResolverOptions{
				Version: version,
			},
		)
		appHandler := &app.Handler{
			Name:        "Demo1",
			Title:       "Demo 1",
			Description: "A demo application hosted in a subfolder",
			Icon:        app.Icon{},
			RawHeaders: []string{
				"<base href=\"/demo1/\">",
			},
			Resources:          resourceResolver,
			Styles:             []string{},
			Scripts:            []string{},
			CacheableResources: []string{},
		}

		err := app.GenerateStaticWebsite("static_output/demo1", appHandler)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Static site generated in static_output/demo1 (fixed paths)")
		return
	}

}
