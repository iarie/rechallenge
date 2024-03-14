package main

import (
	"github.com/iarie/rechallenge/app"
)

func main() {
	cfg := app.NewConfig(
		app.UsePort(8080),
		app.UsePacker("V1"),
	)

	app.Run(cfg)
}
