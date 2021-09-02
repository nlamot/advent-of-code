package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/nlamot/advent-of-code/util"
)

func NumberOfLegalPasswordsInFile(r io.Reader, positional bool) (int, error) {
	input, err := util.ReadExerciseInputString(r)
	if err != nil {
		return -1, err
	}
	count := 0
	for _, line := range input {
		pi, e := MapPasswordInput(line)
		if e != nil {
			return -1, e
		}
		if !positional {
			c := strings.Count(pi.Password, pi.Rule.Sequence)
			if (c >= pi.Rule.First && c <= pi.Rule.Second) {
				count++
			}
		} else {
			first := string(pi.Password[pi.Rule.First-1])
			second := string(pi.Password[pi.Rule.Second-1])
			if (first == pi.Rule.Sequence && second != pi.Rule.Sequence) ||
				(first != pi.Rule.Sequence && second == pi.Rule.Sequence) {
					count++
				}
		}

	}
	return count, nil
}

type PasswordInput struct {
	Rule PasswordRule
	Password string
}

type PasswordRule struct {
	First int
	Second int
	Sequence string
}

func MapPasswordInput(element string)(PasswordInput, error) {
	parts := strings.Split(element, ": ")
	if len(parts) != 2 {
		return PasswordInput{}, fmt.Errorf("invalid input - this is not a valid password line - %v", element)
	}
	password := parts[1]
	rule := strings.Split(parts[0], " ")
	if len(rule) != 2 {
		return PasswordInput{}, fmt.Errorf("invalid input - this is not a valid password line - %v", element)
	}
	sequence := rule[1]
	minMax := strings.Split(rule[0], "-")
	if len(minMax) != 2 {
		return PasswordInput{}, fmt.Errorf("invalid input - this is not a valid password line - %v", element)
	}
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])
	return PasswordInput{
		Rule: PasswordRule{
			First: min,
			Second: max,
			Sequence: sequence,
		},
		Password: password,
	}, nil
}