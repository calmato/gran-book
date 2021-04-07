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

func (s *BookServer) ShowBook(ctx context.Context, req *pb.ShowBookRequest) (*pb.BookResponse, error) {
	_, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	b, err := s.BookApplication.Show(ctx, req.GetIsbn())
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookResponse(b)
	return res, nil
}

func (s *BookServer) ReadBookshelf(
	ctx context.Context, req *pb.ReadBookshelfRequest,
) (*pb.BookshelfResponse, error) {
	cuid, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.Bookshelf{
		UserID: cuid,
		BookID: int(req.GetBookId()),
		Status: book.ReadStatus,
		ReadOn: req.GetReadOn(),
	}

	_, bs, err := s.BookApplication.CreateOrUpdateBookshelf(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

func (s *BookServer) ReadingBookshelf(
	ctx context.Context, req *pb.ReadingBookshelfRequest,
) (*pb.BookshelfResponse, error) {
	cuid, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.Bookshelf{
		UserID: cuid,
		BookID: int(req.GetBookId()),
		Status: book.ReadingStatus,
	}

	_, bs, err := s.BookApplication.CreateOrUpdateBookshelf(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

func (s *BookServer) StackBookshelf(
	ctx context.Context, req *pb.StackBookshelfRequest,
) (*pb.BookshelfResponse, error) {
	cuid, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.Bookshelf{
		UserID: cuid,
		BookID: int(req.GetBookId()),
		Status: book.StackStatus,
	}

	_, bs, err := s.BookApplication.CreateOrUpdateBookshelf(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

func (s *BookServer) WantBookshelf(
	ctx context.Context, req *pb.WantBookshelfRequest,
) (*pb.BookshelfResponse, error) {
	cuid, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.Bookshelf{
		UserID: cuid,
		BookID: int(req.GetBookId()),
		Status: book.WantStatus,
	}

	_, bs, err := s.BookApplication.CreateOrUpdateBookshelf(ctx, in) // TODO: refactor
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

func (s *BookServer) ReleaseBookshelf(
	ctx context.Context, req *pb.ReleaseBookshelfRequest,
) (*pb.BookshelfResponse, error) {
	cuid, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.Bookshelf{
		UserID: cuid,
		BookID: int(req.GetBookId()),
		Status: book.ReleaseStatus,
	}

	_, bs, err := s.BookApplication.CreateOrUpdateBookshelf(ctx, in) // TODO: refactor
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

func getBookResponse(b *book.Book) *pb.BookResponse {
	as := make([]*pb.BookResponse_Author, len(b.Authors))
	for i, v := range b.Authors {
		a := &pb.BookResponse_Author{
			Name:     v.Name,
			NameKana: v.NamaKana,
		}

		as[i] = a
	}

	rs := make([]*pb.BookResponse_Review, len(b.Reviews))
	for i, v := range b.Reviews {
		r := &pb.BookResponse_Review{
			Id:         int64(v.ID),
			UserId:     v.UserID,
			Score:      int32(v.Score),
			Impression: v.Impression,
			CreatedAt:  datetime.TimeToString(v.CreatedAt),
			UpdatedAt:  datetime.TimeToString(v.UpdatedAt),
		}

		rs[i] = r
	}

	bs := &pb.BookResponse_Bookshelf{}
	if b.Bookshelf != nil {
		bs.Id = int64(b.Bookshelf.ID)
		bs.Status = int32(b.Bookshelf.Status)
		bs.ReadOn = datetime.DateToString(b.Bookshelf.ReadOn)
		bs.CreatedAt = datetime.TimeToString(b.Bookshelf.CreatedAt)
		bs.UpdatedAt = datetime.TimeToString(b.Bookshelf.UpdatedAt)
	}

	return &pb.BookResponse{
		Id:             int64(b.ID),
		Title:          b.Title,
		TitleKana:      b.TitleKana,
		Description:    b.Description,
		Isbn:           b.Isbn,
		Publisher:      b.Publisher,
		PublishedOn:    datetime.DateToString(b.PublishedOn),
		ThumbnailUrl:   b.ThumbnailURL,
		RakutenUrl:     b.RakutenURL,
		RakutenGenreId: b.RakutenGenreID,
		CreatedAt:      datetime.TimeToString(b.CreatedAt),
		UpdatedAt:      datetime.TimeToString(b.UpdatedAt),
		Authors:        as,
		Reviews:        rs,
		Bookshelf:      bs,
	}
}

func getBookshelfResponse(bs *book.Bookshelf) *pb.BookshelfResponse {
	return &pb.BookshelfResponse{
		Id:        int64(bs.ID),
		BookId:    int64(bs.BookID),
		UserId:    bs.UserID,
		Status:    int32(bs.Status),
		ReadOn:    datetime.DateToString(bs.ReadOn),
		CreatedAt: datetime.TimeToString(bs.CreatedAt),
		UpdatedAt: datetime.TimeToString(bs.UpdatedAt),
	}
}
