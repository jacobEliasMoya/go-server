package main

import (
	"fmt"
	"net/http"
)

func main() {
}

func homeHandler(w http.ResponseWriter, r http.Request) {
	name := r.URL.Query().Get("name")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	_, err := fmt.Fprintf(w, "hello %s!\n", name)

	if err != nil {
		fmt.Println("Response write failed", err)
	}
}
