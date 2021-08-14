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

type BookHandler interface {
	ListReview(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetReview(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type bookHandler struct {
	bookClient pb.BookServiceClient
	authClient pb.AuthServiceClient
	userClient pb.UserServiceClient
}

func NewBookHandler(bookConn *grpc.ClientConn, authConn *grpc.ClientConn, userConn *grpc.ClientConn) BookHandler {
	bc := pb.NewBookServiceClient(bookConn)
	ac := pb.NewAuthServiceClient(authConn)
	uc := pb.NewUserServiceClient(userConn)

	return &bookHandler{
		bookClient: bc,
		authClient: ac,
		userClient: uc,
	}
}

// ListReview - 書籍のレビュー一覧取得
func (h *bookHandler) ListReview(ctx *gin.Context) {
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.LIST_LIMIT_DEFAULT))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.LIST_OFFSET_DEFAULT))
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	reviewsInput := &pb.ListBookReviewRequest{
		BookId: bookID,
		Limit:  limit,
		Offset: offset,
	}

	reviewsOuput, err := h.bookClient.ListBookReview(ctx, reviewsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	userIDs := make([]string, len(reviewsOuput.GetReviews()))
	for i, r := range reviewsOuput.GetReviews() {
		userIDs[i] = r.GetUserId()
	}

	usersInput := &pb.MultiGetUserRequest{
		UserIds: userIDs,
	}

	usersOutput, err := h.userClient.MultiGetUser(ctx, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookReviewListResponse(reviewsOuput, usersOutput)
	ctx.JSON(http.StatusOK, res)
}

// Get - 書籍情報取得 (ISBN指定) ※廃止予定
func (h *bookHandler) Get(ctx *gin.Context) {
	isbn := ctx.Param("isbn")

	authOutput, err := h.authClient.GetAuth(ctx, &pb.Empty{})
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bookInput := &pb.GetBookByIsbnRequest{
		Isbn: isbn,
	}

	bookOutput, err := h.bookClient.GetBookByIsbn(ctx, bookInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bookshelfInput := &pb.GetBookshelfRequest{
		BookId: bookOutput.GetId(),
		UserId: authOutput.GetId(),
	}

	bookshelfOutput, err := h.bookClient.GetBookshelf(ctx, bookshelfInput)
	if err != nil && !util.IsNotFound(err) {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookResponse(bookOutput, bookshelfOutput)
	ctx.JSON(http.StatusOK, res)
}

// GetReview - 書籍のレビュー情報取得
func (h *bookHandler) GetReview(ctx *gin.Context) {
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	reviewID, err := strconv.ParseInt(ctx.Param("reviewID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	bookInput := &pb.GetBookRequest{
		BookId: bookID,
	}

	_, err = h.bookClient.GetBook(ctx, bookInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	reviewInput := &pb.GetReviewRequest{
		ReviewId: reviewID,
	}

	reviewOutput, err := h.bookClient.GetReview(ctx, reviewInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	userInput := &pb.GetUserRequest{
		UserId: reviewOutput.GetUserId(),
	}

	userOutput, err := h.userClient.GetUser(ctx, userInput)
	if err != nil && !util.IsNotFound(err) {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookReviewResponse(reviewOutput, userOutput)
	ctx.JSON(http.StatusOK, res)
}

// Create - 書籍登録
func (h *bookHandler) Create(ctx *gin.Context) {
	req := &request.CreateBookRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	authorNames := strings.Split(req.Author, "/")
	authorNameKanas := strings.Split(req.AuthorKana, "/")
	authors := make([]*pb.CreateBookRequest_Author, len(authorNames))
	for i := range authorNames {
		author := &pb.CreateBookRequest_Author{
			Name:     authorNames[i],
			NameKana: authorNameKanas[i],
		}

		authors[i] = author
	}

	in := &pb.CreateBookRequest{
		Title:          req.Title,
		TitleKana:      req.TitleKana,
		Description:    req.ItemCaption,
		Isbn:           req.Isbn,
		Publisher:      req.PublisherName,
		PublishedOn:    req.SalesDate,
		ThumbnailUrl:   h.getThumbnailURLByRequest(req.SmaillImageURL, req.MediumImageURL, req.LargeImageURL),
		RakutenUrl:     req.ItemURL,
		RakutenSize:    req.Size,
		RakutenGenreId: req.BooksGenreID,
		Authors:        authors,
	}

	out, err := h.bookClient.CreateBook(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookResponse(out, nil)
	ctx.JSON(http.StatusOK, res)
}

// Update - 書籍更新
func (h *bookHandler) Update(ctx *gin.Context) {
	req := &request.UpdateBookRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	authorNames := strings.Split(req.Author, "/")
	authorNameKanas := strings.Split(req.AuthorKana, "/")
	authors := make([]*pb.UpdateBookRequest_Author, len(authorNames))
	for i := range authorNames {
		author := &pb.UpdateBookRequest_Author{
			Name:     authorNames[i],
			NameKana: authorNameKanas[i],
		}

		authors[i] = author
	}

	in := &pb.UpdateBookRequest{
		Title:          req.Title,
		TitleKana:      req.TitleKana,
		Description:    req.ItemCaption,
		Isbn:           req.Isbn,
		Publisher:      req.PublisherName,
		PublishedOn:    req.SalesDate,
		ThumbnailUrl:   h.getThumbnailURLByRequest(req.SmaillImageURL, req.MediumImageURL, req.LargeImageURL),
		RakutenUrl:     req.ItemURL,
		RakutenSize:    req.Size,
		RakutenGenreId: req.BooksGenreID,
		Authors:        authors,
	}

	out, err := h.bookClient.UpdateBook(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookResponse(out, nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *bookHandler) getThumbnailURLByRequest(smallImageURL, mediumImageURL, largeImageURL string) string {
	if largeImageURL != "" {
		return largeImageURL
	}

	if mediumImageURL != "" {
		return mediumImageURL
	}

	if smallImageURL != "" {
		return smallImageURL
	}

	return ""
}

func (h *bookHandler) getBookResponse(
	bookOutput *pb.BookResponse, bookshelfOutput *pb.BookshelfResponse,
) *response.BookResponse {
	authorNames := make([]string, len(bookOutput.GetAuthors()))
	authorNameKanas := make([]string, len(bookOutput.GetAuthors()))
	for i, a := range bookOutput.GetAuthors() {
		authorNames[i] = a.GetName()
		authorNameKanas[i] = a.GetNameKana()
	}

	res := &response.BookResponse{
		ID:           bookOutput.GetId(),
		Title:        bookOutput.GetTitle(),
		TitleKana:    bookOutput.GetTitleKana(),
		Description:  bookOutput.GetDescription(),
		Isbn:         bookOutput.GetIsbn(),
		Publisher:    bookOutput.GetPublisher(),
		PublishedOn:  bookOutput.GetPublishedOn(),
		ThumbnailURL: bookOutput.GetThumbnailUrl(),
		RakutenURL:   bookOutput.GetRakutenUrl(),
		Size:         bookOutput.GetRakutenSize(),
		Author:       strings.Join(authorNames, "/"),
		AuthorKana:   strings.Join(authorNameKanas, "/"),
		CreatedAt:    bookOutput.GetCreatedAt(),
		UpdatedAt:    bookOutput.GetUpdatedAt(),
	}

	if bookshelfOutput != nil {
		bookshelf := &response.BookBookshelf{
			ID:        bookshelfOutput.GetId(),
			Status:    entity.BookshelfStatus(bookshelfOutput.GetStatus()).Name(),
			ReadOn:    bookshelfOutput.GetReadOn(),
			CreatedAt: bookshelfOutput.GetCreatedAt(),
			UpdatedAt: bookshelfOutput.GetUpdatedAt(),
		}

		if bookshelfOutput.GetReview() != nil {
			bookshelf.Impression = bookshelfOutput.GetReview().GetImpression()
		}

		res.Bookshelf = bookshelf
	}

	return res
}

func (h *bookHandler) getBookReviewResponse(
	reviewOutput *pb.ReviewResponse, userOutput *pb.UserResponse,
) *response.BookReviewResponse {
	user := &response.BookReviewUser{
		ID:       reviewOutput.GetUserId(),
		Username: "unknown",
	}

	if userOutput != nil {
		user.Username = userOutput.GetUsername()
		user.ThumbnailURL = userOutput.GetThumbnailUrl()
	}

	return &response.BookReviewResponse{
		ID:         reviewOutput.GetId(),
		Impression: reviewOutput.GetImpression(),
		CreatedAt:  reviewOutput.GetCreatedAt(),
		UpdatedAt:  reviewOutput.GetUpdatedAt(),
		User:       user,
	}
}

func (h *bookHandler) getBookReviewListResponse(
	reviewsOutput *pb.ReviewListResponse, usersOutput *pb.UserMapResponse,
) *response.BookReviewListResponse {
	reviews := make([]*response.BookReviewListReview, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		user := &response.BookReviewListUser{
			Username: "unknown",
		}

		if usersOutput.GetUsers()[r.GetUserId()] != nil {
			user.Username = usersOutput.GetUsers()[r.GetUserId()].GetUsername()
			user.ThumbnailURL = usersOutput.GetUsers()[r.GetUserId()].GetThumbnailUrl()
		}

		review := &response.BookReviewListReview{
			ID:         r.GetId(),
			Impression: r.GetImpression(),
			CreatedAt:  r.GetCreatedAt(),
			UpdatedAt:  r.GetUpdatedAt(),
			User:       user,
		}

		reviews[i] = review
	}

	return &response.BookReviewListResponse{
		Reviews: reviews,
		Limit:   reviewsOutput.GetLimit(),
		Offset:  reviewsOutput.GetOffset(),
		Total:   reviewsOutput.GetTotal(),
	}
}
