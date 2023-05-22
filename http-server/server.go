package main

import (
	"log"
	"net/http"

	"github.com/akarshippili/networking/http-server/router"
)

func main() {
	http.HandleFunc("/view/", router.ViewHandler)
	http.HandleFunc("/", router.RootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
