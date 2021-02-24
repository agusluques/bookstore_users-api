package utils

import "time"

const (
	apiDateFormat   = "2006-01-02T15:04:05Z"
	apiDBDateFormat = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}

func GetNowDBString() string {
	return GetNow().Format(apiDBDateFormat)
}
