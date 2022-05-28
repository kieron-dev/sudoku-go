package solver

import (
	"fmt"
	"strings"
)

type Solver struct {
	possibilities Possibilities
}

func New(grid [][]byte) *Solver {
	s := &Solver{possibilities: NewPossibilities()}
	for r := range grid {
		for c, b := range grid[r] {
			if isNum(b) {
				s.setVal(r, c, int(b-'0'))
			}
		}
	}
	s.possibilities.Print()
	fmt.Println()

	return s
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func (s *Solver) Print() {
	fmt.Println()
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			pos := s.possibilities.possibilities(r, c)
			sep := " "
			if c == 2 || c == 5 {
				sep = "|"
			}
			fmt.Printf(" %d %s", pos[0], sep)
		}
		fmt.Println()
		if r == 2 || r == 5 {
			fmt.Println("-----------+-----------+-----------")
		} else if r != 8 {
			fmt.Println("           |           |           ")
		}
	}
}

func (s *Solver) setVal(row, col, val int) bool {
	queue := [][3]int{{row, col, val}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		r := cur[0]
		c := cur[1]
		v := cur[2]

		for i := 0; i < 9; i++ {
			if i == c {
				continue
			}
			lenBefore := len(s.possibilities.possibilities(r, i))
			s.possibilities.setNotPossible(r, i, v)
			lenAfter := len(s.possibilities.possibilities(r, i))
			if lenBefore > 1 && lenAfter == 1 {
				queue = append(queue, [3]int{r, i, s.possibilities.possibilities(r, i)[0]})
			}
			if lenAfter == 0 {
				return false
			}
		}

		for i := 0; i < 9; i++ {
			if i == r {
				continue
			}
			lenBefore := len(s.possibilities.possibilities(i, c))
			s.possibilities.setNotPossible(i, c, v)
			lenAfter := len(s.possibilities.possibilities(i, c))
			if lenBefore > 1 && lenAfter == 1 {
				queue = append(queue, [3]int{i, c, s.possibilities.possibilities(i, c)[0]})
			}
			if lenAfter == 0 {
				return false
			}
		}

		blockR := 3 * (r / 3)
		blockC := 3 * (c / 3)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				rc := blockR + i
				cc := blockC + j
				if rc == r && cc == c {
					continue
				}
				lenBefore := len(s.possibilities.possibilities(rc, cc))
				s.possibilities.setNotPossible(rc, cc, v)
				lenAfter := len(s.possibilities.possibilities(rc, cc))
				if lenBefore > 1 && lenAfter == 1 {
					queue = append(queue, [3]int{rc, cc, s.possibilities.possibilities(rc, cc)[0]})
				}
				if lenAfter == 0 {
					return false
				}
			}
		}

		s.possibilities.setVal(r, c, v)
	}

	return true
}

func (s *Solver) isComplete() bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if len(s.possibilities.possibilities(r, c)) > 1 {
				return false
			}
		}
	}

	return true
}

func (s *Solver) getMin() (int, int) {
	for n := 2; n < 10; n++ {
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				if len(s.possibilities.possibilities(r, c)) == n {
					return r, c
				}
			}
		}
	}
	panic("getMin when none exists")
}

func (s *Solver) Solve(depth int) bool {
	if s.isComplete() {
		s.Print()
		return true
	}

	r, c := s.getMin()
	for _, v := range s.possibilities.possibilities(r, c) {
		fmt.Printf("%strying %d,%d = %d\n", strings.Repeat("  ", depth), r, c, v)
		prevPoss := s.possibilities.clone()
		if s.setVal(r, c, v) {
			ok := s.Solve(depth + 1)
			if ok {
				return true
			}
		}
		s.possibilities = prevPoss
	}

	return false
}
