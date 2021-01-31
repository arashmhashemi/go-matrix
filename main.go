package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
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

	return records, nil
}
