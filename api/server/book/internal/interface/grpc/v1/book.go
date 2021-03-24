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

	in := &input.BookItem{
		Title:        req.GetTitle(),
		Description:  req.GetDescription(),
		Isbn:         req.GetIsbn(),
		ThumbnailURL: req.GetThumbnailUrl(),
		Version:      req.GetVersion(),
		Publisher:    req.GetPublisher(),
		PublishedOn:  req.GetPublishedOn(),
		Authors:      req.GetAuthors(),
		Categories:   req.GetCategories(),
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

	books := make([]*input.BookItem, len(req.GetBooks()))
	for i, item := range req.GetBooks() {
		bookItem := &input.BookItem{
			Title:        item.GetTitle(),
			Description:  item.GetDescription(),
			Isbn:         item.GetIsbn(),
			ThumbnailURL: item.GetThumbnailUrl(),
			Version:      item.GetVersion(),
			Publisher:    item.GetPublisher(),
			PublishedOn:  item.GetPublishedOn(),
			Authors:      item.GetAuthors(),
			Categories:   item.GetCategories(),
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
	as := make([]string, len(b.Authors))
	for i, v := range b.Authors {
		as[i] = v.Name
	}

	cs := make([]string, len(b.Categories))
	for i, v := range b.Categories {
		cs[i] = v.Name
	}

	return &pb.BookResponse{
		Id:           int64(b.ID),
		Title:        b.Title,
		Description:  b.Description,
		Isbn:         b.Isbn,
		ThumbnailUrl: b.ThumbnailURL,
		Version:      b.Version,
		Publisher:    b.Publisher,
		PublishedOn:  datetime.DateToString(b.PublishedOn),
		Authors:      as,
		Categories:   cs,
		CreatedAt:    datetime.TimeToString(b.CreatedAt),
		UpdatedAt:    datetime.TimeToString(b.UpdatedAt),
	}
}

func getBookListResponse(bs []*book.Book) *pb.BookListResponse {
	books := make([]*pb.BookListResponse_Book, len(bs))
	for i, b := range bs {
		as := make([]string, len(b.Authors))
		for i, v := range b.Authors {
			as[i] = v.Name
		}

		cs := make([]string, len(b.Categories))
		for i, v := range b.Categories {
			cs[i] = v.Name
		}

		item := &pb.BookListResponse_Book{
			Id:           int64(b.ID),
			Title:        b.Title,
			Description:  b.Description,
			Isbn:         b.Isbn,
			ThumbnailUrl: b.ThumbnailURL,
			Version:      b.Version,
			Publisher:    b.Publisher,
			PublishedOn:  datetime.DateToString(b.PublishedOn),
			Authors:      as,
			Categories:   cs,
			CreatedAt:    datetime.TimeToString(b.CreatedAt),
			UpdatedAt:    datetime.TimeToString(b.UpdatedAt),
		}

		books[i] = item
	}

	return &pb.BookListResponse{
		Books: books,
	}
}
