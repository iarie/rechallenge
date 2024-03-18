package internal_test

import (
	"reflect"
	"testing"

	"github.com/iarie/rechallenge/data"
	"github.com/iarie/rechallenge/internal"
)

func TestPackerV2_Input(t *testing.T) {
	_, err := internal.PackerV2(-1, []data.Package{})

	if err == nil {
		t.Errorf("Vaildation Failed")
	}
}

func TestPackerV2_Required(t *testing.T) {
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
		{0, data.Order{Requested: 0}},
		{1, data.Order{
			Requested: 1,
			LineItems: []data.LineItem{
				{Package: pack250, Qty: 1},
			},
		}},
		{250, data.Order{
			Requested: 250,
			LineItems: []data.LineItem{
				{Package: pack250, Qty: 1},
			},
		}},
		{251, data.Order{
			Requested: 251,
			LineItems: []data.LineItem{
				{Package: pack500, Qty: 1},
			},
		}},
		{501, data.Order{
			Requested: 501,
			LineItems: []data.LineItem{
				{Package: pack500, Qty: 1},
				{Package: pack250, Qty: 1},
			},
		}},
		{12001, data.Order{
			Requested: 12001,
			LineItems: []data.LineItem{
				{Package: pack5000, Qty: 2},
				{Package: pack2000, Qty: 1},
				{Package: pack250, Qty: 1},
			},
		}},
	}

	for i, test := range cases {
		r, err := internal.PackerV2(test.qty, inventory)

		if err != nil {
			t.Errorf("Unexpected Error [case#%v]: %v. Want: %v", i+1, err, test.expectation)
		}

		eq := reflect.DeepEqual(r, test.expectation)

		if !eq {
			t.Errorf("Bad Result [case#%v]: %v. Want: %v", i+1, r, test.expectation)
		}
	}
}

func TestPackerV2_Custom(t *testing.T) {
	pack20 := data.Package{Sku: "S", Size: 20}
	pack10 := data.Package{Sku: "S", Size: 10}
	pack5 := data.Package{Sku: "S", Size: 5}
	pack3 := data.Package{Sku: "S", Size: 3}

	inventory := []data.Package{pack20, pack10, pack5, pack3}

	cases := []struct {
		qty         int
		expectation data.Order
	}{
		{0, data.Order{Requested: 0}},
		{15, data.Order{
			Requested: 15,
			LineItems: []data.LineItem{
				{Package: pack10, Qty: 1},
				{Package: pack5, Qty: 1},
			},
		}},
	}

	for i, test := range cases {
		r, err := internal.PackerV2(test.qty, inventory)

		if err != nil {
			t.Errorf("Unexpected Error [case#%v]: %v. Want: %v", i+1, err, test.expectation)
		}

		eq := reflect.DeepEqual(r, test.expectation)

		if !eq {
			t.Errorf("Bad Result [case#%v]: %v. Want: %v", i+1, r, test.expectation)
		}
	}
}

func TestPackerV2_PrimalPacks(t *testing.T) {
	pack17 := data.Package{Sku: "S", Size: 17}
	pack7 := data.Package{Sku: "S", Size: 7}
	pack5 := data.Package{Sku: "S", Size: 5}
	pack3 := data.Package{Sku: "S", Size: 3}

	inventory := []data.Package{pack17, pack7, pack5, pack3}

	cases := []struct {
		qty         int
		expectation data.Order
	}{
		{0, data.Order{Requested: 0}},
		{42, data.Order{
			Requested: 42,
			LineItems: []data.LineItem{
				{Package: pack17, Qty: 2},
				{Package: pack5, Qty: 1},
				{Package: pack3, Qty: 1},
			},
		}},
		{88, data.Order{
			Requested: 88,
			LineItems: []data.LineItem{
				{Package: pack17, Qty: 5},
				{Package: pack3, Qty: 1},
			},
		}},
		{1256, data.Order{
			Requested: 1256,
			LineItems: []data.LineItem{
				{Package: pack17, Qty: 73},
				{Package: pack5, Qty: 3},
			},
		}},
		{1_000_000, data.Order{
			Requested: 1_000_000,
			LineItems: []data.LineItem{
				{Package: pack17, Qty: 58822},
				{Package: pack7, Qty: 3},
				{Package: pack5, Qty: 1},
			},
		}},
	}

	for i, test := range cases {
		r, err := internal.PackerV2(test.qty, inventory)

		if err != nil {
			t.Errorf("Unexpected Error [case#%v]: %v. Want: %v", i+1, err, test.expectation)
		}

		eq := reflect.DeepEqual(r, test.expectation)

		if !eq {
			t.Errorf("Bad Result [case#%v]: %v. Want: %v", i+1, r, test.expectation)
		}
	}
}
