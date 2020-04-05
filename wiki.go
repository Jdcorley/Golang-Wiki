package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
