package util

import "time"

func StringToTime(value string) time.Time {
	convertedTime, err := time.Parse("2006-01-02 15:04", value)

	if err != nil {
		panic(err)
	}

	return convertedTime
}
