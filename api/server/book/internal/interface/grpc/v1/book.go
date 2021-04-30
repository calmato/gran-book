package v1

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/application"
	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/calmato/gran-book/api/server/book/internal/application/output"
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

func (s *BookServer) ListBookshelf(
	ctx context.Context, req *pb.ListBookshelfRequest,
) (*pb.BookshelfListResponse, error) {
	_, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.ListBookshelf{
		UserID: req.GetUserId(),
		Limit:  int(req.GetLimit()),
		Offset: int(req.GetOffset()),
	}

	bss, out, err := s.BookApplication.ListBookshelf(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfListResponse(bss, out)
	return res, nil
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

func (s *BookServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.BookResponse, error) {
	_, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	as := make([]*input.BookAuthor, len(req.GetAuthors()))
	for i, v := range req.GetAuthors() {
		a := &input.BookAuthor{
			Name:     v.GetName(),
			NameKana: v.GetNameKana(),
		}

		as[i] = a
	}

	in := &input.Book{
		Title:          req.GetTitle(),
		TitleKana:      req.GetTitleKana(),
		Description:    req.GetDescription(),
		Isbn:           req.GetIsbn(),
		Publisher:      req.GetPublisher(),
		PublishedOn:    req.GetPublishedOn(),
		ThumbnailURL:   req.GetThumbnailUrl(),
		RakutenURL:     req.GetRakutenUrl(),
		RakutenGenreID: req.GetRakutenGenreId(),
		Authors:        as,
	}

	b, err := s.BookApplication.Create(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookResponse(b)
	return res, nil
}

func (s *BookServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.BookResponse, error) {
	_, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	as := make([]*input.BookAuthor, len(req.GetAuthors()))
	for i, v := range req.GetAuthors() {
		a := &input.BookAuthor{
			Name:     v.GetName(),
			NameKana: v.GetNameKana(),
		}

		as[i] = a
	}

	in := &input.Book{
		Title:          req.GetTitle(),
		TitleKana:      req.GetTitleKana(),
		Description:    req.GetDescription(),
		Isbn:           req.GetIsbn(),
		Publisher:      req.GetPublisher(),
		PublishedOn:    req.GetPublishedOn(),
		ThumbnailURL:   req.GetThumbnailUrl(),
		RakutenURL:     req.GetRakutenUrl(),
		RakutenGenreID: req.GetRakutenGenreId(),
		Authors:        as,
	}

	b, err := s.BookApplication.Update(ctx, in)
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

	bs, err := s.BookApplication.CreateOrUpdateBookshelf(ctx, in)
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

	bs, err := s.BookApplication.CreateOrUpdateBookshelf(ctx, in)
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

	bs, err := s.BookApplication.CreateOrUpdateBookshelf(ctx, in)
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

	bs, err := s.BookApplication.CreateOrUpdateBookshelf(ctx, in)
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

	bs, err := s.BookApplication.CreateOrUpdateBookshelf(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

func (s *BookServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.EmptyBook, error) {
	_, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.BookApplication.Delete(ctx, int(req.GetBookId()))
	if err != nil {
		return nil, errorHandling(err)
	}

	return &pb.EmptyBook{}, nil
}

func getBookResponse(b *book.Book) *pb.BookResponse {
	as := make([]*pb.BookResponse_Author, len(b.Authors))
	for i, v := range b.Authors {
		a := &pb.BookResponse_Author{
			Name:     v.Name,
			NameKana: v.NameKana,
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
		PublishedOn:    b.PublishedOn,
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

func (s *BookServer) DeleteBookshelf(ctx context.Context, req *pb.DeleteBookshelfRequest) (*pb.EmptyBook, error) {
	cuid, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.BookApplication.DeleteBookshelf(ctx, int(req.GetBookId()), cuid)
	if err != nil {
		return nil, errorHandling(err)
	}

	return &pb.EmptyBook{}, nil
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

func getBookshelfListResponse(bss []*book.Bookshelf, out *output.ListQuery) *pb.BookshelfListResponse {
	bookshelves := make([]*pb.BookshelfListResponse_Bookshelf, len(bss))
	for i, bs := range bss {
		authors := make([]*pb.BookshelfListResponse_Author, len(bs.Book.Authors))
		for j, a := range bs.Book.Authors {
			author := &pb.BookshelfListResponse_Author{
				Name:     a.Name,
				NameKana: a.NameKana,
			}

			authors[j] = author
		}

		book := &pb.BookshelfListResponse_Book{
			Id:             int64(bs.Book.ID),
			Title:          bs.Book.Title,
			TitleKana:      bs.Book.TitleKana,
			Description:    bs.Book.Description,
			Isbn:           bs.Book.Isbn,
			Publisher:      bs.Book.Publisher,
			PublishedOn:    bs.Book.PublishedOn,
			ThumbnailUrl:   bs.Book.ThumbnailURL,
			RakutenUrl:     bs.Book.RakutenURL,
			RakutenGenreId: bs.Book.RakutenGenreID,
			CreatedAt:      datetime.TimeToString(bs.CreatedAt),
			UpdatedAt:      datetime.TimeToString(bs.UpdatedAt),
			Authors:        authors,
		}

		bookshelf := &pb.BookshelfListResponse_Bookshelf{
			Id:         int64(bs.ID),
			BookId:     int64(bs.BookID),
			UserId:     bs.UserID,
			Status:     int32(bs.Status),
			Impression: "",
			ReadOn:     datetime.DateToString(bs.ReadOn),
			CreatedAt:  datetime.TimeToString(bs.CreatedAt),
			UpdatedAt:  datetime.TimeToString(bs.UpdatedAt),
			Book:       book,
		}

		bookshelves[i] = bookshelf
	}

	return &pb.BookshelfListResponse{
		Bookshelves: bookshelves,
		Limit:       int64(out.Limit),
		Offset:      int64(out.Offset),
		Total:       int64(out.Total),
	}
}
