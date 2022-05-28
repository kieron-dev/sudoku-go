package solver

import (
	"fmt"
	"strconv"
)

type Possibilities []bool

func NewPossibilities() Possibilities {
	return make(Possibilities, 9*9*9)
}

func (p Possibilities) getIndex(r, c, val int) int {
	return 9*(9*r+c) + val - 1
}

func (p Possibilities) isPossible(r, c, val int) bool {
	return !p[p.getIndex(r, c, val)]
}

func (p Possibilities) setNotPossible(r, c, val int) {
	p[p.getIndex(r, c, val)] = true
}

func (p Possibilities) setVal(r, c, val int) {
	for i := 1; i < 10; i++ {
		if i != val {
			p.setNotPossible(r, c, i)
		}
	}
}

func (p Possibilities) Print() {
	for r := 0; r < 9; r++ {
		for i := 0; i < 3; i++ {
			for c := 0; c < 9; c++ {
				for j := 0; j < 3; j++ {
					val := 3*i + j + 1
					valS := strconv.Itoa(val)
					if !p.isPossible(r, c, val) {
						valS = "."
					}
					fmt.Printf("%s", valS)
				}
				fmt.Print("   ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func (p Possibilities) possibilities(r, c int) []int {
	res := []int{}
	for i := 1; i <= 9; i++ {
		if p.isPossible(r, c, i) {
			res = append(res, i)
		}
	}
	return res
}

func (p Possibilities) clone() Possibilities {
	n := NewPossibilities()
	copy(n, p)

	return n
}
