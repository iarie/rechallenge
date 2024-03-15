package internal

import "github.com/iarie/rechallenge/data"

type Repository interface {
	Get() []data.Package
}

type HardcodedInventoryRepo struct{}

func (r *HardcodedInventoryRepo) Get() []data.Package {
	return []data.Package{
		{Sku: "xxxx0250", Size: 250},
		{Sku: "xxxx0500", Size: 500},
		{Sku: "xxxx1000", Size: 1000},
		{Sku: "xxxx2000", Size: 2000},
		{Sku: "xxxx5000", Size: 5000},
	}
}
