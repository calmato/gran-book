package entity

// BookshelfStatus - 読書ステータス
type BookshelfStatus int32

const (
	BOOKSHELF_STATUS_NONE    BookshelfStatus = 0 // 不明なステータス
	BOOKSHELF_STATUS_READ    BookshelfStatus = 1 // 読み終えた本
	BOOKSHELF_STATUS_READING BookshelfStatus = 2 // 読んでいる本
	BOOKSHELF_STATUS_STACKED BookshelfStatus = 3 // 積読本
	BOOKSHELF_STATUS_WANT    BookshelfStatus = 4 // 読みたい本
	BOOKSHELF_STATUS_RELEASE BookshelfStatus = 5 // 手放したい本
)

var (
	bookshelfStatusName = map[BookshelfStatus]string{
		0: "none",
		1: "read",
		2: "reading",
		3: "stack",
		4: "want",
		5: "release",
	}
	bookshelfStatusValue = map[string]int32{
		"none":    0,
		"read":    1,
		"reading": 2,
		"stack":   3,
		"want":    4,
		"release": 5,
	}
)

func (s BookshelfStatus) Name() string {
	if name, ok := bookshelfStatusName[s]; ok {
		return name
	}

	return ""
}

func (s BookshelfStatus) Value(key string) BookshelfStatus {
	if value, ok := bookshelfStatusValue[key]; ok {
		return BookshelfStatus(value)
	}

	return BOOKSHELF_STATUS_NONE
}
