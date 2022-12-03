package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, world!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})
	http.ListenAndServe(fmt.Sprintf(":%d", 5050), nil)
}
