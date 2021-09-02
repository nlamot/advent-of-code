package main_test

import (
	"os"
	"testing"

	main "github.com/nlamot/advent-of-code/2020/day-02"
)

func TestDay2Puzzle1(t *testing.T) {
	file, _ := os.Open("resources/puzzle.txt")
	result, _ := main.NumberOfLegalPasswordsInFile(file, false)

	t.Logf("The result for day 2, star 1 is %v", result)
}

func TestDay2Puzzle2(t *testing.T) {
	file, _ := os.Open("resources/puzzle.txt")
	result, _ := main.NumberOfLegalPasswordsInFile(file, true)

	t.Logf("The result for day 2, star 2 is %v", result)
}