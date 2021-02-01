package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Echos an input matrix without any changes
func echo(records [][]string) string {
	var response string

	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	return response
}

// Returns the matrix as a string in matrix format where the columns and rows are inverted
// example:
// 1,2,3		1,4,7
// 4,5,6   ==>	2,5,8
// 7,8,9		3,6,9
func invert(records [][]string) string {
	var response string

	for i := 0; i < len(records)-1; i++ {
		for j := i + 1; j < len(records); j++ {
			// * swap
			records[i][j], records[j][i] = records[j][i], records[i][j]
		}
		response = fmt.Sprintf("%s%s\n", response, strings.Join(records[i], ","))
	}
	response = fmt.Sprintf("%s%s\n", response, strings.Join(records[len(records)-1], ","))

	return response
}

// Returns the matrix as a 1 line string, with values separated by commas
// example:
// 1,2,3
// 4,5,6   ==>	1,2,3,4,5,6,7,8,9
// 7,8,9
func flatten(records [][]string) string {
	var response string

	for i, row := range records {
		if i == 0 {
			response = strings.Join(row, ",")
		} else {
			response = fmt.Sprintf("%s,%s", response, strings.Join(row, ","))
		}
	}

	return response + "\n"
}

// Returns the sum of the integers in the matrix
// 1,2,3
// 4,5,6   ==>	45
// 7,8,9
func sum(records [][]string) string {
	var result int

	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[i]); j++ {
			x, err := strconv.Atoi(records[i][j])
			if err == nil {
				result += x
			} else {
				fmt.Println(err.Error())
				return "Error!"
			}
		}
	}

	return strconv.Itoa(result) + "\n"
}

// Return the product of the integers in the matrix. If overflow happens, returns 'Overflow!'
// 1,2,3
// 4,5,6   ==>	362880
// 7,8,9
func multiply(records [][]string) string {
	var result int64 = 1
	var ok bool

	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[i]); j++ {
			x, err := strconv.ParseInt(records[i][j], 10, 64)
			if err == nil {
				result, ok = mul64(x, result)
				if !ok {
					return "Overflow!\n"
				}
			} else {
				fmt.Println(err.Error())
				return "Error!\n"
			}
		}
	}

	return strconv.FormatInt(result, 10) + "\n"
}

// multiplies two numbers and checks for overflow.
func mul64(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	c := a * b
	if (c < 0) == ((a < 0) != (b < 0)) {
		if c/b == a {
			return c, true
		}
	}
	return c, false
}
