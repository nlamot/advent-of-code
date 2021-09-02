package main_test

import (
	"os"
	"testing"

	"github.com/nlamot/advent-of-code/util"
	main "github.com/nlamot/advent-of-code/2020/day-01"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMultiplicationOnSumWithExample(t *testing.T) {
	file, _ := os.Open("resources/example.txt")
	input, _ := util.ReadExerciseInputInt(file)
	resultDuo, errDuo := main.CalculateDuoMultiplicationOnSum(input, 2020)
	resultTrio, errTrio := main.CalculateTrioMultiplicationOnSum(input, 2020)
	
	assert.Nil(t, errDuo)
	assert.Equal(t, 514579, resultDuo)
	assert.Nil(t, errTrio)
	assert.Equal(t, 241861950, resultTrio)
}

func TestCalculateMultiplicationOnSumReturnsErrorWhenSumDoesntExist (t *testing.T) {
	file, _ := os.Open("resources/example.txt")
	input, _ := util.ReadExerciseInputInt(file)
	resultDuo, errDuo := main.CalculateDuoMultiplicationOnSum(input, 2021)
	resultTrio, errTrio := main.CalculateTrioMultiplicationOnSum(input, 2021)
	
	assert.EqualError(t, errDuo, "invalid input - no 2 rows sum up to 2021")
	assert.Equal(t, -1, resultDuo)
	assert.EqualError(t, errTrio, "invalid input - no 2 rows sum up to 2021")
	assert.Equal(t, -1, resultTrio)
}