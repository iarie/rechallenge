package internal

import (
	"errors"
	"fmt"
	"sort"

	"github.com/iarie/rechallenge/data"
)

type Packer interface {
	Pack(int, []data.Package) (data.Order, error)
}

type PackerFunc func(int, []data.Package) (data.Order, error)

func (f PackerFunc) Pack(qty int, packs []data.Package) (data.Order, error) {
	return f(qty, packs)
}

func PackerV1(total int, packs []data.Package) (data.Order, error) {
	o := data.Order{}

	// Simple validations
	if total < 0 {
		err := errors.New("negative total value")
		return o, err
	}

	if total == 0 {
		return o, nil
	}

	// Shared var to aggregate solutions
	mem := []Solution{}

	// Initial empty solution
	start := newSolution([]data.LineItem{}, total)

	// Start search
	recursionV1(start, total, packs, &mem)

	// Rank results
	sort.Sort(ByScore(mem))

	fmt.Println("Permuations:", len(mem))

	// Build Order
	o.LineItems = append(o.LineItems, mem[0].Items...)
	return o, nil
}

// Search for permuations
// Accept rolling solution node, remainder to fullfil, packs to iterate, and memo to store results
func recursionV1(in Solution, remainder int, packs []data.Package, mem *[]Solution) {
	// Algorithm assumes that packs slice is sorted by size (desc)
	// Start from the biggest pack
	for i, p := range packs {

		// Calculate how many packs is needed to fulfill an order
		// Assuming we have unlimited stock, otherwise we can limit q value
		q := (remainder / p.Size)

		// Check combinations starting with max fulfilling to 0 packs
		for j := q + 1; j >= 0; j-- {
			// Spawn new solution branch:
			// copy rolling line items
			pNext := make([]data.LineItem, len(in.Items))
			copy(pNext, in.Items)

			// Do not append null package
			if j != 0 {
				pNext = append(pNext, data.LineItem{Package: p, Qty: j})
			}

			// Clone stuct
			sNext := newSolution(pNext, in.Target)

			// Save solved solution & continue
			if sNext.IsSolved() {
				// call Finalize to assign the Score
				sNext.Finalize()
				*mem = append(*mem, sNext)

				continue
			}

			// Calculate new remainder
			newRemainder := remainder - p.Size*j

			// Go to next smaller slice: packs[i+1:]
			recursionV1(sNext, newRemainder, packs[i+1:], mem)
		}
	}
}

// Ranking functions

type Solution struct {
	Items  []data.LineItem
	Target int
	Score  [2]int

	Debug string
}

func newSolution(items []data.LineItem, target int) Solution {
	return Solution{
		Items:  items,
		Target: target,
	}
}

func (s *Solution) CheckScore() [2]int {
	return getScore(s.Items, s.Target)
}

func (s *Solution) Finalize() {
	s.Score = getScore(s.Items, s.Target)
}

func (s *Solution) IsSolved() bool {
	sc := s.CheckScore()
	return sc[0] <= 0
}

type ByScore []Solution

func (a ByScore) Len() int {
	return len(a)
}

func (a ByScore) Less(i, j int) bool {
	if a[i].Score[0] == a[j].Score[0] {
		return a[i].Score[1] < a[j].Score[1]
	} else {
		return a[i].Score[0] > a[j].Score[0]
	}
}

func (a ByScore) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func getScore(items []data.LineItem, target int) [2]int {
	var totalItems int
	var totalPacks int

	for _, li := range items {
		totalItems += li.Qty * li.Size
		totalPacks += li.Qty
	}

	return [2]int{target - totalItems, totalPacks}
}
