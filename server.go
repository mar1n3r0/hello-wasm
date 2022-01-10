// Our empty version of the httpServer for usage with the wasm target
// this way we will not include any of the related code
//go:build !wasm

package main

import (
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func AppServer() {
	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	withGz := gziphandler.GzipHandler(&app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})
	http.Handle("/", withGz)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
