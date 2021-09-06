package test

import (
	"errors"
	"time"
)

var (
	ErrMock = errors.New("some error")

	jst, _   = time.LoadLocation("Asia/Tokyo")
	TimeMock = time.Date(2021, time.Month(7), 24, 20, 0, 0, 0, jst).Local().String()
	DateMock = time.Date(2021, time.Month(7), 24, 0, 0, 0, 0, jst).Local().String()
)
