package main

import (
	"log"
	"net/http"

	"github.com/akarshippili/networking/http-server/router"
)

func main() {
	http.HandleFunc("/view/", router.ViewHandler)
	http.HandleFunc("/edit/", router.EditHandler)
	http.HandleFunc("/save/", router.SaveHandler)
	http.HandleFunc("/", router.RootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
