package date_utils

import "time"

const (
	apilayout = "2006-01-02T15:05:14Z"
	dblayout  = "2006-01-02 15:05:14Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apilayout)
}

func GetNowDBString() string {
	return GetNow().Format(dblayout)
}
