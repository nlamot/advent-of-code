package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadExerciseInputIntReturnsErrorOnInvalidInput(t *testing.T) {
	input, _ := os.Open("resources/invalid.txt")
	result, err := ReadExerciseInputInt(input)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid input - all rows should contain a valid integer")
}

func TestReadExerciseInputIntReturnsExpectedIntSlice(t *testing.T) {
	input, _ := os.Open("resources/example_int.txt")
	result, err := ReadExerciseInputInt(input)
	assert.Nil(t, err)
	assert.Equal(t, []int{1721,979,366,299,675,1456}, result)
}

func TestReadExerciseInputStringReturnsExpectedStringSlice(t *testing.T) {
	input, _ := os.Open("resources/example_string.txt")
	result, err := ReadExerciseInputString(input)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc",}, result)
}