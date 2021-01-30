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

func TestEcho(t *testing.T) {
	var testData = []struct {
		input [][]string
		want  string
	}{
		{data[0].input, "1\n"},
		{data[1].input, "1,2,3\n4,5,6\n7,8,9\n"},
		{data[2].input, "11,12,13,14\n21,22,23,24\n31,32,33,34\n41,42,43,44\n"},
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
		{data[0].input, "1\n"},
		{data[1].input, "1,4,7\n2,5,8\n3,6,9\n"},
		//{data[2].input, "11,12,13,14\n21,22,23,24\n31,32,33,34\n41,42,43,44\n"},
	}

	for _, testItem := range testData {
		got := invert(testItem.input)

		if got != testItem.want {
			t.Errorf("got %q, want %q", got, testItem.want)
		}
	}
}
