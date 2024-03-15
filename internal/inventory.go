package internal

import (
	"sort"

	"github.com/iarie/rechallenge/data"
)

type Repository interface {
	Get() []data.Package
}

type HardcodedInventoryRepo struct{}

func (r *HardcodedInventoryRepo) Get() []data.Package {
	inv := []data.Package{
		{Sku: "xxxx2000", Size: 2000},
		{Sku: "xxxx0500", Size: 500},
		{Sku: "xxxx1000", Size: 1000},
		{Sku: "xxxx0250", Size: 250},
		{Sku: "xxxx5000", Size: 5000},
	}

	sort.Sort(bySize(inv))

	return inv
}

type bySize []data.Package

func (a bySize) Len() int {
	return len(a)
}

func (a bySize) Less(i, j int) bool {
	return a[i].Size > a[j].Size
}

func (a bySize) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
