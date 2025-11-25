package main

import (
	my_app "demo-project/pkg/app"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {

	my_app.SetupRoutes("demo1")
	app.RunWhenOnBrowser()
}
