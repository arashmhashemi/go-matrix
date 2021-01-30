package main

import (
	"testing"
)

var data = []struct {
	input [][]string
}{
	{[][]string{ // 1 row 1 column
		{"1"},
	}},
	{[][]string{ // 3 rows 3 columns
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}},
	{[][]string{ // 4 rows 4 columns
		{"11", "12", "13", "14"},
		{"21", "22", "23", "24"},
		{"31", "32", "33", "34"},
		{"41", "42", "43", "44"},
	}},
}

func getMatrix(index int) [][]string {
	matrix := data[index].input

	duplicate := make([][]string, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]string, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}

	return duplicate
}
func TestEcho(t *testing.T) {
	var testData = []struct {
		input [][]string
		want  string
	}{
		{getMatrix(0), "1\n"},
		{getMatrix(1), "1,2,3\n4,5,6\n7,8,9\n"},
		{getMatrix(2), "11,12,13,14\n21,22,23,24\n31,32,33,34\n41,42,43,44\n"},
	}

	for _, testItem := range testData {
		got := echo(testItem.input)

		if got != testItem.want {
			t.Errorf("got %q, want %q", got, testItem.want)
		}
	}
}

func TestInvert(t *testing.T) {
	var testData = []struct {
		input [][]string
		want  string
	}{
		{getMatrix(0), "1\n"},
		{getMatrix(1), "1,4,7\n2,5,8\n3,6,9\n"},
		{getMatrix(2), "11,21,31,41\n12,22,32,42\n13,23,33,43\n14,24,34,44\n"},
	}

	for _, testItem := range testData {
		got := invert(testItem.input)

		if got != testItem.want {
			t.Errorf("got %q, want %q", got, testItem.want)
		}
	}
}

func TestFlatten(t *testing.T) {
	var testData = []struct {
		input [][]string
		want  string
	}{
		{getMatrix(0), "1\n"},
		{getMatrix(1), "1,2,3,4,5,6,7,8,9\n"},
		{getMatrix(2), "11,12,13,14,21,22,23,24,31,32,33,34,41,42,43,44\n"},
	}

	for _, testItem := range testData {
		got := flatten(testItem.input)

		if got != testItem.want {
			t.Errorf("got %q, want %q", got, testItem.want)
		}
	}
}
