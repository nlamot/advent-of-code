package main

import (
	"fmt"
)

func CalculateDuoMultiplicationOnSum(input []int, sum int) (int, error) {
	length := len(input)
	for i1 := 0; i1 < length; i1++ {
		for i2 := i1+1; i2 < length; i2++ {
			if input[i1] + input[i2] == sum {
				return input[i1] * input[i2], nil
			}
		}
	}
	return -1, fmt.Errorf("invalid input - no 2 rows sum up to %v", sum)
}

func CalculateTrioMultiplicationOnSum(input []int, sum int) (int, error) {
	length := len(input)
	for i1 := 0; i1 < length; i1++ {
		for i2 := i1+1; i2 < length; i2++ {
			for i3 := i2+1; i3 < length; i3++ {
				if input[i1] + input[i2] + input[i3] == sum {
					return input[i1] * input[i2] * input[i3], nil
				}
			}	
		}
	}
	return -1, fmt.Errorf("invalid input - no 2 rows sum up to %v", sum)
}