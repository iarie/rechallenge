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
	LineItems []LineItem
}
