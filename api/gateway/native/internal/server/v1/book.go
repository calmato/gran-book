package v1

import (
	"net/http"
	"strings"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type BookHandler interface {
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type bookHandler struct {
	bookClient pb.BookServiceClient
	authClient pb.AuthServiceClient
}

func NewBookHandler(bookConn *grpc.ClientConn, authConn *grpc.ClientConn) BookHandler {
	bc := pb.NewBookServiceClient(bookConn)
	ac := pb.NewAuthServiceClient(authConn)

	return &bookHandler{
		bookClient: bc,
		authClient: ac,
	}
}

// Get - 書籍情報取得 (ISBN指定) ※廃止予定
func (h *bookHandler) Get(ctx *gin.Context) {
	isbn := ctx.Param("bookID")

	c := util.SetMetadata(ctx)
	authOutput, err := h.authClient.GetAuth(c, &pb.Empty{})
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bookInput := &pb.GetBookByIsbnRequest{
		Isbn: isbn,
	}

	bookOutput, err := h.bookClient.GetBookByIsbn(c, bookInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bookshelfInput := &pb.GetBookshelfRequest{
		BookId: bookOutput.GetId(),
		UserId: authOutput.GetId(),
	}

	bookshelfOutput, err := h.bookClient.GetBookshelf(c, bookshelfInput)
	if err != nil && !util.IsNotFound(err) {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookResponse(bookOutput, bookshelfOutput)
	ctx.JSON(http.StatusOK, res)
}

// Create - 書籍登録
func (h *bookHandler) Create(ctx *gin.Context) {
	req := &pb.CreateBookV1Request{}
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
		ThumbnailUrl:   h.getThumbnailURLByRequest(req.SmallImageUrl, req.MediumImageUrl, req.LargeImageUrl),
		RakutenUrl:     req.ItemUrl,
		RakutenSize:    req.Size,
		RakutenGenreId: req.BooksGenreId,
		Authors:        authors,
	}

	c := util.SetMetadata(ctx)
	out, err := h.bookClient.CreateBook(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookResponse(out, nil)
	ctx.JSON(http.StatusOK, res)
}

// Update - 書籍更新
func (h *bookHandler) Update(ctx *gin.Context) {
	req := &pb.UpdateBookV1Request{}
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
		ThumbnailUrl:   h.getThumbnailURLByRequest(req.SmallImageUrl, req.MediumImageUrl, req.LargeImageUrl),
		RakutenUrl:     req.ItemUrl,
		RakutenSize:    req.Size,
		RakutenGenreId: req.BooksGenreId,
		Authors:        authors,
	}

	c := util.SetMetadata(ctx)
	out, err := h.bookClient.UpdateBook(c, in)
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

type BookV1Response struct {
	Id           int64                     `json:"id,omitempty"`           // 書籍ID
	Title        string                    `json:"title,omitempty"`        // タイトル
	TitleKana    string                    `json:"title_kana,omitempty"`   // タイトル(かな)
	Description  string                    `json:"description,omitempty"`  // 説明
	Isbn         string                    `json:"isbn,omitempty"`         // ISBN
	Publisher    string                    `json:"publisher,omitempty"`    // 出版社名
	PublishedOn  string                    `json:"published_on,omitempty"` // 出版日
	ThumbnailUrl string                    `json:"thumbnailUrl,omitempty"` // サムネイルURL
	RakutenUrl   string                    `json:"rakutenUrl,omitempty"`   // 楽天ショップURL
	Size         string                    `json:"size,omitempty"`         // 楽天書籍サイズ
	Author       string                    `json:"author,omitempty"`       // 著者名一覧
	AuthorKana   string                    `json:"authorKana,omitempty"`   /// 著者名一覧(かな)
	Bookshelf    *BookV1Response_Bookshelf `json:"bookshelf,omitempty"`    // ユーザーの本棚情報
	CreatedAt    string                    `json:"createdAt,omitempty"`    // 登録日時
	UpdatedAt    string                    `json:"updatedAt,omitempty"`    // 更新日時
}

type BookV1Response_Bookshelf struct {
	Id         int64  `json:"id,omitempty"`         // 本棚ID
	Status     string `json:"status,omitempty"`     // 読書ステータス
	ReadOn     string `json:"readOn,omitempty"`     // 読み終えた日
	Impression string `json:"impression,omitempty"` // 感想
	CreatedAt  string `json:"createdAt,omitempty"`  // 登録日時
	UpdatedAt  string `json:"updatedAt,omitempty"`  // 更新日時
}

func (h *bookHandler) getBookResponse(
	bookOutput *pb.BookResponse, bookshelfOutput *pb.BookshelfResponse,
) *BookV1Response {
	authorNames := make([]string, len(bookOutput.GetAuthors()))
	authorNameKanas := make([]string, len(bookOutput.GetAuthors()))
	for i, a := range bookOutput.GetAuthors() {
		authorNames[i] = a.GetName()
		authorNameKanas[i] = a.GetNameKana()
	}

	res := &BookV1Response{
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
	}

	if bookshelfOutput != nil {
		bookshelf := &BookV1Response_Bookshelf{
			Id:        bookshelfOutput.GetId(),
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
