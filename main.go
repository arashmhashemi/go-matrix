package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r, echo)
	})
	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r, invert)
	})
	http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r, flatten)
	})
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r, sum)
	})
	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r, multiply)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}

func handle(w http.ResponseWriter, r *http.Request, fn func([][]string) string) {
	records, err := readFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	fmt.Fprint(w, fn(records))
}

func readFile(r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	// TODO: properly return error codes and text + documentation
	// TODO: add unit test for the validation logic

	if len(records) == 0 {
		return nil, errors.New("Validation Error: input file must have at least 1 record\n")
	}
	if len(records) != len(records[0]) {
		return nil, errors.New("Validation Error: input file must have equal rows and columns (be square)\n")
	}
	for i, row := range records {
		for j, r := range row {
			v, err := strconv.Atoi(r)
			if err != nil {
				return nil,
					errors.New(fmt.Sprintf("Validation Error: input file must be a valid csv file containing integers with values > 0. Please check (row: %d, col: %d, val: %v).\n", i, j, r))
			}

			if v < 0 {
				return nil,
					errors.New(fmt.Sprintf("Validation Error: input file must be a valid csv file containing integers with values > 0. Please check (row: %d, col: %d, val: %v).\n", i, j, r))
			}
		}
	}

	return records, nil
}
