package book

type BookShelf struct {
	books []*Book
	last  int64
}

func NewBookShelf(maxSize int64) *BookShelf {
	books := make([]*Book, 0, maxSize)
	return &BookShelf{
		books: books,
		last:  0,
	}
}

func (bs *BookShelf) GetBookAt(index int64) *Book {
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book *Book) {
	bs.books = append(bs.books, book)
	bs.last++
}

func (bs *BookShelf) GetLength() int64 {
	return bs.last
}

func (bs *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(bs)
}
