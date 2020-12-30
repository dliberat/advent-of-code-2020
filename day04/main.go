package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type document map[string]string

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func parseDocument(txt string) document {
	doc := make(document, 0)
	txt = strings.ReplaceAll(txt, "\n", " ")
	fields := strings.Split(txt, " ")
	for _, field := range fields {
		split := strings.Split(field, ":")
		key, val := split[0], split[1]
		doc[key] = val
	}
	return doc
}

func parseBatchFile(file string) []document {
	documents := make([]document, 0)
	entries := strings.Split(file, "\n\n")
	for _, entry := range entries {
		doc := parseDocument(entry)
		documents = append(documents, doc)
	}
	return documents
}

func minLenValidator(val string, length int) bool {
	return len(val) >= length
}
func exactLenValidator(val string, length int) bool {
	return len(val) == length
}
func intRangeValidator(val string, min, max int) bool {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return intVal >= min && intVal <= max
}
func byrValidator(val string) bool {
	return exactLenValidator(val, 4) && intRangeValidator(val, 1920, 2002)
}
func iyrValidator(val string) bool {
	return exactLenValidator(val, 4) && intRangeValidator(val, 2010, 2020)
}
func eyrValidator(val string) bool {
	return exactLenValidator(val, 4) && intRangeValidator(val, 2020, 2030)
}
func hgtValidator(val string) bool {
	if !minLenValidator(val, 1) {
		return false
	}
	unit := val[len(val)-2:]
	if unit != "cm" && unit != "in" {
		return false
	}
	val = val[0 : len(val)-2]
	if unit == "cm" {
		return intRangeValidator(val, 150, 193)
	}
	return intRangeValidator(val, 59, 76)
}
func hclValidator(val string) bool {
	re := regexp.MustCompile(`^#[0-9a-z]{6}$`)
	return re.MatchString(val)
}
func eclValidator(val string) bool {
	allowedValues := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, allowed := range allowedValues {
		if val == allowed {
			return true
		}
	}
	return false
}
func pidValidator(val string) bool {
	return exactLenValidator(val, 9) && intRangeValidator(val, 0, 999999999)
}

func isValidPassport(document map[string]string, reqdFields []string) bool {
	for _, field := range reqdFields {
		if document[field] == "" {
			return false
		}
	}
	return true
}

func isValidPassportPart2(doc document) (bool, error) {
	if !byrValidator(doc["byr"]) {
		return false, errors.New("byr")
	}
	if !iyrValidator(doc["iyr"]) {
		return false, errors.New("iyr")
	}
	if !eyrValidator(doc["eyr"]) {
		return false, errors.New("eyr")
	}
	if !hgtValidator(doc["hgt"]) {
		return false, errors.New("hgt")
	}
	if !hclValidator(doc["hcl"]) {
		return false, errors.New("hcl")
	}
	if !eclValidator(doc["ecl"]) {
		return false, errors.New("ecl")
	}
	if !pidValidator(doc["pid"]) {
		return false, errors.New("pid")
	}
	return true, nil
}

func part1(txt string) int {
	documents := parseBatchFile(txt)
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validCount := 0
	for _, doc := range documents {
		if isValidPassport(doc, required) {
			validCount++
		}
	}
	return validCount
}

func part2(txt string) int {
	documents := parseBatchFile(txt)
	validCount := 0
	for _, doc := range documents {
		isValid, _ := isValidPassportPart2(doc)
		if isValid {
			validCount++
		}
	}
	return validCount
}

func main() {
	txt := readInput("input.txt")
	p1 := part1(txt)
	fmt.Println("[PART 1] Total valid passports:", p1)
	p2 := part2(txt)
	fmt.Println("[PART 2] Total valid passports:", p2)
}
