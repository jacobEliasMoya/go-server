package main

import (
	"fmt"
	"net/http"
)

func main() {

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello From Go")
}
