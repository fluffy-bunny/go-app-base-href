package app

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Home struct {
	app.Compo
	Name string
}

func (h *Home) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Home Page - "+h.Name),
		app.A().Href("page1").Text("Go to Page 1"),
	)
}

type Page1 struct {
	app.Compo
}

func (p *Page1) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Page 1"),
		app.A().Href("/").Text("Back to Home"),
	)
}

func SetupRoutes(name string, baseRef string) {
	fixPath := func(path string) string {
		if baseRef != "" {
			return "/" + baseRef + path
		}
		return path
	}
	app.Route(fixPath("/"), func() app.Composer { return &Home{Name: name} })
	app.Route(fixPath("/page1"), func() app.Composer { return &Page1{} })
}
