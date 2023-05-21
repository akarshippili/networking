package fs

import "os"

type Page struct {
	Title string
	Body  []byte // byte slice
}

func (page *Page) Save() error {
	// save data from page.Body to [page.Title + ".txt"]
	filepath := page.Title + ".txt"
	return os.WriteFile(filepath, page.Body, 0600)
}

func Load(title string) (page *Page, err error) {
	filepath := title + ".txt"
	byteArray, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: byteArray}, nil
}
