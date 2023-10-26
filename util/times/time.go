package times

import (
	"time"
)

func StringToTime(t string) (time.Time, error) {
	return time.Parse(time.RFC3339, t)
}
