package internal_test

import (
	"reflect"
	"testing"

	"github.com/iarie/rechallenge/data"
	"github.com/iarie/rechallenge/internal"
)

func TestPackerV1_Input(t *testing.T) {
	_, err := internal.PackerV1(-1, []data.Package{})

	if err == nil {
		t.Errorf("Vaildation Failed")
	}
}

func TestPackerV1_Required(t *testing.T) {
	pack250 := data.Package{Sku: "A", Size: 250}
	pack500 := data.Package{Sku: "A", Size: 500}
	pack1000 := data.Package{Sku: "A", Size: 1000}
	pack2000 := data.Package{Sku: "A", Size: 2000}
	pack5000 := data.Package{Sku: "A", Size: 5000}

	inventory := []data.Package{pack5000, pack2000, pack1000, pack500, pack250}

	cases := []struct {
		qty         int
		expectation data.Order
	}{
		{0, data.Order{}},
		{1, data.Order{
			LineItems: []data.LineItem{
				{Package: pack250, Qty: 1},
			},
		}},
		{250, data.Order{
			LineItems: []data.LineItem{
				{Package: pack250, Qty: 1},
			},
		}},
		{251, data.Order{
			LineItems: []data.LineItem{
				{Package: pack500, Qty: 1},
			},
		}},
		{501, data.Order{
			LineItems: []data.LineItem{
				{Package: pack500, Qty: 1},
				{Package: pack250, Qty: 1},
			},
		}},
		{12001, data.Order{
			LineItems: []data.LineItem{
				{Package: pack5000, Qty: 2},
				{Package: pack2000, Qty: 1},
				{Package: pack250, Qty: 1},
			},
		}},
	}

	for i, test := range cases {
		r, err := internal.PackerV1(test.qty, inventory)

		if err != nil {
			t.Errorf("Unexpected Error [case#%v]: %v. Want: %v", i+1, err, test.expectation)
		}

		eq := reflect.DeepEqual(r, test.expectation)

		if !eq {
			t.Errorf("Bad Result [case#%v]: %v. Want: %v", i+1, r, test.expectation)
		}
	}
}

func TestPackerV1_Custom(t *testing.T) {
	pack10 := data.Package{Sku: "S", Size: 10}
	pack5 := data.Package{Sku: "S", Size: 5}
	pack3 := data.Package{Sku: "S", Size: 3}

	inventory := []data.Package{pack10, pack5, pack3}

	cases := []struct {
		qty         int
		expectation data.Order
	}{
		{0, data.Order{}},
		{15, data.Order{
			LineItems: []data.LineItem{
				{Package: pack10, Qty: 1},
				{Package: pack5, Qty: 1},
			},
		}},
	}

	for i, test := range cases {
		r, err := internal.PackerV1(test.qty, inventory)

		if err != nil {
			t.Errorf("Unexpected Error [case#%v]: %v. Want: %v", i+1, err, test.expectation)
		}

		eq := reflect.DeepEqual(r, test.expectation)

		if !eq {
			t.Errorf("Bad Result [case#%v]: %v. Want: %v", i+1, r, test.expectation)
		}
	}
}
