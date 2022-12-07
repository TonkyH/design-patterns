package factory

import "os"

type IPage interface {
	MakeHTML() string
}

type Page struct {
	IPage
	Title   string
	Author  string
	Content []Item
}

func NewPage(title string, author string) *Page {
	return &Page{
		Title:  title,
		Author: author,
	}
}

func (p *Page) Add(item Item) {
	p.Content = append(p.Content, item)
}

func (p *Page) Out(filename string) {
	f, _ := os.Create(filename)
	defer f.Close()

	f.WriteString(p.MakeHTML())
}
