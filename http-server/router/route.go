package router

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/akarshippili/networking/http-server/fs"
)

func renderTemplate(w http.ResponseWriter, tmpl string, page *fs.Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fileName := path[len("/view/"):]
	fmt.Printf("User trying access %s\n", fileName)

	page, err := fs.Load(fileName)
	if err != nil {
		http.Redirect(w, r, "/edit/"+fileName, http.StatusFound)
		return
	}

	renderTemplate(w, "view", page)
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

	renderTemplate(w, "edit", page)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")

	page := fs.Page{
		Title: fileName,
		Body:  []byte(body),
	}

	if err := page.Save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/view/"+fileName, http.StatusFound)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey there, what's your name ?\n")
}
