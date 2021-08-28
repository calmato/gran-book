package v2

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

// Get - 本棚の書籍情報取得
func (h *bookshelfHandler) Get(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	bookshelfInput := &pb.GetBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}

	c := util.SetMetadata(ctx)
	bookshelfOutput, err := h.bookClient.GetBookshelf(c, bookshelfInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	reviewsInput := &pb.ListBookReviewRequest{
		BookId: bookshelfOutput.GetBookId(),
		Limit:  20,
		Offset: 0,
	}

	reviewsOutput, err := h.bookClient.ListBookReview(c, reviewsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	userIDs := make([]string, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		userIDs[i] = r.GetUserId()
	}

	usersInput := &pb.MultiGetUserRequest{
		UserIds: userIDs,
	}

	usersOutput, err := h.userClient.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(bookshelfOutput, reviewsOutput, usersOutput)
	ctx.JSON(http.StatusOK, res)
}

// 本棚の書籍情報
type BookshelfV2Response struct {
	Id           int64                          `json:"id"`           // 書籍ID
	Title        string                         `json:"title"`        // タイトル
	TitleKana    string                         `json:"titleKana"`    // タイトル(かな)
	Description  string                         `json:"description"`  // 説明
	Isbn         string                         `json:"isbn"`         // ISBN
	Publisher    string                         `json:"publisher"`    // 出版社名
	PublishedOn  string                         `json:"publishedOn"`  // 出版日
	ThumbnailUrl string                         `json:"thumbnailUrl"` // サムネイルURL
	RakutenUrl   string                         `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string                         `json:"size"`         // 楽天書籍サイズ
	Author       string                         `json:"author"`       // 著者名一覧
	AuthorKana   string                         `json:"author_kana"`  /// 著者名一覧(かな)
	Bookshelf    *BookshelfV2Response_Bookshelf `json:"bookshelf"`    // ユーザーの本棚情報
	Reviews      []*BookshelfV2Response_Review  `json:"reviewsList"`  // レビュー一覧
	ReviewLimit  int64                          `json:"reviewLimit"`  // レビュー取得上限
	ReviewOffset int64                          `json:"reviewOffset"` // レビュー取得開始位置
	ReviewTotal  int64                          `json:"reviewTotal"`  // レビュー検索一致件
	CreatedAt    string                         `json:"createdAt"`    // 登録日時
	UpdatedAt    string                         `json:"updatedAt"`    // 更新日時
}

type BookshelfV2Response_Bookshelf struct {
	Status    string `json:"status"`    // 読書ステータス
	ReadOn    string `json:"readOn"`    // 読み終えた日
	ReviewId  int64  `json:"reviewId"`  // レビューID
	CreatedAt string `json:"createdAt"` // 登録日時
	UpdatedAt string `json:"updatedAt"` // 更新日時
}

type BookshelfV2Response_Review struct {
	Id         int64                     `json:"id"`         // レビューID
	Impression string                    `json:"impression"` // 感想
	User       *BookshelfV2Response_User `json:"user"`       // 投稿者
	CreatedAt  string                    `json:"createdAt"`  // 登録日時
	UpdatedAt  string                    `json:"updatedAt"`  // 更新日時
}

type BookshelfV2Response_User struct {
	Id           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl"` // サムネイルURL
}

func (h *bookshelfHandler) getBookshelfResponse(
	bookshelfOutput *pb.BookshelfResponse, reviewsOutput *pb.ReviewListResponse, usersOutput *pb.UserMapResponse,
) *BookshelfV2Response {
	bookshelf := &BookshelfV2Response_Bookshelf{
		Status:    entity.BookshelfStatus(bookshelfOutput.GetStatus()).Name(),
		ReadOn:    bookshelfOutput.GetReadOn(),
		ReviewId:  bookshelfOutput.GetReviewId(),
		CreatedAt: bookshelfOutput.GetCreatedAt(),
		UpdatedAt: bookshelfOutput.GetUpdatedAt(),
	}

	reviews := make([]*BookshelfV2Response_Review, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		user := &BookshelfV2Response_User{
			Id:       r.GetUserId(),
			Username: "unknown",
		}

		if usersOutput.GetUsers()[r.GetUserId()] != nil {
			user.Username = usersOutput.GetUsers()[r.GetUserId()].GetUsername()
			user.ThumbnailUrl = usersOutput.GetUsers()[r.GetUserId()].GetThumbnailUrl()
		}

		review := &BookshelfV2Response_Review{
			Id:         r.GetId(),
			Impression: r.GetImpression(),
			CreatedAt:  r.GetCreatedAt(),
			UpdatedAt:  r.GetUpdatedAt(),
			User:       user,
		}

		reviews[i] = review
	}

	authorNames := make([]string, len(bookshelfOutput.GetBook().GetAuthors()))
	authorNameKanas := make([]string, len(bookshelfOutput.GetBook().GetAuthors()))
	for i, a := range bookshelfOutput.GetBook().GetAuthors() {
		authorNames[i] = a.GetName()
		authorNameKanas[i] = a.GetNameKana()
	}

	return &BookshelfV2Response{
		Id:           bookshelfOutput.GetBook().GetId(),
		Title:        bookshelfOutput.GetBook().GetTitle(),
		TitleKana:    bookshelfOutput.GetBook().GetTitleKana(),
		Description:  bookshelfOutput.GetBook().GetDescription(),
		Isbn:         bookshelfOutput.GetBook().GetIsbn(),
		Publisher:    bookshelfOutput.GetBook().GetPublisher(),
		PublishedOn:  bookshelfOutput.GetBook().GetPublishedOn(),
		ThumbnailUrl: bookshelfOutput.GetBook().GetThumbnailUrl(),
		RakutenUrl:   bookshelfOutput.GetBook().GetRakutenUrl(),
		Size:         bookshelfOutput.GetBook().GetRakutenSize(),
		Author:       strings.Join(authorNames, "/"),
		AuthorKana:   strings.Join(authorNameKanas, "/"),
		CreatedAt:    bookshelfOutput.GetBook().GetCreatedAt(),
		UpdatedAt:    bookshelfOutput.GetBook().GetUpdatedAt(),
		Bookshelf:    bookshelf,
		Reviews:      reviews,
		ReviewLimit:  reviewsOutput.GetLimit(),
		ReviewOffset: reviewsOutput.GetOffset(),
		ReviewTotal:  reviewsOutput.GetTotal(),
	}
}

// 本棚の書籍一覧
type BookshelfListV2Response struct {
	Books  []*BookshelfListV2Response_Book `json:"booksList"` // 書籍一覧
	Limit  int64                           `json:"limit"`     // 取得上限数
	Offset int64                           `json:"offset"`    // 取得開始位置
	Total  int64                           `json:"total"`     // 検索一致数
}

type BookshelfListV2Response_Book struct {
	Id           int64                              `json:"id"`           // 書籍ID
	Title        string                             `json:"title"`        // タイトル
	TitleKana    string                             `json:"titleKana"`    // タイトル(かな)
	Description  string                             `json:"description"`  // 説明
	Isbn         string                             `json:"isbn"`         // ISBN
	Publisher    string                             `json:"publisher"`    // 出版社名
	PublishedOn  string                             `json:"publishedOn"`  // 出版日
	ThumbnailUrl string                             `json:"thumbnailUrl"` // サムネイルURL
	RakutenUrl   string                             `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string                             `json:"size"`         // 楽天書籍サイズ
	Author       string                             `json:"author"`       // 著者名一覧
	AuthorKana   string                             `json:"authorKana"`   /// 著者名一覧(かな)
	Bookshelf    *BookshelfListV2Response_Bookshelf `json:"bookshelf"`    // ユーザーの本棚情報
	CreatedAt    string                             `json:"createdAt"`    // 登録日時
	UpdatedAt    string                             `json:"updatedAt"`    // 更新日時
}

type BookshelfListV2Response_Bookshelf struct {
	Status    string `json:"status"`    // 読書ステータス
	ReadOn    string `json:"readOn"`    // 読み終えた日
	ReviewId  int64  `json:"reviewId"`  // レビューID
	CreatedAt string `json:"createdAt"` // 登録日時
	UpdatedAt string `json:"updatedAt"` // 更新日時
}

func (h *bookshelfHandler) getBookshelfListResponse(out *pb.BookshelfListResponse) *BookshelfListV2Response {
	books := make([]*BookshelfListV2Response_Book, len(out.GetBookshelves()))
	for i, b := range out.GetBookshelves() {
		bookshelf := &BookshelfListV2Response_Bookshelf{
			Status:    entity.BookshelfStatus(b.GetStatus()).Name(),
			ReadOn:    b.GetReadOn(),
			ReviewId:  b.GetReviewId(),
			CreatedAt: b.GetCreatedAt(),
			UpdatedAt: b.GetUpdatedAt(),
		}

		authorNames := make([]string, len(b.GetBook().GetAuthors()))
		authorNameKanas := make([]string, len(b.GetBook().GetAuthors()))
		for i, a := range b.GetBook().GetAuthors() {
			authorNames[i] = a.GetName()
			authorNameKanas[i] = a.GetNameKana()
		}

		book := &BookshelfListV2Response_Book{
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

	return &BookshelfListV2Response{
		Books:  books,
		Limit:  out.GetLimit(),
		Offset: out.GetOffset(),
		Total:  out.GetTotal(),
	}
}
