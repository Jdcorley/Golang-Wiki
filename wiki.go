package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// title and body of wiki page
type Page struct {
	Title string
	Body  []byte
}

// save method for wiki page persistence
// take a pointer p to Page
// takes no parameters and returns a value of type error
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// constructs file anme from the title parameter
// reads the file's contents into a new variable body.
// returns a pointer to a page literal constructed w/
// proper title and body values.
// if second parameter is nil the page successfully loaded.
// otherwise returns an error that can be handled by the caller.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// extracts the page title from r.URL.Path
// The path is re-sliced with [len("/view/":]
// this drops the leading "/view/" component of the request path.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
