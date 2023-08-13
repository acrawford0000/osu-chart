package model

import (
	"strings"
	"time"
)

// DateTime is used to parse various timestamps returned by the osu! API.
// The API returns dates in more than one format.
type DateTime time.Time

func (t *DateTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	var layout string

	if strings.Contains(s, "T") {
		if strings.Contains(s, "Z") {
			layout = "2006-01-02T15:04:05Z"
		} else {
			layout = "2006-01-02T15:04:05+00:00"
		}
	} else if strings.Contains(s, ":") {
		layout = "2006-01-02 15:04:05"
	} else {
		layout = "2006-01-02"
	}

	if s == "null" {
		t = nil
		return nil
	}

	nativeTime, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	*t = DateTime(nativeTime)

	return nil
}

func (t *DateTime) GetTime() time.Time {
	return time.Time(*t)
}

func (t *DateTime) String() string {
	return t.GetTime().Format("2006-01-02 15:04:05.000000")
}
