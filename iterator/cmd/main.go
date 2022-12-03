package main

import (
	"fmt"
	"iterator/pkg/book"
)

func main() {
	bookShelf := book.NewBookShelf(4)
	bookShelf.AppendBook(book.NewBook("Hello World1"))
	bookShelf.AppendBook(book.NewBook("Hello World2"))
	bookShelf.AppendBook(book.NewBook("Hello World3"))
	bookShelf.AppendBook(book.NewBook("Hello World4"))

	it := bookShelf.Iterator()
	for it.HasNext() {
		book, _ := it.Next()
		fmt.Println(book.GetName())
	}
}
