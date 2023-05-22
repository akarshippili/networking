package router

import (
	"fmt"
	"log"
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

func EditHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fileName := path[len("/edit/"):]
	fmt.Printf("User trying access %s\n", fileName)

	page, err := fs.Load(fileName)
	if err != nil {
		page = &fs.Page{
			Title: fileName,
		}
	}

	fmt.Fprintf(w, "<Title>%s.txt</Title>"+
		"<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		page.Title, page.Title, page.Title, page.Body)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")

	page := fs.Page{
		Title: fileName,
		Body:  []byte(body),
	}
	if err := page.Save(); err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "Bad Request")
	}

	fmt.Fprintf(w, "201 Created")
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey there, what's your name ?\n")
}
