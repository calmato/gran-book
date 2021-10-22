package entity

import "github.com/calmato/gran-book/api/service/proto/book"

type Bookshelf struct {
	*book.Bookshelf
}

type Bookshelves []*Bookshelf

func NewBookshelf(b *book.Bookshelf) *Bookshelf {
	return &Bookshelf{b}
}

func (b *Bookshelf) Status() BookshelfStatus {
	return NewBookshelfStatus(b.GetStatus())
}

func NewBookshelves(bs []*book.Bookshelf) Bookshelves {
	res := make(Bookshelves, len(bs))
	for i := range bs {
		res[i] = NewBookshelf(bs[i])
	}
	return res
}

func (bs Bookshelves) BookIDs() []int64 {
	bookIDs := []int64{}
	books := map[int64]bool{}
	for _, b := range bs {
		if _, ok := books[b.BookId]; ok {
			continue
		}

		books[b.BookId] = true
		bookIDs = append(bookIDs, b.BookId)
	}
	return bookIDs
}
