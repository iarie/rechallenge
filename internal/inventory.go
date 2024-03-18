package internal

import (
	"sort"

	"github.com/iarie/rechallenge/data"
)

type Repository interface {
	Get() []data.Package
	New(int) error
	Delete(int) error
}

type HardcodedInventoryRepo struct {
	Data []data.Package
}

func (r *HardcodedInventoryRepo) Get() []data.Package {
	inv := r.Data

	sort.Sort(data.BySizeDesc(inv))

	return inv
}

func (r *HardcodedInventoryRepo) New(size int) error {
	r.Data = append(r.Data, data.Package{Sku: "A", Size: size})
	return nil
}

func (r *HardcodedInventoryRepo) Delete(size int) error {
	var indexToDelete int

	for i, d := range r.Data {
		if d.Size == size {
			indexToDelete = i
		}
	}

	r.Data = append(r.Data[:indexToDelete], r.Data[indexToDelete+1:]...)
	return nil
}
