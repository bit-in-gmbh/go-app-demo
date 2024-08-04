package main

import (
	"fmt"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type hello struct {
	app.Compo
	name string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Hello, "),
			app.If(h.name != "", func() app.UI {
				return app.Text(h.name)
			}).Else(func() app.UI {
				return app.Text("World")
			}),
		),
		app.Input().
			Value(h.name).
			Placeholder("What is your name?").
			AutoFocus(true).
			OnInput(h.OnInputChange),
	)
}

func (h *hello) OnInputChange(ctx app.Context, e app.Event) {
	h.name = ctx.Src().JSValue().Get("value").String()
	h.Render()
}

func main() {
	fmt.Println("starting local server")

	app.Route("/", app.NewZeroComponentFactory(&hello{}))
	app.RunWhenOnBrowser()

	h := &app.Handler{
		Title:  "Hello Demo",
		Author: "Maxence Charriere",
	}

	if err := http.ListenAndServe(":7000", h); err != nil {
		panic(err)
	}
}
