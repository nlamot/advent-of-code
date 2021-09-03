package main_test

import (
	"os"
	"testing"

	main "github.com/nlamot/advent-of-code/2020/day-03"
	"github.com/stretchr/testify/assert"
)


func TestNumberOfTreesWithExample(t *testing.T) {
	file, _ := os.Open("resources/example.txt")
	result, err := main.NumberOfTrees(file, []main.Slope{{X: 3, Y: 1}})
	
	assert.Nil(t, err)
	assert.Equal(t, 7, result)
}

func TestNumberOfTreesWithExampleMultiple(t *testing.T) {
	file, _ := os.Open("resources/example.txt")
	result, err := main.NumberOfTrees(file, []main.Slope{{X: 1, Y: 1},{X: 3, Y: 1},{X: 5, Y: 1},{X: 7, Y: 1},{X: 1, Y: 2}})
	
	assert.Nil(t, err)
	assert.Equal(t, 336, result)
}