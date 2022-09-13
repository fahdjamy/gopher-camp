package utils

import "time"

func DateTime(time time.Time, layout string) string {
	return time.Format(layout)
}
