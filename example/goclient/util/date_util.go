package util

import "time"

// GetCurrentTimeStamp gets current timestamp.
func GetCurrentTimeStamp() string {
	return time.Now().Format("20060102150405")
}
