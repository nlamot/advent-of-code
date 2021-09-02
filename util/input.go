package util

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Can be refactored when Go generics are added, probably in 1.18
func ReadExerciseInputInt(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	result := []int{}
	
	for scanner.Scan() {
		element, e := strconv.Atoi(scanner.Text())
		if e != nil {
			return nil, fmt.Errorf("invalid input - all rows should contain a valid integer")
		}
		result = append(result, element)
	}

	return result, nil
}

func ReadExerciseInputString(r io.Reader) ([]string, error){
	scanner := bufio.NewScanner(r)
	result := []string{}
	
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}