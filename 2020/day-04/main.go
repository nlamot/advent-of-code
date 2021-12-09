package main

import (
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/nlamot/advent-of-code/util"
)

func NumberOfValidPassports(r io.Reader, fieldsOnly bool)(int) {
	input, _ := util.ReadExerciseInputString(r)
	passports := extractPassports(input)

	count := 0
	for _, passport := range passports {
		if validPassportFieldsAvailable(passport) && (fieldsOnly || validPassportFieldsRules(passport)) {
			log.Println(passport)
			count++
		}
	}

	return count
}

func extractPassports(s []string)([]map[string]string) {
	var passports []map[string]string
	passport := make(map[string]string)
	for _, line := range s {
		if line == "" {
			passports = append(passports, passport)
			passport = make(map[string]string)
		} else {
			for _, element := range strings.Split(line, " ") {
				e := strings.Split(element, ":")
				passport[e[0]] = e[1]
			}
		}
	}
	passports = append(passports, passport)
	return passports
}

func validPassportFieldsAvailable(passport map[string]string) bool {
	_, byr := passport["byr"]
	_, iyr := passport["iyr"]
	_, eyr := passport["eyr"]
	_, hgt := passport["hgt"]
	_, hcl := passport["hcl"]
	_, ecl := passport["ecl"]
	_, pid := passport["pid"]
	return byr && iyr && eyr && hgt && hcl && ecl && pid
}

func validPassportFieldsRules(passport map[string]string) bool {
	byr,_ := strconv.Atoi(passport["byr"])
	iyr,_ := strconv.Atoi(passport["iyr"])
	eyr,_ := strconv.Atoi(passport["eyr"])
	hgt,_ := passport["hgt"]
	hcl,_ := passport["hcl"]
	ecl,_ := passport["ecl"]
	pid,_ := passport["pid"]

    validHgtString,_ := regexp.MatchString(`^[1-9][0-9]*(cm|in)$`, hgt)
	if !validHgtString {
		return false
	}
	hgtValue, _ := strconv.Atoi(hgt[:len(hgt)-2])
	hgtUnit := hgt[len(hgt)-2:]
	
	validHcl,_ := regexp.MatchString(`^#[0-9a-f]{6}$`, hcl)
	validEcl,_ := regexp.MatchString(`^amb|blu|brn|gry|grn|hzl|oth$`, ecl)
	validPid,_ := regexp.MatchString(`^[0-9]{9}$`, pid)

	return (byr >= 1920 && byr <= 2002) &&
		   (iyr >= 2010 && iyr <= 2010) &&
		   (eyr >= 2020 && eyr <= 2030) &&
		   validHgtString && ((hgtUnit == "cm" && (hgtValue >= 150 || hgtValue <= 193)) || (hgtUnit == "in" && (hgtValue >= 59 || hgtValue <= 76))) &&
		   validHcl &&
		   validEcl &&
		   validPid
}