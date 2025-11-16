package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from my-first-api!")
	})

	fmt.Println("Server läuft auf http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

