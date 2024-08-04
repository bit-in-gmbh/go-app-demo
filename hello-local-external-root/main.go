package main

import (
	"fmt"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	fmt.Println("starting local server with external root")

	h := &app.Handler{
		Title:  "Hello Demo with external root",
		Author: "Maxence Charriere",
		// RootDir: "../hello-local/web",
		Resources: app.LocalDir("../hello-local/web"),
	}

	if err := http.ListenAndServe(":7000", h); err != nil {
		panic(err)
	}
}
