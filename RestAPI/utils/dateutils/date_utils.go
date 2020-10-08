package dateutils

import "time"

const (
	apiDateLayout = "2020-10-05T5:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	dateCreated := GetNow().Format(apiDateLayout)
	return dateCreated
}
