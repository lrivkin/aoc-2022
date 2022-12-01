package utils

import (
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contents), "\n"), nil
}

func ReadBlock(path string) ([][]string, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	tests := strings.Split(string(contents), "\n\n")
	var broken [][]string
	for _, t := range tests {
		broken = append(broken, strings.Split(t, "\n"))
	}
	return broken, nil
}

func StringSliceToIntSlice(input []string) ([]int, error) {
	var numbers []int
	for _, x := range input {
		num, err := strconv.Atoi(x)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func ParseIntGrid(path string) ([][]int, error) {
	contents, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	var results [][]int
	for _, line := range contents {
		var row []int
		for _, x := range strings.Split(line, "") {
			num, _ := strconv.Atoi(x)
			row = append(row, num)
		}
		results = append(results, row)
	}
	return results, nil
}

func StringListToIntSlice(input string) ([]int, error) {
	var numbers []int
	separated := strings.Split(input, ",")
	for _, x := range separated {
		num, err := strconv.Atoi(strings.TrimSpace(x))
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func GetMapKeys(mymap map[string]struct{}) []string {
	keys := make([]string, len(mymap))
	i := 0
	for k := range mymap {
		keys[i] = k
		i++
	}
	return keys
}
