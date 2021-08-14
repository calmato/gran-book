package server

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/application"
	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/internal/interface/validation"
	"github.com/calmato/gran-book/api/server/book/pkg/database"
	"github.com/calmato/gran-book/api/server/book/pkg/datetime"
	pb "github.com/calmato/gran-book/api/server/book/proto"
)

type bookServer struct {
	pb.UnimplementedBookServiceServer
	bookRequestValidation validation.BookRequestValidation
	bookApplication       application.BookApplication
}

func NewBookServer(brv validation.BookRequestValidation, ba application.BookApplication) pb.BookServiceServer {
	return &bookServer{
		bookRequestValidation: brv,
		bookApplication:       ba,
	}
}

// ListBookshelf - 本棚の書籍一覧取得
func (s *bookServer) ListBookshelf(
	ctx context.Context, req *pb.ListBookshelfRequest,
) (*pb.BookshelfListResponse, error) {
	err := s.bookRequestValidation.ListBookshelf(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	q := &database.ListQuery{
		Limit:  limit,
		Offset: offset,
		Conditions: []*database.ConditionQuery{
			{
				Field:    "user_id",
				Operator: "==",
				Value:    req.GetUserId(),
			},
		},
	}

	if req.GetOrder() != nil {
		o := &database.OrderQuery{
			Field:   req.GetOrder().GetField(),
			OrderBy: int(req.GetOrder().GetOrderBy()),
		}

		q.Order = o
	}

	bss, total, err := s.bookApplication.ListBookshelf(ctx, q)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfListResponse(bss, limit, offset, total)
	return res, nil
}

// ListBookReview - 書籍のレビュー一覧取得
func (s *bookServer) ListBookReview(
	ctx context.Context, req *pb.ListBookReviewRequest,
) (*pb.ReviewListResponse, error) {
	err := s.bookRequestValidation.ListBookReview(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	rs, total, err := s.bookApplication.ListBookReview(ctx, int(req.GetBookId()), limit, offset)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getReviewListResponse(rs, limit, offset, total)
	return res, nil
}

// ListUserReview - ユーザーのレビュー一覧取得
func (s *bookServer) ListUserReview(
	ctx context.Context, req *pb.ListUserReviewRequest,
) (*pb.ReviewListResponse, error) {
	err := s.bookRequestValidation.ListUserReview(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	rs, total, err := s.bookApplication.ListUserReview(ctx, req.GetUserId(), limit, offset)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getReviewListResponse(rs, limit, offset, total)
	return res, nil
}

// MultiGetBooks - 書籍一覧取得 (ID指定)
func (s *bookServer) MultiGetBooks(ctx context.Context, req *pb.MultiGetBooksRequest) (*pb.BookMapResponse, error) {
	err := s.bookRequestValidation.MultiGetBooks(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	bookIDs := make([]int, len(req.GetBookIds()))
	for i, bookID := range req.BookIds {
		bookIDs[i] = int(bookID)
	}
	bs, err := s.bookApplication.MultiGet(ctx, bookIDs)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookMapResponse(bs)
	return res, nil
}

// GetBook - 書籍取得
func (s *bookServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.BookResponse, error) {
	err := s.bookRequestValidation.GetBook(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	b, err := s.bookApplication.Get(ctx, int(req.GetBookId()))
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookResponse(b)
	return res, nil
}

// GetBook - 書籍取得 (ISBN指定)
func (s *bookServer) GetBookByIsbn(ctx context.Context, req *pb.GetBookByIsbnRequest) (*pb.BookResponse, error) {
	err := s.bookRequestValidation.GetBookByIsbn(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	b, err := s.bookApplication.GetByIsbn(ctx, req.GetIsbn())
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookResponse(b)
	return res, nil
}

// GetBookshlf - 本棚の書籍取得
func (s *bookServer) GetBookshelf(ctx context.Context, req *pb.GetBookshelfRequest) (*pb.BookshelfResponse, error) {
	err := s.bookRequestValidation.GetBookshelf(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	bs, err := s.bookApplication.GetBookshelfByUserIDAndBookID(ctx, req.GetUserId(), int(req.GetBookId()))
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

// GetReview - レビュー取得
func (s *bookServer) GetReview(ctx context.Context, req *pb.GetReviewRequest) (*pb.ReviewResponse, error) {
	err := s.bookRequestValidation.GetReview(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	r, err := s.bookApplication.GetReview(ctx, int(req.GetReviewId()))
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getReviewResponse(r)
	return res, nil
}

// CreateBook - 書籍登録
func (s *bookServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.BookResponse, error) {
	err := s.bookRequestValidation.CreateBook(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	as := make([]*book.Author, len(req.GetAuthors()))
	for i, author := range req.GetAuthors() {
		a := &book.Author{
			Name:     author.GetName(),
			NameKana: author.GetNameKana(),
		}

		as[i] = a
	}

	b := &book.Book{
		Title:          req.GetTitle(),
		TitleKana:      req.GetTitleKana(),
		Description:    req.GetDescription(),
		Isbn:           req.GetIsbn(),
		Publisher:      req.GetPublisher(),
		PublishedOn:    req.GetPublishedOn(),
		ThumbnailURL:   req.GetThumbnailUrl(),
		RakutenURL:     req.GetRakutenUrl(),
		RakutenSize:    req.GetRakutenSize(),
		RakutenGenreID: req.GetRakutenGenreId(),
		Authors:        as,
	}
	err = s.bookApplication.Create(ctx, b)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookResponse(b)
	return res, nil
}

// UpdateBook - 書籍更新
func (s *bookServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.BookResponse, error) {
	err := s.bookRequestValidation.UpdateBook(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	b, err := s.bookApplication.GetByIsbn(ctx, req.GetIsbn())
	if err != nil {
		return nil, errorHandling(err)
	}

	as := make([]*book.Author, len(req.GetAuthors()))
	for i, author := range req.GetAuthors() {
		a := &book.Author{
			Name:     author.GetName(),
			NameKana: author.GetNameKana(),
		}

		as[i] = a
	}

	b.Title = req.GetTitle()
	b.TitleKana = req.GetTitleKana()
	b.Description = req.GetDescription()
	b.Isbn = req.GetIsbn()
	b.Publisher = req.GetPublisher()
	b.PublishedOn = req.GetPublishedOn()
	b.ThumbnailURL = req.GetThumbnailUrl()
	b.RakutenURL = req.GetRakutenUrl()
	b.RakutenSize = req.GetRakutenSize()
	b.RakutenGenreID = req.GetRakutenGenreId()
	b.Authors = as

	err = s.bookApplication.Update(ctx, b)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookResponse(b)
	return res, nil
}

// ReadBookshelf - 読んだ本の登録
func (s *bookServer) ReadBookshelf(ctx context.Context, req *pb.ReadBookshelfRequest) (*pb.BookshelfResponse, error) {
	err := s.bookRequestValidation.ReadBookshelf(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	bs, err := s.bookApplication.GetBookshelfByUserIDAndBookIDWithRelated(ctx, req.GetUserId(), int(req.GetBookId()))
	if err != nil {
		return nil, errorHandling(err)
	}

	bs.BookID = int(req.GetBookId())
	bs.UserID = req.GetUserId()
	bs.Status = book.ReadStatus
	bs.ReadOn = datetime.StringToDate(req.GetReadOn())

	bs.ReviewID = bs.Review.ID
	bs.Review.BookID = int(req.GetBookId())
	bs.Review.UserID = req.GetUserId()
	bs.Review.Impression = req.GetImpression()

	err = s.bookApplication.CreateOrUpdateBookshelf(ctx, bs)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

// ReadBookshelf - 読んでいる本の登録
func (s *bookServer) ReadingBookshelf(
	ctx context.Context, req *pb.ReadingBookshelfRequest,
) (*pb.BookshelfResponse, error) {
	err := s.bookRequestValidation.ReadingBookshelf(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	bs, err := s.bookApplication.GetBookshelfByUserIDAndBookIDWithRelated(ctx, req.GetUserId(), int(req.GetBookId()))
	if err != nil {
		return nil, errorHandling(err)
	}

	bs.BookID = int(req.GetBookId())
	bs.UserID = req.GetUserId()
	bs.Status = book.ReadingStatus

	err = s.bookApplication.CreateOrUpdateBookshelf(ctx, bs)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

// StackedBookshelf - 積読本の登録
func (s *bookServer) StackedBookshelf(
	ctx context.Context, req *pb.StackedBookshelfRequest,
) (*pb.BookshelfResponse, error) {
	err := s.bookRequestValidation.StackedBookshelf(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	bs, err := s.bookApplication.GetBookshelfByUserIDAndBookIDWithRelated(ctx, req.GetUserId(), int(req.GetBookId()))
	if err != nil {
		return nil, errorHandling(err)
	}

	bs.BookID = int(req.GetBookId())
	bs.UserID = req.GetUserId()
	bs.Status = book.StackedStatus

	err = s.bookApplication.CreateOrUpdateBookshelf(ctx, bs)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

// WantBookshelf - 読みたい本の登録
func (s *bookServer) WantBookshelf(ctx context.Context, req *pb.WantBookshelfRequest) (*pb.BookshelfResponse, error) {
	err := s.bookRequestValidation.WantBookshelf(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	bs, err := s.bookApplication.GetBookshelfByUserIDAndBookIDWithRelated(ctx, req.GetUserId(), int(req.GetBookId()))
	if err != nil {
		return nil, errorHandling(err)
	}

	bs.BookID = int(req.GetBookId())
	bs.UserID = req.GetUserId()
	bs.Status = book.WantStatus

	err = s.bookApplication.CreateOrUpdateBookshelf(ctx, bs)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

// ReleaseBookshelf - 手放したい本の登録
func (s *bookServer) ReleaseBookshelf(
	ctx context.Context, req *pb.ReleaseBookshelfRequest,
) (*pb.BookshelfResponse, error) {
	err := s.bookRequestValidation.ReleaseBookshelf(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	bs, err := s.bookApplication.GetBookshelfByUserIDAndBookIDWithRelated(ctx, req.GetUserId(), int(req.GetBookId()))
	if err != nil {
		return nil, errorHandling(err)
	}

	bs.BookID = int(req.GetBookId())
	bs.UserID = req.GetUserId()
	bs.Status = book.ReleaseStatus

	err = s.bookApplication.CreateOrUpdateBookshelf(ctx, bs)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getBookshelfResponse(bs)
	return res, nil
}

// DeleteBook - 書籍削除
func (s *bookServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.Empty, error) {
	err := s.bookRequestValidation.DeleteBook(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	b, err := s.bookApplication.Get(ctx, int(req.GetBookId()))
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.bookApplication.Delete(ctx, b)
	if err != nil {
		return nil, errorHandling(err)
	}

	return &pb.Empty{}, nil
}

// DeleteBookshelf - 本棚の書籍削除
func (s *bookServer) DeleteBookshelf(ctx context.Context, req *pb.DeleteBookshelfRequest) (*pb.Empty, error) {
	err := s.bookRequestValidation.DeleteBookshelf(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	bs, err := s.bookApplication.GetBookshelfByUserIDAndBookID(ctx, req.GetUserId(), int(req.GetBookId()))
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.bookApplication.DeleteBookshelf(ctx, bs)
	if err != nil {
		return nil, errorHandling(err)
	}

	return &pb.Empty{}, nil
}

func getBookResponse(b *book.Book) *pb.BookResponse {
	authors := make([]*pb.BookResponse_Author, len(b.Authors))
	for i, a := range b.Authors {
		author := &pb.BookResponse_Author{
			Name:     a.Name,
			NameKana: a.NameKana,
		}

		authors[i] = author
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
		RakutenSize:    b.RakutenSize,
		RakutenGenreId: b.RakutenGenreID,
		CreatedAt:      datetime.TimeToString(b.CreatedAt),
		UpdatedAt:      datetime.TimeToString(b.UpdatedAt),
		Authors:        authors,
	}
}

func getBookMapResponse(bs []*book.Book) *pb.BookMapResponse {
	books := map[int64]*pb.BookMapResponse_Book{}
	for _, b := range bs {
		authors := make([]*pb.BookMapResponse_Author, len(b.Authors))
		for j, a := range b.Authors {
			author := &pb.BookMapResponse_Author{
				Name:     a.Name,
				NameKana: a.NameKana,
			}

			authors[j] = author
		}

		book := &pb.BookMapResponse_Book{
			Id:             int64(b.ID),
			Title:          b.Title,
			TitleKana:      b.TitleKana,
			Description:    b.Description,
			Isbn:           b.Isbn,
			Publisher:      b.Publisher,
			PublishedOn:    b.PublishedOn,
			ThumbnailUrl:   b.ThumbnailURL,
			RakutenUrl:     b.RakutenURL,
			RakutenSize:    b.RakutenSize,
			RakutenGenreId: b.RakutenGenreID,
			CreatedAt:      datetime.TimeToString(b.CreatedAt),
			UpdatedAt:      datetime.TimeToString(b.UpdatedAt),
			Authors:        authors,
		}

		books[book.Id] = book
	}

	return &pb.BookMapResponse{
		Books: books,
	}
}

func getBookshelfResponse(bs *book.Bookshelf) *pb.BookshelfResponse {
	res := &pb.BookshelfResponse{
		Id:        int64(bs.ID),
		BookId:    int64(bs.BookID),
		UserId:    bs.UserID,
		ReviewId:  int64(bs.ReviewID),
		Status:    int32(bs.Status),
		ReadOn:    datetime.DateToString(bs.ReadOn),
		CreatedAt: datetime.TimeToString(bs.CreatedAt),
		UpdatedAt: datetime.TimeToString(bs.UpdatedAt),
	}

	if bs.Book != nil {
		authors := make([]*pb.BookshelfResponse_Author, len(bs.Book.Authors))
		for i, a := range bs.Book.Authors {
			author := &pb.BookshelfResponse_Author{
				Name:     a.Name,
				NameKana: a.NameKana,
			}

			authors[i] = author
		}

		book := &pb.BookshelfResponse_Book{
			Id:             int64(bs.Book.ID),
			Title:          bs.Book.Title,
			TitleKana:      bs.Book.TitleKana,
			Description:    bs.Book.Description,
			Isbn:           bs.Book.Isbn,
			Publisher:      bs.Book.Publisher,
			PublishedOn:    bs.Book.PublishedOn,
			ThumbnailUrl:   bs.Book.ThumbnailURL,
			RakutenUrl:     bs.Book.RakutenURL,
			RakutenSize:    bs.Book.RakutenSize,
			RakutenGenreId: bs.Book.RakutenGenreID,
			CreatedAt:      datetime.TimeToString(bs.Book.CreatedAt),
			UpdatedAt:      datetime.TimeToString(bs.Book.UpdatedAt),
			Authors:        authors,
		}

		res.Book = book
	}

	if bs.Review != nil {
		review := &pb.BookshelfResponse_Review{
			Score:      int32(bs.Review.Score),
			Impression: bs.Review.Impression,
		}

		res.Review = review
	}

	return res
}

func getBookshelfListResponse(bss []*book.Bookshelf, limit, offset, total int) *pb.BookshelfListResponse {
	bookshelves := make([]*pb.BookshelfListResponse_Bookshelf, len(bss))
	for i, bs := range bss {
		bookshelf := &pb.BookshelfListResponse_Bookshelf{
			Id:        int64(bs.ID),
			BookId:    int64(bs.BookID),
			UserId:    bs.UserID,
			ReviewId:  int64(bs.ReviewID),
			Status:    int32(bs.Status),
			ReadOn:    datetime.DateToString(bs.ReadOn),
			CreatedAt: datetime.TimeToString(bs.CreatedAt),
			UpdatedAt: datetime.TimeToString(bs.UpdatedAt),
		}

		if bs.Book != nil {
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
				RakutenSize:    bs.Book.RakutenSize,
				RakutenGenreId: bs.Book.RakutenGenreID,
				CreatedAt:      datetime.TimeToString(bs.CreatedAt),
				UpdatedAt:      datetime.TimeToString(bs.UpdatedAt),
				Authors:        authors,
			}

			bookshelf.Book = book
		}

		bookshelves[i] = bookshelf
	}

	return &pb.BookshelfListResponse{
		Bookshelves: bookshelves,
		Limit:       int64(limit),
		Offset:      int64(offset),
		Total:       int64(total),
	}
}

func getReviewResponse(rv *book.Review) *pb.ReviewResponse {
	return &pb.ReviewResponse{
		Id:         int64(rv.ID),
		BookId:     int64(rv.BookID),
		UserId:     rv.UserID,
		Score:      int32(rv.Score),
		Impression: rv.Impression,
		CreatedAt:  datetime.TimeToString(rv.CreatedAt),
		UpdatedAt:  datetime.TimeToString(rv.UpdatedAt),
	}
}

func getReviewListResponse(rvs []*book.Review, limit, offset, total int) *pb.ReviewListResponse {
	reviews := make([]*pb.ReviewListResponse_Review, len(rvs))
	for i, rv := range rvs {
		review := &pb.ReviewListResponse_Review{
			Id:         int64(rv.ID),
			BookId:     int64(rv.BookID),
			UserId:     rv.UserID,
			Score:      int32(rv.Score),
			Impression: rv.Impression,
			CreatedAt:  datetime.TimeToString(rv.CreatedAt),
			UpdatedAt:  datetime.TimeToString(rv.UpdatedAt),
		}

		reviews[i] = review
	}

	return &pb.ReviewListResponse{
		Reviews: reviews,
		Limit:   int64(limit),
		Offset:  int64(offset),
		Total:   int64(total),
	}
}
