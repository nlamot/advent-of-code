package main_test

import (
	"os"
	"testing"

	main "github.com/nlamot/advent-of-code/2020/day-03"
)

func TestDay3Puzzle1(t *testing.T) {
	file, _ := os.Open("resources/puzzle.txt")
	result, _ := main.NumberOfTrees(file, []main.Slope{{X: 3, Y: 1}})

	t.Logf("The result for day 3, star 1 is %v", result)
}

func TestDay3Puzzle2(t *testing.T) {
	file, _ := os.Open("resources/puzzle.txt")
	result, _ := main.NumberOfTrees(file, []main.Slope{{X: 1, Y: 1},{X: 3, Y: 1},{X: 5, Y: 1},{X: 7, Y: 1},{X: 1, Y: 2}})

	t.Logf("The result for day 3, star 2 is %v", result)
}