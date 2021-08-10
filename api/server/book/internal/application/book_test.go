package application

import (
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
)

func testBook(id int) *book.Book {
	current := time.Now().Local()

	return &book.Book{
		ID:             id,
		Title:          "小説　ちはやふる　上の句",
		TitleKana:      "ショウセツ チハヤフルカミノク",
		Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
		Isbn:           "9784062938426",
		Publisher:      "講談社",
		PublishedOn:    "2018-01-16",
		ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
		RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
		RakutenSize:    "コミック",
		RakutenGenreID: "001004008001/001004008003/001019001",
		CreatedAt:      current,
		UpdatedAt:      current,
		Authors: []*book.Author{
			{
				ID:        1,
				Name:      "有沢 ゆう希",
				NameKana:  "アリサワ ユウキ",
				CreatedAt: current,
				UpdatedAt: current,
			},
			{
				ID:        2,
				Name:      "末次 由紀",
				NameKana:  "スエツグ ユキ",
				CreatedAt: current,
				UpdatedAt: current,
			},
		},
	}
}
