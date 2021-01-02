package datetime

import "time"

const layout = "2006-01-02 15:04:05"

// TimeToString - Time型 -> String型
func TimeToString(t time.Time) string {
	str := t.Format(layout)
	return str
}

// StringToTime - String型 -> Time型
func StringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}
