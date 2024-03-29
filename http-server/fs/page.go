package fs

import "os"

type Page struct {
	Title string
	Body  []byte // byte slice
}

var ErrorPage = Page{
	Title: "ERROR",
	Body:  []byte("check if the url entered is right"),
}

func (page *Page) Save() error {
	// save data from page.Body to [page.Title + ".txt"]
	filepath := "./data/" + page.Title + ".txt"
	return os.WriteFile(filepath, page.Body, 0600)
}

func Load(title string) (page *Page, err error) {
	filepath := "./data/" + title + ".txt"
	byteArray, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: byteArray}, nil
}
