package main

import (
	"github.com/iarie/rechallenge/app"
)

func main() {
	// load config
	cfg := app.NewConfig(8080)
	// init store
	// start web server

	app.Run(cfg)
}
