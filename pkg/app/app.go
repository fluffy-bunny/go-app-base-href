package app

import (
	"strings"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Home struct {
	app.Compo
	Name string
}

func (h *Home) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Home Page - "+h.Name),
		app.A().Href(fixHRef("/page1")).Text("Go to Page 1"),
	)
}

type Page1 struct {
	app.Compo
}

func fixHRef(href string) string {
	rootPrefix := app.Getenv("GOAPP_ROOT_PREFIX")

	if rootPrefix == "" || rootPrefix == "/" {
		return href
	}
	rr := strings.TrimRight(rootPrefix, "/")
	return rr + href

}
func (p *Page1) Render() app.UI {

	return app.Div().Body(
		app.H1().Text("Page 1"),
		app.A().Href(fixHRef("/")).Text("Back to Home"),
	)
}

func SetupRoutes(name string) {
	fixPath := func(path string) string {

		return path
	}
	app.Route(fixPath("/"), func() app.Composer { return &Home{Name: name} })
	app.Route(fixPath("/page1"), func() app.Composer { return &Page1{} })
}
