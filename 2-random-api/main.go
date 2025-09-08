package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/random", randomNumberHandler)
	http.ListenAndServe(":8080", nil)
}

func randomNumberHandler(w http.ResponseWriter, r *http.Request) {
	n := 1 + rand.Intn(6)
	w.Write([]byte(strconv.Itoa(n)))
	return
}
