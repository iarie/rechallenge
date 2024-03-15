package main

import (
	"github.com/iarie/rechallenge/app"
	"github.com/iarie/rechallenge/internal"
)

func main() {
	cfg := app.NewConfig(
		app.UsePort(80),
		app.UsePacker("V1"),
		app.UseInventoryRepo(&internal.HardcodedInventoryRepo{}),
	)

	app.Run(cfg)
}
