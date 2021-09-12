package entity

import pb "github.com/calmato/gran-book/api/gateway/native/proto/service/book"

type Bookshelf struct {
	*pb.Bookshelf
}

type Bookshelves []*Bookshelf

func NewBookshelf(b *pb.Bookshelf) *Bookshelf {
	return &Bookshelf{b}
}

func (b *Bookshelf) Status() BookshelfStatus {
	return NewBookshelfStatus(b.GetStatus())
}

func NewBookshelves(bs []*pb.Bookshelf) Bookshelves {
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
