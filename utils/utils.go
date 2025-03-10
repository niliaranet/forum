package utils

import (
	"time"
)

func FormatTimestamp(timestamp time.Time) string {
	return timestamp.Local().Format("2006-01-02 - 03:04 PM")
}
