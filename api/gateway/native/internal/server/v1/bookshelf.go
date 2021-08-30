package v1

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/native/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type BookshelfHandler interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Read(ctx *gin.Context)
	Reading(ctx *gin.Context)
	Stacked(ctx *gin.Context)
	Want(ctx *gin.Context)
	Release(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookshelfHandler struct {
	bookClient pb.BookServiceClient
	authClient pb.AuthServiceClient
}

func NewBookshelfHandler(bookConn *grpc.ClientConn, authConn *grpc.ClientConn) BookshelfHandler {
	bc := pb.NewBookServiceClient(bookConn)
	ac := pb.NewAuthServiceClient(authConn)

	return &bookshelfHandler{
		bookClient: bc,
		authClient: ac,
	}
}

// List - 本棚の書籍一覧取得 ※廃止予定
func (h *bookshelfHandler) List(ctx *gin.Context) {
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))
	userID := ctx.Param("userID")

	in := &pb.ListBookshelfRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}

	c := util.SetMetadata(ctx)
	out, err := h.bookClient.ListBookshelf(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfListResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Get - 本棚の書籍情報取得 ※廃止予定
func (h *bookshelfHandler) Get(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.GetBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.bookClient.GetBookshelf(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Read - 読んだ本の登録
func (h *bookshelfHandler) Read(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	req := &request.ReadBookshelfRequest{}
	err = ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.ReadBookshelfRequest{
		UserId:     userID,
		BookId:     bookID,
		Impression: req.Impression,
		ReadOn:     req.ReadOn,
	}

	c := util.SetMetadata(ctx)
	out, err := h.bookClient.ReadBookshelf(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Reading - 読んでいる本の登録
func (h *bookshelfHandler) Reading(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.ReadingBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.bookClient.ReadingBookshelf(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Stacked - 積読本の登録
func (h *bookshelfHandler) Stacked(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.StackedBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.bookClient.StackedBookshelf(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Want - 欲しい本の登録
func (h *bookshelfHandler) Want(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.WantBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.bookClient.WantBookshelf(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Release - 手放したい本の登録
func (h *bookshelfHandler) Release(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.ReleaseBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.bookClient.ReleaseBookshelf(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Delete - 本棚から書籍の削除
func (h *bookshelfHandler) Delete(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.DeleteBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}

	c := util.SetMetadata(ctx)
	_, err = h.bookClient.DeleteBookshelf(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *bookshelfHandler) getBookshelfResponse(out *pb.BookshelfResponse) *response.BookshelfResponse {
	bookshelf := &response.BookshelfResponse_Bookshelf{
		ID:        out.GetId(),
		Status:    entity.BookshelfStatus(out.GetStatus()).Name(),
		ReadOn:    out.GetReadOn(),
		CreatedAt: out.GetCreatedAt(),
		UpdatedAt: out.GetUpdatedAt(),
	}

	if out.GetReview() != nil {
		bookshelf.Impression = out.GetReview().GetImpression()
	}

	authorNames := make([]string, len(out.GetBook().GetAuthors()))
	authorNameKanas := make([]string, len(out.GetBook().GetAuthors()))
	for i, a := range out.GetBook().GetAuthors() {
		authorNames[i] = a.GetName()
		authorNameKanas[i] = a.GetNameKana()
	}

	return &response.BookshelfResponse{
		ID:           out.GetBook().GetId(),
		Title:        out.GetBook().GetTitle(),
		TitleKana:    out.GetBook().GetTitleKana(),
		Description:  out.GetBook().GetDescription(),
		Isbn:         out.GetBook().GetIsbn(),
		Publisher:    out.GetBook().GetPublisher(),
		PublishedOn:  out.GetBook().GetPublishedOn(),
		ThumbnailURL: out.GetBook().GetThumbnailUrl(),
		RakutenURL:   out.GetBook().GetRakutenUrl(),
		Size:         out.GetBook().GetRakutenSize(),
		Author:       strings.Join(authorNames, "/"),
		AuthorKana:   strings.Join(authorNameKanas, "/"),
		CreatedAt:    out.GetBook().GetCreatedAt(),
		UpdatedAt:    out.GetBook().GetUpdatedAt(),
		Bookshelf:    bookshelf,
	}
}

func (h *bookshelfHandler) getBookshelfListResponse(out *pb.BookshelfListResponse) *response.BookshelfListResponse {
	books := make([]*response.BookshelfListResponse_Book, len(out.GetBookshelves()))
	for i, b := range out.GetBookshelves() {
		bookshelf := &response.BookshelfListResponse_Bookshelf{
			ID:        b.GetId(),
			Status:    entity.BookshelfStatus(b.GetStatus()).Name(),
			ReadOn:    b.GetReadOn(),
			CreatedAt: b.GetCreatedAt(),
			UpdatedAt: b.GetUpdatedAt(),
		}

		authorNames := make([]string, len(b.GetBook().GetAuthors()))
		authorNameKanas := make([]string, len(b.GetBook().GetAuthors()))
		for i, a := range b.GetBook().GetAuthors() {
			authorNames[i] = a.GetName()
			authorNameKanas[i] = a.GetNameKana()
		}

		book := &response.BookshelfListResponse_Book{
			ID:           b.GetBook().GetId(),
			Title:        b.GetBook().GetTitle(),
			TitleKana:    b.GetBook().GetTitleKana(),
			Description:  b.GetBook().GetDescription(),
			Isbn:         b.GetBook().GetIsbn(),
			Publisher:    b.GetBook().GetPublisher(),
			PublishedOn:  b.GetBook().GetPublishedOn(),
			ThumbnailURL: b.GetBook().GetThumbnailUrl(),
			RakutenURL:   b.GetBook().GetRakutenUrl(),
			Size:         b.GetBook().GetRakutenSize(),
			Author:       strings.Join(authorNames, "/"),
			AuthorKana:   strings.Join(authorNameKanas, "/"),
			CreatedAt:    b.GetBook().GetCreatedAt(),
			UpdatedAt:    b.GetBook().GetUpdatedAt(),
			Bookshelf:    bookshelf,
		}

		books[i] = book
	}

	return &response.BookshelfListResponse{
		Books:  books,
		Limit:  out.GetLimit(),
		Offset: out.GetOffset(),
		Total:  out.GetTotal(),
	}
}
