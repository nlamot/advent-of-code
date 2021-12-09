package main_test

import (
	"os"
	"testing"

	main "github.com/nlamot/advent-of-code/2020/day-04"
	"github.com/stretchr/testify/assert"
)

func TestNumberOfValidPassportsFields(t *testing.T) {
	file, _ := os.Open("resources/example.txt")
	result := main.NumberOfValidPassports(file, true)
	
	assert.Equal(t, 2, result)
}

func TestNumberOfValidPassportsValid(t *testing.T) {
	file, _ := os.Open("resources/valid.txt")
	result := main.NumberOfValidPassports(file, false)
	
	assert.Equal(t, 4, result)
}

func TestNumberOfValidPassportsInvalid(t *testing.T) {
	file, _ := os.Open("resources/invalid.txt")
	result := main.NumberOfValidPassports(file, false)
	
	assert.Equal(t, 0, result)
}