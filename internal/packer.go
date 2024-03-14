package internal

import "github.com/iarie/rechallenge/data"

type Packer interface {
	Pack(int, []data.Package) data.Order
}

type PackerFunc func(int, []data.Package) data.Order

func (f PackerFunc) Pack(qty int, packs []data.Package) data.Order {
	return f(qty, packs)
}

func PackerV1(qty int, packs []data.Package) data.Order {
	pkg_50 := data.Package{Sku: "xxxx0200", Size: 50}
	pkg_100 := data.Package{Sku: "xxxx0200", Size: 100}

	o := data.Order{}

	l1 := data.LineItem{Package: pkg_50, Qty: 5}
	l2 := data.LineItem{Package: pkg_100, Qty: 3}
	o.LineItems = append(o.LineItems, l1, l2)

	return o
}
