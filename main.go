package main

import (
	"fmt"
	"net/http"
)

func main() {
	mainScore := 0

	fmt.Println(mainScore)

	newScore := incrementOriginal(&mainScore, 55)

	fmt.Println(newScore, ">", mainScore)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello From Go")
}

func incrementOriginal(score *int, incrimentor int) int {
	*score = *score + incrimentor
	return *score
}
