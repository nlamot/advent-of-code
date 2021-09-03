package main

import (
	"io"

	"github.com/nlamot/advent-of-code/util"
)

func NumberOfTrees(r io.Reader, slopes []Slope)(int, error) {
	input, fileError := util.ReadExerciseInputString(r)
	if fileError != nil {
		return -1, fileError
	}
	result := 1
	for _,s := range slopes {
		numberOfTrees := 0
		x := 0
		y := 0
		width := len(input[0])
		height := len(input)
		for y < height {
			if string(input[y][x]) == "#" {
				numberOfTrees++
			}
			x = (x + s.X) % width
			y+=s.Y
		}
		result*=numberOfTrees
	}
	
	return result, nil
}

type Slope struct {
	X int
	Y int
}