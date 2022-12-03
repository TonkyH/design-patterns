package book

import "errors"

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int64
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{
		bookShelf: bookShelf,
		index:     0,
	}
}

func (bsi *BookShelfIterator) HasNext() bool {
	return bsi.index < bsi.bookShelf.GetLength()
}

func (bsi *BookShelfIterator) Next() (*Book, error) {
	if !bsi.HasNext() {
		return nil, errors.New("No Element")
	}

	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return book, nil
}
