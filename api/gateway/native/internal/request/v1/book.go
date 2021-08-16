package v1

type CreateBookRequest struct {
	Title          string `json:"title"`
	TitleKana      string `json:"titleKana"`
	ItemCaption    string `json:"itemCaption"`
	Isbn           string `json:"isbn"`
	PublisherName  string `json:"publisherName"`
	SalesDate      string `json:"salesDate"`
	SmaillImageURL string `json:"smallImageUrl"`
	MediumImageURL string `json:"mediumImageUrl"`
	LargeImageURL  string `json:"largeImageUrl"`
	ItemURL        string `json:"itemUrl"`
	Size           string `json:"size"`
	BooksGenreID   string `json:"booksGenreId"`
	Author         string `json:"author"`
	AuthorKana     string `json:"authorKana"`
}

type UpdateBookRequest struct {
	Title          string `json:"title"`
	TitleKana      string `json:"titleKana"`
	ItemCaption    string `json:"itemCaption"`
	Isbn           string `json:"isbn"`
	PublisherName  string `json:"publisherName"`
	SalesDate      string `json:"salesDate"`
	SmaillImageURL string `json:"smallImageUrl"`
	MediumImageURL string `json:"mediumImageUrl"`
	LargeImageURL  string `json:"largeImageUrl"`
	ItemURL        string `json:"itemUrl"`
	Size           string `json:"size"`
	BooksGenreID   string `json:"booksGenreId"`
	Author         string `json:"author"`
	AuthorKana     string `json:"authorKana"`
}
