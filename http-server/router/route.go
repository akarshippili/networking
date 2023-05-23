package router

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"github.com/akarshippili/networking/http-server/fs"
)

var templates = template.Must(template.ParseFiles("./templates/view.html", "./templates/edit.html"))

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	path := r.URL.Path
	exp := regexp.MustCompile("^/(edit|view|save)/([A-Za-z0-9]+)$")
	if !exp.MatchString(path) {
		return "", errors.New("invalid page title")
	}

	return path[len("/view/"):], nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, page *fs.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	fileName, err := getTitle(w, r)
	fmt.Printf("User trying access %s\n", fileName)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	page, err := fs.Load(fileName)
	if err != nil {
		http.Redirect(w, r, "/edit/"+fileName, http.StatusFound)
		return
	}
	renderTemplate(w, "view", page)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	fileName, err := getTitle(w, r)
	fmt.Printf("User trying access %s\n", fileName)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	page, err := fs.Load(fileName)
	if err != nil {
		page = &fs.Page{
			Title: fileName,
		}
	}

	renderTemplate(w, "edit", page)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	fileName, err := getTitle(w, r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

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
