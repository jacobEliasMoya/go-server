package main

import (
	"fmt"
	"net/http"
)

func main() {
	score := 10
	pointerToScore := &score
	actualScore := *pointerToScore + 4

	fmt.Println(score, ">", pointerToScore, ">", actualScore, ">", score)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello From Go")
}
