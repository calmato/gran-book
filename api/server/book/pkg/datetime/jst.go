package datetime

import (
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	dateFormat = "2006-01-02"
)

var (
	jst = time.FixedZone("JST", 9*60*60)
)

func Now() time.Time {
	return time.Now().In(jst)
}

func BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, jst)
}

func EndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 1, 23, 59, 59, 0, jst).AddDate(0, 0, -1)
}

func ParseTime(str string) (time.Time, error) {
	return time.ParseInLocation(timeFormat, str, jst)
}

func ParseDate(str string) (time.Time, error) {
	return time.ParseInLocation(dateFormat, str, jst)
}

func FormatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format(timeFormat)
}

func FormatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format(dateFormat)
}
