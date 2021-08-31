package v2

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v2"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type BookshelfHandler interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type bookshelfHandler struct {
	bookClient pb.BookServiceClient
	userClient pb.UserServiceClient
}

func NewBookshelfHandler(bookConn *grpc.ClientConn, userConn *grpc.ClientConn) BookshelfHandler {
	bc := pb.NewBookServiceClient(bookConn)
	uc := pb.NewUserServiceClient(userConn)

	return &bookshelfHandler{
		bookClient: bc,
		userClient: uc,
	}
}

// List - 本棚の書籍一覧取得
func (h *bookshelfHandler) List(ctx *gin.Context) {
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))
	userID := ctx.Param("userID")

	bookshelvesInput := &pb.ListBookshelfRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}

	c := util.SetMetadata(ctx)
	bookshelvesOutput, err := h.bookClient.ListBookshelf(c, bookshelvesInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bss := entity.NewBookshelves(bookshelvesOutput.Bookshelves)

	booksInput := &pb.MultiGetBooksRequest{
		BookIds: bss.BookIDs(),
	}

	booksOutput, err := h.bookClient.MultiGetBooks(c, booksInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bs := entity.NewBooks(booksOutput.Books)
	res := h.getBookshelfListResponse(
		bss, bs.Map(), bookshelvesOutput.Limit, bookshelvesOutput.Offset, bookshelvesOutput.Total,
	)
	ctx.JSON(http.StatusOK, res)
}

// Get - 本棚の書籍情報取得
func (h *bookshelfHandler) Get(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	c := util.SetMetadata(ctx)
	eg, ectx := errgroup.WithContext(c)

	var b *entity.Book
	eg.Go(func() error {
		in := &pb.GetBookRequest{
			BookId: bookID,
		}
		out, err := h.bookClient.GetBook(ectx, in)
		if err != nil {
			return err
		}
		b = entity.NewBook(out.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		in := &pb.GetBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		out, err := h.bookClient.GetBookshelf(ectx, in)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(out.Bookshelf)
		return nil
	})

	var rs entity.Reviews
	var limit, offset, total int64
	eg.Go(func() error {
		in := &pb.ListBookReviewRequest{
			BookId: bookID,
			Limit:  20,
			Offset: 0,
		}
		out, err := h.bookClient.ListBookReview(ectx, in)
		if err != nil {
			return err
		}
		limit = out.Limit
		offset = out.Offset
		total = out.Total
		rs = entity.NewReviews(out.Reviews)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	usersInput := &pb.MultiGetUserRequest{
		UserIds: rs.UserIDs(),
	}

	usersOutput, err := h.userClient.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := entity.NewUsers(usersOutput.Users)
	res := h.getBookshelfResponse(bs, b, rs, us.Map(), limit, offset, total)
	ctx.JSON(http.StatusOK, res)
}

func (h *bookshelfHandler) getBookshelfResponse(
	bs *entity.Bookshelf, b *entity.Book,
	rs entity.Reviews, us map[string]*entity.User,
	reviewLimit int64,
	reviewOffset int64,
	reviewTotal int64,
) *response.BookshelfResponse {
	bookshelf := &response.BookshelfResponse_Bookshelf{
		Status:    bs.Status().Name(),
		ReadOn:    bs.ReadOn,
		ReviewID:  bs.ReviewId,
		CreatedAt: bs.CreatedAt,
		UpdatedAt: bs.UpdatedAt,
	}

	reviews := make([]*response.BookshelfResponse_Review, len(rs))
	for i, r := range rs {
		user := &response.BookshelfResponse_User{
			ID:           r.UserId,
			Username:     "unknown",
			ThumbnailURL: "",
		}

		if us[r.UserId] != nil {
			user.Username = us[r.UserId].Username
			user.ThumbnailURL = us[r.UserId].ThumbnailUrl
		}

		review := &response.BookshelfResponse_Review{
			ID:         r.Id,
			Impression: r.Impression,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
			User:       user,
		}

		reviews[i] = review
	}

	return &response.BookshelfResponse{
		ID:           b.Id,
		Title:        b.Title,
		TitleKana:    b.TitleKana,
		Description:  b.Description,
		Isbn:         b.Isbn,
		Publisher:    b.Publisher,
		PublishedOn:  b.PublishedOn,
		ThumbnailURL: b.ThumbnailUrl,
		RakutenURL:   b.RakutenUrl,
		Size:         b.RakutenSize,
		Author:       strings.Join(b.AuthorNames(), "/"),
		AuthorKana:   strings.Join(b.AuthorNameKanas(), "/"),
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
		Bookshelf:    bookshelf,
		Reviews:      reviews,
		ReviewLimit:  reviewLimit,
		ReviewOffset: reviewOffset,
		ReviewTotal:  reviewTotal,
	}
}

func (h *bookshelfHandler) getBookshelfListResponse(
	bss entity.Bookshelves, bm map[int64]*entity.Book, limit, offset, total int64,
) *response.BookshelfListResponse {
	books := make([]*response.BookshelfListResponse_Book, 0, len(bss))
	for _, bs := range bss {
		b, ok := bm[bs.BookId]
		if !ok {
			continue
		}

		bookshelf := &response.BookshelfListResponse_Bookshelf{
			Status:    bs.Status().Name(),
			ReadOn:    bs.ReadOn,
			ReviewID:  bs.ReviewId,
			CreatedAt: bs.CreatedAt,
			UpdatedAt: bs.UpdatedAt,
		}

		book := &response.BookshelfListResponse_Book{
			ID:           b.Id,
			Title:        b.Title,
			TitleKana:    b.TitleKana,
			Description:  b.Description,
			Isbn:         b.Isbn,
			Publisher:    b.Publisher,
			PublishedOn:  b.PublishedOn,
			ThumbnailURL: b.ThumbnailUrl,
			RakutenURL:   b.RakutenUrl,
			Size:         b.RakutenSize,
			Author:       strings.Join(b.AuthorNames(), "/"),
			AuthorKana:   strings.Join(b.AuthorNameKanas(), "/"),
			CreatedAt:    b.CreatedAt,
			UpdatedAt:    b.UpdatedAt,
			Bookshelf:    bookshelf,
		}

		books = append(books, book)
	}

	return &response.BookshelfListResponse{
		Books:  books,
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}
