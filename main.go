package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go Web Server!")
	})

	fmt.Println("Server running on http://localhost:8080")

	fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	server := http.ListenAndServe(":8080", nil)
	if server != nil {
		fmt.Println("Error:", server)
	}
}