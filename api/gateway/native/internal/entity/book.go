package entity

import (
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
)

type Book struct {
	*pb.Book
}

type Books []*Book

func NewBook(b *pb.Book) *Book {
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

func NewBooks(bs []*pb.Book) Books {
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

type Review struct {
	*pb.Review
}

type Reviews []*Review

func NewReview(r *pb.Review) *Review {
	return &Review{r}
}

func NewReviews(rs []*pb.Review) Reviews {
	res := make(Reviews, len(rs))
	for i := range rs {
		res[i] = NewReview(rs[i])
	}
	return res
}

func (rs Reviews) UserIDs() []string {
	res := make([]string, len(rs))
	for i := range rs {
		res[i] = rs[i].UserId
	}
	return res
}

func (rs Reviews) BookIDs() []int64 {
	res := make([]int64, len(rs))
	for i := range rs {
		res[i] = rs[i].BookId
	}
	return res
}
