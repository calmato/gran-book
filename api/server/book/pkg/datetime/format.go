package datetime

import (
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	dateFormat = "2006-01-02"
)

// TimeToString - Time型 -> String型
func TimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	str := t.Local().Format(timeFormat)
	return str
}

// DateToString - Time型 -> String型
func DateToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	str := t.Local().Format(dateFormat)
	return str
}

// StringToTime - String型 -> Time型
func StringToTime(str string) time.Time {
	t, err := time.ParseInLocation(timeFormat, str, time.Local)
	if err != nil {
		return time.Time{}
	}

	return t.Local()
}

// StringToDate - String型 -> Time型
func StringToDate(str string) time.Time {
	t, err := time.ParseInLocation(dateFormat, str, time.Local)
	if err != nil {
		return time.Time{}
	}

	return t.Local()
}

func BeginningOfMonth(str string) time.Time {
	t, err := time.ParseInLocation(dateFormat, str, time.Local)
	if err != nil {
		return time.Time{}
	}

	t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	return t.Local()
}

func EndOfMonth(str string) time.Time {
	t, err := time.ParseInLocation(dateFormat, str, time.Local)
	if err != nil {
		return time.Time{}
	}

	t = time.Date(t.Year(), t.Month()+1, 1, 23, 59, 59, 0, time.Local).AddDate(0, 0, -1)
	return t.Local()
}
