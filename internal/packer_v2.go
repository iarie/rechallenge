package internal

import (
	"errors"
	"math"
	"sort"

	"github.com/iarie/rechallenge/data"
)

func PackerV2(total int, packs []data.Package) (data.Order, error) {
	o := data.Order{Requested: total}

	// Simple validations
	if total < 0 {
		err := errors.New("negative total value")
		return o, err
	}

	if total == 0 {
		return o, nil
	}

	// sort Packs
	sort.Sort(data.BySizeAsc(packs))

	res := search(total, packs)

	// Build Order
	// from biggest pack to lowest
	for i := len(res.Packs) - 1; i >= 0; i-- {
		q := res.Packs[i]
		if q == 0 {
			continue
		}
		newItem := data.LineItem{Package: packs[i], Qty: q}
		o.LineItems = append(o.LineItems, newItem)
	}

	return o, nil
}

func search(total int, Packs []data.Package) SolutionV2 {
	// Init solutions matrix
	solutions := make([][]SolutionV2, total+1)
	for i := range solutions {
		solutions[i] = make([]SolutionV2, len(Packs)+1)
		solutions[i][0] = SolutionV2{Score: math.MinInt}
	}

	for i := range solutions[0] {
		solutions[0][i] = SolutionV2{Score: math.MinInt}
	}

	var (
		s0 SolutionV2
		s1 SolutionV2
		s2 SolutionV2
		sr SolutionV2
	)

	for n := 1; n <= total; n++ {
		for c, pack := range Packs {
			val := pack.Size

			if n < val {
				// prev best solution without current pack
				s0 = solutions[n][c]

				// prev solution using current pack for n-1
				s1 = solutions[n-1][c+1]
				// adjust score
				s1.Save(n)

				// use pack
				s2 = SolutionV2{Score: n - val, Sum: val, Qty: 1, Packs: make([]int, len(Packs))}
				s2.Packs[c] = 1

				// save best of 3
				solutions[n][c+1] = findBest(s0, s1, s2)
			} else if n == val {
				// use pack
				s1 = SolutionV2{Score: 0, Sum: val, Qty: 1, Packs: make([]int, len(Packs))}
				s1.Packs[c] = 1
				solutions[n][c+1] = s1
			} else if n > val {
				// prev best solution without current pack
				s0 = solutions[n][c]

				// prev solution using current pack for n-1
				s1 = solutions[n-1][c+1]
				// adjust score
				s1.Save(n)

				// prev best soltion that fit current pack size
				sr = solutions[n-val][c+1]
				s2 = SolutionV2{Packs: make([]int, len(Packs))}
				copy(s2.Packs, sr.Packs)
				s2.Packs[c] += 1
				s2.Sum = sr.Sum + val
				s2.Qty = sr.Qty + 1
				s2.Save(n)

				// save best of 3
				solutions[n][c+1] = findBest(s0, s1, s2)
			}
		}
	}

	return solutions[total][len(Packs)]
}

type SolutionV2 struct {
	// an indexed quantity of packages
	// example [0,0,4,5]
	Packs []int
	// Score represnenting order fulfillment
	// example: 0 means precise fulfillment
	// example: negative means overwlown fulfillment
	Score int

	// Total sum of Packs
	Sum int
	// Total quantity of Packs
	Qty int
}

func (s *SolutionV2) Save(n int) {
	s.Score = n - s.Sum
}

// Best solution by Score max [-inf, 0], min Qty
func findBest(s ...SolutionV2) SolutionV2 {
	best := SolutionV2{Score: math.MinInt, Qty: math.MaxInt}

	for _, cnd := range s {
		if cnd.Score > 0 {
			continue
		}

		if cnd.Score > best.Score {
			best = cnd
		} else if cnd.Score == best.Score {
			if cnd.Qty < best.Qty {
				best = cnd
			}
		}
	}

	return best
}
