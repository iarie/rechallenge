package main

import (
	"github.com/iarie/rechallenge/app"
	"github.com/iarie/rechallenge/data"
	"github.com/iarie/rechallenge/internal"
)

func main() {
	repo := &internal.HardcodedInventoryRepo{
		Data: []data.Package{
			{Sku: "A", Size: 2000},
			{Sku: "A", Size: 500},
			{Sku: "A", Size: 1000},
			{Sku: "A", Size: 250},
			{Sku: "A", Size: 5000},
		},
	}

	cfg := app.NewConfig(
		app.UsePort(80),
		app.UsePacker("V2"),
		app.UseInventoryRepo(repo),
	)

	app.Run(cfg)
}
