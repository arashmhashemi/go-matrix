package main

import (
	"fmt"
	"strconv"
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

func sum(records [][]string) string {
	var result int

	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[i]); j++ {
			x, err := strconv.Atoi(records[i][j])
			if err == nil {
				result += x
			} else {
				fmt.Println(err.Error())
			}
		}
	}

	return strconv.Itoa(result) + "\n"
}
