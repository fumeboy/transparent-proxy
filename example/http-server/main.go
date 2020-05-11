package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Done!")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":20900", nil)
	fmt.Println(err)
}