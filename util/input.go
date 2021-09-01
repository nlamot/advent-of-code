package util

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func ReadExerciseIntegerInput(r io.Reader) ([]int, error) {
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