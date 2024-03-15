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
