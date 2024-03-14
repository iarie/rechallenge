package internal

import "github.com/iarie/rechallenge/data"

type Packer interface {
	Pack(int) data.Order
}

type PackerFunc func(int) data.Order

func (f PackerFunc) Pack(qty int) data.Order {
	return f(qty)
}

func PackerV1(qty int) data.Order {
	pkg_50 := data.Package{Sku: "xxxx0200", Size: 50}
	pkg_100 := data.Package{Sku: "xxxx0200", Size: 100}

	o := data.Order{}

	l1 := data.LineItem{Package: pkg_50, Qty: 5}
	l2 := data.LineItem{Package: pkg_100, Qty: 3}
	o.LineItems = append(o.LineItems, l1, l2)

	return o
}
