package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		fmt.Fprintf(w, "Hey there, what's your name ?\n")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
