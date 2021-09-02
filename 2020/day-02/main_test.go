package main_test

import (
	"os"
	"testing"

	main "github.com/nlamot/advent-of-code/2020/day-02"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfLegalPasswordsInFileWithExampleNonPositional(t *testing.T) {
	file, _ := os.Open("resources/example.txt")
	result, err := main.NumberOfLegalPasswordsInFile(file, false)
	
	assert.Nil(t, err)
	assert.Equal(t, 2, result)
}

func TestNumberOfLegalPasswordsInFileWithExamplePositional(t *testing.T) {
	file, _ := os.Open("resources/example.txt")
	result, err := main.NumberOfLegalPasswordsInFile(file, true)
	
	assert.Nil(t, err)
	assert.Equal(t, 1, result)
}
