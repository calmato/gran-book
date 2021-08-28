package v1

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
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
	userClient pb.UserServiceClient
}

func NewBookshelfHandler(
	bookConn *grpc.ClientConn, authConn *grpc.ClientConn, userConn *grpc.ClientConn,
) BookshelfHandler {
	bc := pb.NewBookServiceClient(bookConn)
	ac := pb.NewAuthServiceClient(authConn)
	uc := pb.NewUserServiceClient(userConn)

	return &bookshelfHandler{
		bookClient: bc,
		authClient: ac,
		userClient: uc,
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

	req := &pb.ReadBookshelfV1Request{}
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

func (h *bookshelfHandler) getBookshelfResponse(out *pb.BookshelfResponse) *pb.BookshelfV1Response {
	bookshelf := &pb.BookshelfV1Response_Bookshelf{
		Id:        out.GetId(),
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

	return &pb.BookshelfV1Response{
		Id:           out.GetBook().GetId(),
		Title:        out.GetBook().GetTitle(),
		TitleKana:    out.GetBook().GetTitleKana(),
		Description:  out.GetBook().GetDescription(),
		Isbn:         out.GetBook().GetIsbn(),
		Publisher:    out.GetBook().GetPublisher(),
		PublishedOn:  out.GetBook().GetPublishedOn(),
		ThumbnailUrl: out.GetBook().GetThumbnailUrl(),
		RakutenUrl:   out.GetBook().GetRakutenUrl(),
		Size:         out.GetBook().GetRakutenSize(),
		Author:       strings.Join(authorNames, "/"),
		AuthorKana:   strings.Join(authorNameKanas, "/"),
		CreatedAt:    out.GetBook().GetCreatedAt(),
		UpdatedAt:    out.GetBook().GetUpdatedAt(),
		Bookshelf:    bookshelf,
	}
}

// TODO: refactor
type BookshelfListV1Response struct {
	Books  []*pb.BookshelfListV1Response_Book `protobuf:"bytes,1,rep,name=books,proto3" json:"booksList,omitempty"` // 書籍一覧
	Limit  int64                              `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`    // 取得上限数
	Offset int64                              `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`  // 取得開始位置
	Total  int64                              `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`    // 検索一致数
}

func (h *bookshelfHandler) getBookshelfListResponse(out *pb.BookshelfListResponse) *BookshelfListV1Response {
	books := make([]*pb.BookshelfListV1Response_Book, len(out.GetBookshelves()))
	for i, b := range out.GetBookshelves() {
		bookshelf := &pb.BookshelfListV1Response_Bookshelf{
			Id:        b.GetId(),
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

		book := &pb.BookshelfListV1Response_Book{
			Id:           b.GetBook().GetId(),
			Title:        b.GetBook().GetTitle(),
			TitleKana:    b.GetBook().GetTitleKana(),
			Description:  b.GetBook().GetDescription(),
			Isbn:         b.GetBook().GetIsbn(),
			Publisher:    b.GetBook().GetPublisher(),
			PublishedOn:  b.GetBook().GetPublishedOn(),
			ThumbnailUrl: b.GetBook().GetThumbnailUrl(),
			RakutenUrl:   b.GetBook().GetRakutenUrl(),
			Size:         b.GetBook().GetRakutenSize(),
			Author:       strings.Join(authorNames, "/"),
			AuthorKana:   strings.Join(authorNameKanas, "/"),
			CreatedAt:    b.GetBook().GetCreatedAt(),
			UpdatedAt:    b.GetBook().GetUpdatedAt(),
			Bookshelf:    bookshelf,
		}

		books[i] = book
	}

	return &BookshelfListV1Response{
		Books:  books,
		Limit:  out.GetLimit(),
		Offset: out.GetOffset(),
		Total:  out.GetTotal(),
	}
}
