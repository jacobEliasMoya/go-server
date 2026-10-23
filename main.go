package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	var message strings.Builder
	fmt.Fprintln(&message, "blah blah blah")

	fmt.Println(message)

	http.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello From Go")
}
