package date_utils

import "time"

const (
	layout = "2006-01-02T15:05:14Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(layout)
}
