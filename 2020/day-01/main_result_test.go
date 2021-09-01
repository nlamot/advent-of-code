package main_test

import (
	"os"
	"testing"

	"github.com/nlamot/advent-of-code/util"
	main "github.com/nlamot/advent-of-code/2020/day-01"
)

func TestDay1Puzzle1(t *testing.T) {
	file, _ := os.Open("resources/puzzle.txt")
	input, _ := util.ReadExerciseIntegerInput(file)
	result, _ := main.CalculateDuoMultiplicationOnSum(input, 2020)

	t.Logf("The result for day 1, star 1 is %v", result)
}

func TestDay1Puzzle2(t *testing.T) {
	file, _ := os.Open("resources/puzzle.txt")
	input, _ := util.ReadExerciseIntegerInput(file)
	result, _ := main.CalculateTrioMultiplicationOnSum(input, 2020)

	t.Logf("The result for day 1, star 2 is %v", result)
}