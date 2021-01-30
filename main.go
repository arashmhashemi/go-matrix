package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
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

	http.ListenAndServe(":8080", nil)
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

	return records, nil
}
