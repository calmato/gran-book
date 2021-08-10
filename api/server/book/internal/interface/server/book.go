package server

import (
	"github.com/calmato/gran-book/api/server/book/internal/application"
	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/internal/interface/validation"
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

func getBookListResponse(bs []*book.Book, limit, offset, total int) *pb.BookListResponse {
	books := make([]*pb.BookListResponse_Book, len(bs))
	for i, b := range bs {
		authors := make([]*pb.BookListResponse_Author, len(b.Authors))
		for j, a := range b.Authors {
			author := &pb.BookListResponse_Author{
				Name:     a.Name,
				NameKana: a.NameKana,
			}

			authors[j] = author
		}

		book := &pb.BookListResponse_Book{
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

		books[i] = book
	}

	return &pb.BookListResponse{
		Books:  books,
		Limit:  int64(limit),
		Offset: int64(offset),
		Total:  int64(total),
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

	res := &pb.BookshelfResponse{
		Id:        int64(bs.ID),
		BookId:    int64(bs.BookID),
		UserId:    bs.UserID,
		ReviewId:  int64(bs.ReviewID),
		Status:    int32(bs.Status),
		ReadOn:    datetime.DateToString(bs.ReadOn),
		CreatedAt: datetime.TimeToString(bs.CreatedAt),
		UpdatedAt: datetime.TimeToString(bs.UpdatedAt),
		Book:      book,
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

		bookshelf := &pb.BookshelfListResponse_Bookshelf{
			Id:        int64(bs.ID),
			BookId:    int64(bs.BookID),
			UserId:    bs.UserID,
			ReviewId:  int64(bs.ReviewID),
			Status:    int32(bs.Status),
			ReadOn:    datetime.DateToString(bs.ReadOn),
			CreatedAt: datetime.TimeToString(bs.CreatedAt),
			UpdatedAt: datetime.TimeToString(bs.UpdatedAt),
			Book:      book,
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
