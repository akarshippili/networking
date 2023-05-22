package router

import (
	"fmt"
	"net/http"

	"github.com/akarshippili/networking/http-server/fs"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fileName := path[len("/view/"):]
	fmt.Printf("User trying access %s\n", fileName)

	page, err := fs.Load(fileName)
	if err != nil {
		page = &fs.ErrorPage
	}

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey there, what's your name ?\n")
}
