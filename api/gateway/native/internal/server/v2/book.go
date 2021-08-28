package v2

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type BookHandler interface {
	Get(ctx *gin.Context)
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

// Get - 書籍情報取得
func (h *bookHandler) Get(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	key := ctx.DefaultQuery("key", "id")

	c := util.SetMetadata(ctx)
	bookOutput, err := h.getBook(c, bookID, key)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	reviewsInput := &pb.ListBookReviewRequest{
		BookId: bookOutput.GetId(),
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

	res := h.getBookResponse(bookOutput, reviewsOutput, usersOutput)
	ctx.JSON(http.StatusOK, res)
}

func (h *bookHandler) getBook(ctx context.Context, bookID, key string) (*pb.BookResponse, error) {
	switch key {
	case "id":
		bookID, err := strconv.ParseInt(bookID, 10, 64)
		if err != nil {
			return nil, entity.ErrBadRequest.New(err)
		}

		in := &pb.GetBookRequest{
			BookId: bookID,
		}

		return h.bookClient.GetBook(ctx, in)
	case "isbn":
		in := &pb.GetBookByIsbnRequest{
			Isbn: bookID,
		}

		return h.bookClient.GetBookByIsbn(ctx, in)
	default:
		err := fmt.Errorf("this key is invalid argument")
		return nil, entity.ErrBadRequest.New(err)
	}
}

// 書籍情報
type BookV2Response struct {
	Id           int64                    `json:"id,omitempty"`           // 書籍ID
	Title        string                   `json:"title,omitempty"`        // タイトル
	TitleKana    string                   `json:"titleKana,omitempty"`    // タイトル(かな)
	Description  string                   `json:"description,omitempty"`  // 説明
	Isbn         string                   `json:"isbn,omitempty"`         // ISBN
	Publisher    string                   `json:"publisher,omitempty"`    // 出版社名
	PublishedOn  string                   `json:"publishedOn,omitempty"`  // 出版日
	ThumbnailUrl string                   `json:"thumbnailUrl,omitempty"` // サムネイルURL
	RakutenUrl   string                   `json:"rakutenUrl,omitempty"`   // 楽天ショップURL
	Size         string                   `json:"size,omitempty"`         // 楽天書籍サイズ
	Author       string                   `json:"author,omitempty"`       // 著者名一覧
	AuthorKana   string                   `json:"authorKana,omitempty"`   /// 著者名一覧(かな)
	Reviews      []*BookV2Response_Review `json:"reviewsList,omitempty"`  // レビュー一覧
	ReviewLimit  int64                    `json:"reviewLimit,omitempty"`  // レビュー取得上限
	ReviewOffset int64                    `json:"reviewOffset,omitempty"` // レビュー取得開始位置
	ReviewTotal  int64                    `json:"reviewTotal,omitempty"`  // レビュー検索一致件数
	CreatedAt    string                   `json:"createdAt,omitempty"`    // 登録日時
	UpdatedAt    string                   `json:"updatedAt,omitempty"`    // 更新日時
}

type BookV2Response_Review struct {
	Id         int64                `json:"id,omitempty"`         // レビューID
	Impression string               `json:"impression,omitempty"` // 感想
	User       *BookV2Response_User `json:"user,omitempty"`       // 投稿者
	CreatedAt  string               `json:"createdAt,omitempty"`  // 登録日時
	UpdatedAt  string               `json:"updatedAt,omitempty"`  // 更新日時
}

type BookV2Response_User struct {
	Id           string `json:"id,omitempty"`           // ユーザーID
	Username     string `json:"username,omitempty"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"` // サムネイルURL
}

func (h *bookHandler) getBookResponse(
	bookOutput *pb.BookResponse,
	reviewsOutput *pb.ReviewListResponse,
	usersOutput *pb.UserMapResponse,
) *BookV2Response {
	authorNames := make([]string, len(bookOutput.GetAuthors()))
	authorNameKanas := make([]string, len(bookOutput.GetAuthors()))
	for i, a := range bookOutput.GetAuthors() {
		authorNames[i] = a.GetName()
		authorNameKanas[i] = a.GetNameKana()
	}

	reviews := make([]*BookV2Response_Review, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		user := &BookV2Response_User{
			Id:       r.GetUserId(),
			Username: "unknown",
		}

		if usersOutput.GetUsers()[r.GetUserId()] != nil {
			user.Username = usersOutput.GetUsers()[r.GetUserId()].GetUsername()
			user.ThumbnailUrl = usersOutput.GetUsers()[r.GetUserId()].GetThumbnailUrl()
		}

		review := &BookV2Response_Review{
			Id:         r.GetId(),
			Impression: r.GetImpression(),
			CreatedAt:  r.GetCreatedAt(),
			UpdatedAt:  r.GetUpdatedAt(),
			User:       user,
		}

		reviews[i] = review
	}

	return &BookV2Response{
		Id:           bookOutput.GetId(),
		Title:        bookOutput.GetTitle(),
		TitleKana:    bookOutput.GetTitleKana(),
		Description:  bookOutput.GetDescription(),
		Isbn:         bookOutput.GetIsbn(),
		Publisher:    bookOutput.GetPublisher(),
		PublishedOn:  bookOutput.GetPublishedOn(),
		ThumbnailUrl: bookOutput.GetThumbnailUrl(),
		RakutenUrl:   bookOutput.GetRakutenUrl(),
		Size:         bookOutput.GetRakutenSize(),
		Author:       strings.Join(authorNames, "/"),
		AuthorKana:   strings.Join(authorNameKanas, "/"),
		CreatedAt:    bookOutput.GetCreatedAt(),
		UpdatedAt:    bookOutput.GetUpdatedAt(),
		Reviews:      reviews,
		ReviewLimit:  reviewsOutput.GetLimit(),
		ReviewOffset: reviewsOutput.GetOffset(),
		ReviewTotal:  reviewsOutput.GetTotal(),
	}
}
