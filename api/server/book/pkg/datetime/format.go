package datetime

import "time"

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
	t, _ := time.ParseInLocation(timeFormat, str, time.Local)
	return t.Local()
}

// StringToDate - String型 -> Time型
func StringToDate(str string) time.Time {
	t, _ := time.ParseInLocation(dateFormat, str, time.Local)
	return t.Local()
}
