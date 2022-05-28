package main

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/kieron-dev/sudoku/solver"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	bs := bytes.Fields(input)
	if len(bs) != 81 {
		panic("need a 9 x 9 grid as input")
	}

	grid := make([][]byte, 9)
	for r := range grid {
		grid[r] = make([]byte, 9)
		for c := range grid[r] {
			val := bs[r*9+c]
			if len(val) > 1 {
				panic("invalid input")
			}
			cell := val[0]
			if cell != '.' && (cell < '0' || cell > '9') {
				panic("input should be 1..9 or '.'")
			}
			grid[r][c] = bs[r*9+c][0]
		}
	}

	solver.New(grid).Solve(0)
}
