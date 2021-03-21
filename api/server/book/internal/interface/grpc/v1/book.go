package v1

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/application"
	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/lib/datetime"
	pb "github.com/calmato/gran-book/api/server/book/proto"
)

// BookServer - Bookインターフェースの構造体
type BookServer struct {
	pb.UnimplementedBookServiceServer
	AuthApplication application.AuthApplication
	BookApplication application.BookApplication
}

// CreateBook - 書籍登録
func (s *BookServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.BookResponse, error) {
	_, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	as := make([]*input.BookAuthor, len(req.Authors))
	for i, v := range req.GetAuthors() {
		a := &input.BookAuthor{
			Name: v.GetName(),
		}

		as[i] = a
	}

	cs := make([]*input.BookCategory, len(req.Categories))
	for i, v := range req.GetCategories() {
		c := &input.BookCategory{
			Name: v.GetName(),
		}

		cs[i] = c
	}

	in := &input.BookItem{
		Title:        req.GetTitle(),
		Description:  req.GetDescription(),
		Isbn:         req.GetIsbn(),
		ThumbnailURL: req.GetThumbnailUrl(),
		Version:      req.GetVersion(),
		Publisher:    req.GetPublisher(),
		PublishedOn:  req.GetPublishedOn(),
		Authors:      as,
		Categories:   cs,
	}

	b, err := s.BookApplication.Create(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookResponse(b)
	return res, nil
}

func (s *BookServer) CreateAndUpdateBooks(
	ctx context.Context, req *pb.CreateAndUpdateBooksRequest,
) (*pb.BookListResponse, error) {
	_, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	books := make([]*input.BookItem, len(req.GetItems()))
	for i, item := range req.GetItems() {
		as := make([]*input.BookAuthor, len(item.GetAuthors()))
		for j, v := range item.GetAuthors() {
			a := &input.BookAuthor{
				Name: v.GetName(),
			}

			as[j] = a
		}

		cs := make([]*input.BookCategory, len(item.GetCategories()))
		for j, v := range item.GetCategories() {
			c := &input.BookCategory{
				Name: v.GetName(),
			}

			cs[j] = c
		}

		bookItem := &input.BookItem{
			Title:        item.GetTitle(),
			Description:  item.GetDescription(),
			Isbn:         item.GetIsbn(),
			ThumbnailURL: item.GetThumbnailUrl(),
			Version:      item.GetVersion(),
			Publisher:    item.GetPublisher(),
			PublishedOn:  item.GetPublishedOn(),
			Authors:      as,
			Categories:   cs,
		}

		books[i] = bookItem
	}

	in := &input.CreateAndUpdateBooks{
		Books: books,
	}

	bs, err := s.BookApplication.MultipleCreateAndUpdate(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookListResponse(bs)
	return res, nil
}

func getBookResponse(b *book.Book) *pb.BookResponse {
	p := &pb.BookResponse_Publisher{}
	if b.Publisher != nil {
		p.Id = int64(b.Publisher.ID)
		p.Name = b.Publisher.Name
	}

	as := make([]*pb.BookResponse_Author, len(b.Authors))
	for i, v := range b.Authors {
		a := &pb.BookResponse_Author{
			Id:   int64(v.ID),
			Name: v.Name,
		}

		as[i] = a
	}

	cs := make([]*pb.BookResponse_Category, len(b.Categories))
	for i, v := range b.Categories {
		c := &pb.BookResponse_Category{
			Id:   int64(v.ID),
			Name: v.Name,
		}

		cs[i] = c
	}

	return &pb.BookResponse{
		Id:           int64(b.ID),
		Title:        b.Title,
		Description:  b.Description,
		Isbn:         b.Isbn,
		ThumbnailUrl: b.ThumbnailURL,
		Version:      b.Version,
		PublishedOn:  datetime.DateToString(b.PublishedOn),
		Publisher:    p,
		Authors:      as,
		Categories:   cs,
		CreatedAt:    datetime.TimeToString(b.CreatedAt),
		UpdatedAt:    datetime.TimeToString(b.UpdatedAt),
	}
}

func getBookListResponse(bs []*book.Book) *pb.BookListResponse {
	items := make([]*pb.BookListResponse_Item, len(bs))
	for i, b := range bs {
		p := &pb.BookListResponse_Publisher{}
		if b.Publisher != nil {
			p.Id = int64(b.Publisher.ID)
			p.Name = b.Publisher.Name
		}

		as := make([]*pb.BookListResponse_Author, len(b.Authors))
		for i, v := range b.Authors {
			a := &pb.BookListResponse_Author{
				Id:   int64(v.ID),
				Name: v.Name,
			}

			as[i] = a
		}

		cs := make([]*pb.BookListResponse_Category, len(b.Categories))
		for i, v := range b.Categories {
			c := &pb.BookListResponse_Category{
				Id:   int64(v.ID),
				Name: v.Name,
			}

			cs[i] = c
		}

		item := &pb.BookListResponse_Item{
			Id:           int64(b.ID),
			Title:        b.Title,
			Description:  b.Description,
			Isbn:         b.Isbn,
			ThumbnailUrl: b.ThumbnailURL,
			Version:      b.Version,
			PublishedOn:  datetime.DateToString(b.PublishedOn),
			Publisher:    p,
			Authors:      as,
			Categories:   cs,
			CreatedAt:    datetime.TimeToString(b.CreatedAt),
			UpdatedAt:    datetime.TimeToString(b.UpdatedAt),
		}

		items[i] = item
	}

	return &pb.BookListResponse{
		Items: items,
	}
}
