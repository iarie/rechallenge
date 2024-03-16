package internal

import (
	"errors"
	"math"

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
	o := data.Order{Requested: total}

	// Simple validations
	if total < 0 {
		err := errors.New("negative total value")
		return o, err
	}

	if total == 0 {
		return o, nil
	}

	// Shared var to aggregate solutions
	mem := &Solution{Target: total, Score: [2]int{math.MinInt, math.MaxInt}}

	// Initial empty solution
	start := newSolution([]data.LineItem{}, total)

	// Start search
	recursionV1(start, total, packs, mem)

	// Build Order
	o.LineItems = append(o.LineItems, mem.Items...)
	return o, nil
}

// Search for permuations
// Accept rolling solution node, remainder to fullfil, packs to iterate, and memo to store best
func recursionV1(in Solution, remainder int, packs []data.Package, mem *Solution) {
	if len(packs) == 0 {
		return
	}

	// Shift biggest pack
	p := packs[0]

	// Calculate how many packs needed to fulfill an order
	q := (remainder / p.Size)

	// Check combinations starting with max fulfilling to 1 packs
	for j := q + 1; j >= 0; j-- {
		// Spawn new branch solution
		sNext := copySoltion(in)
		if j != 0 {
			sNext.Items = append(sNext.Items, data.LineItem{Package: p, Qty: j})
		}

		if sNext.IsSolved() {
			// Call Finalize to assign the Score
			sNext.Finalize()

			// Save solved solution
			if isBetterSolution(sNext, *mem) {
				*mem = sNext
			}
		}

		// if we have a precise solution skip branches with worse secondary score
		if mem.Score[0] == 0 && len(sNext.Items) >= len(mem.Items) {
			continue
		}

		// Calculate new remainder
		newRemainder := remainder - p.Size*j

		// Go to next smaller item and repeat
		recursionV1(sNext, newRemainder, packs[1:], mem)
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

func getScore(items []data.LineItem, target int) [2]int {
	var totalItems int
	var totalPacks int

	for _, li := range items {
		totalItems += li.Qty * li.Size
		totalPacks += li.Qty
	}

	return [2]int{target - totalItems, totalPacks}
}

func copySoltion(s Solution) Solution {
	pNext := make([]data.LineItem, len(s.Items))
	copy(pNext, s.Items)
	return newSolution(pNext, s.Target)
}

func isBetterSolution(next, prev Solution) bool {
	ns := next.Score
	ps := prev.Score

	if ns[0] == ps[0] {
		return ns[1] < ps[1]
	}
	return ns[0] > ps[0]
}
