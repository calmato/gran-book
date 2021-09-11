package datetime

import "time"

type Time struct {
	time time.Time
}

var (
	jst = time.FixedZone("JST", 9*60*60)
)

func New(t time.Time) *Time {
	return &Time{t}
}

func Now() time.Time {
	return time.Now().In(jst)
}

func (t *Time) BeginningOfMonth() time.Time {
	return time.Date(t.time.Year(), t.time.Month(), 1, 0, 0, 0, 0, jst)
}

func (t *Time) EndOfMonth() time.Time {
	return time.Date(t.time.Year(), t.time.Month()+1, 1, 23, 59, 59, 0, jst).AddDate(0, 0, -1)
}
