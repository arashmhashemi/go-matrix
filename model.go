package main

import (
	"fmt"
	"strings"
)

func echo(records [][]string) string {
	var response string

	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	return response
}

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
