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

	resourceResolver := ResourceResolvers.ResourceResolverWithBaseHRefResolverOptions(
		ResourceResolvers.BaseHRefResolverOptions{
			Version: version,
		},
	)

	// set up flags.  generate_static, serve static, generate_serve
	generateStatic := flag.Bool("generate_static", false, "Generate static website")
	flag.Parse()

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
		InternalURLs:       []string{"/demo1/"},
	}

	if *generateStatic {
		err := app.GenerateStaticWebsite("static_output/demo1", appHandler)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Static site generated in static_output/demo1 (fixed paths)")
		return
	}

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

}
