package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadExerciseIntegerInputReturnsErrorOnInvalidInput(t *testing.T) {
	input, _ := os.Open("resources/invalid.txt")
	result, err := ReadExerciseIntegerInput(input)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid input - all rows should contain a valid integer")
}

func TestReadExerciseIntegerInputReturnsExpectedIntSlice(t *testing.T) {
	input, _ := os.Open("resources/example.txt")
	result, err := ReadExerciseIntegerInput(input)
	assert.Nil(t, err)
	assert.Equal(t, []int{1721,979,366,299,675,1456}, result)
}