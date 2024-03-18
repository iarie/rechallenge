package data

type Package struct {
	Sku  string
	Size int
}

type LineItem struct {
	Package
	Qty int
}

type Order struct {
	Requested int
	LineItems []LineItem
}

func (o Order) NotEmpty() bool {
	return len(o.LineItems) > 0
}

type BySizeAsc []Package

func (a BySizeAsc) Len() int {
	return len(a)
}

func (a BySizeAsc) Less(i, j int) bool {
	return a[i].Size < a[j].Size
}

func (a BySizeAsc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type BySizeDesc []Package

func (a BySizeDesc) Len() int {
	return len(a)
}

func (a BySizeDesc) Less(i, j int) bool {
	return a[i].Size > a[j].Size
}

func (a BySizeDesc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
