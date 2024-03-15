package internal

import (
	"fmt"
	"sort"

	"github.com/iarie/rechallenge/data"
)

type Packer interface {
	Pack(int, []data.Package) data.Order
}

type PackerFunc func(int, []data.Package) data.Order

func (f PackerFunc) Pack(qty int, packs []data.Package) data.Order {
	return f(qty, packs)
}

func PackerV1(total int, packs []data.Package) data.Order {
	o := data.Order{}

	if total == 0 {
		return o
	}

	// Search for permuations

	mem := []Solution{}
	start := newSolution([]data.LineItem{}, total)
	recursionV1(start, total, packs, &mem)

	// Rank results

	sort.Sort(ByScore(mem))

	// for i, r := range mem {
	// 	// fmt.Printf("# %v - Items: %v Taget: %v Score: %v\n", i+1, r.Items, r.Target, r.Score)
	// 	// fmt.Printf("DEBUG: %v\n", r.Debug)
	// }

	// Build Order

	o.LineItems = append(o.LineItems, mem[0].Items...)
	return o
}

func recursionV1(in Solution, remainder int, packs []data.Package, mem *[]Solution) {
	if remainder == 0 {
		return
	}

	for i, p := range packs {
		q := (remainder / p.Size)

		pMax := make([]data.LineItem, len(in.Items))
		copy(pMax, in.Items)
		pMax = append(pMax, data.LineItem{Package: p, Qty: q + 1})
		sMax := newSolution(pMax, in.Target)

		// save permutation
		if sMax.IsSolved() {
			debug := fmt.Sprintf("isSolved: %v Target:%v Score:%v", sMax.IsSolved(), sMax.Target, sMax.CheckScore())
			sMax.Finalize(debug)
			*mem = append(*mem, sMax)
		}

		// go to next pack
		if q == 0 {
			continue
		}

		// Check combinations starting with max fulfilling to 0 packs
		for j := q; j >= 0; j-- {
			pNext := make([]data.LineItem, len(in.Items))
			copy(pNext, in.Items)

			// if j != 0 {
			pNext = append(pNext, data.LineItem{Package: p, Qty: j})
			// }

			sNext := newSolution(pNext, in.Target)

			fulfilled := p.Size * j

			if sNext.IsSolved() {
				debug := fmt.Sprintf("isSolved: %v Target:%v Score:%v", sNext.IsSolved(), sNext.Target, sNext.CheckScore())
				sNext.Finalize(debug)
				*mem = append(*mem, sNext)
			}

			var newRemainder int

			if fulfilled != 0 {
				newRemainder = remainder % fulfilled
			} else {
				newRemainder = remainder
			}

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

func (s *Solution) Finalize(debug string) {
	s.Score = getScore(s.Items, s.Target)
	s.Debug = debug
}

func (s *Solution) IsSolved() bool {
	sc := s.CheckScore()
	// fmt.Printf("IsSolved: %v | %v\n", sc, sc[0] <= 0)
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
