package util

import (
	"time"
)

func TimeZoneName(t time.Time) string {
	name, _ := t.Zone()
	return name
}
