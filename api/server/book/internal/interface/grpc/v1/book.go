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

	as := make([]*input.CreateBookAuthor, len(req.Authors))
	for i, v := range req.Authors {
		a := &input.CreateBookAuthor{
			Name: v.Name,
		}

		as[i] = a
	}

	cs := make([]*input.CreateBookCategory, len(req.Categories))
	for i, v := range req.Categories {
		c := &input.CreateBookCategory{
			Name: v.Name,
		}

		cs[i] = c
	}

	in := &input.CreateBook{
		Title:        req.Title,
		Description:  req.Description,
		Isbn:         req.Isbn,
		ThumbnailURL: req.ThumbnailUrl,
		Version:      req.Version,
		Publisher:    req.Publisher,
		PublishedOn:  req.PublishedOn,
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

func getBookResponse(b *book.Book) *pb.BookResponse {
	as := make([]*pb.BookResponse_Author, len(b.Authors))
	for i, v := range b.Authors {
		a := &pb.BookResponse_Author{
			Id:   v.ID,
			Name: v.Name,
		}

		as[i] = a
	}

	cs := make([]*pb.BookResponse_Category, len(b.Categories))
	for i, v := range b.Categories {
		c := &pb.BookResponse_Category{
			Id:   v.ID,
			Name: v.Name,
		}

		cs[i] = c
	}

	return &pb.BookResponse{
		Id:           b.ID,
		PublisherId:  b.PublisherID,
		Title:        b.Title,
		Description:  b.Description,
		Isbn:         b.Isbn,
		ThumbnailUrl: b.ThumbnailURL,
		Version:      b.Version,
		PublishedOn:  datetime.DateToString(b.PublishedOn),
		CreatedAt:    datetime.TimeToString(b.CreatedAt),
		UpdatedAt:    datetime.TimeToString(b.UpdatedAt),
		Authors:      as,
		Categories:   cs,
	}
}
