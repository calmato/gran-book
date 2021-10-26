package entity

import "github.com/calmato/gran-book/api/proto/book"

type Book struct {
	*book.Book
}

type Books []*Book

func NewBook(b *book.Book) *Book {
	return &Book{b}
}

func (b *Book) AuthorNames() []string {
	names := make([]string, len(b.Authors))
	for i, a := range b.Authors {
		names[i] = a.Name
	}
	return names
}

func (b *Book) AuthorNameKanas() []string {
	names := make([]string, len(b.Authors))
	for i, a := range b.Authors {
		names[i] = a.NameKana
	}
	return names
}

func NewBooks(bs []*book.Book) Books {
	res := make(Books, len(bs))
	for i := range bs {
		res[i] = NewBook(bs[i])
	}
	return res
}

func (bs Books) Map() map[int64]*Book {
	res := make(map[int64]*Book, len(bs))
	for _, b := range bs {
		res[b.Id] = b
	}
	return res
}
