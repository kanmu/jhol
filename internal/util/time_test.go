package util_test

import (
	"testing"
	"time"

	"github.com/kanmu/jhol/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestTimeZoneName(_t *testing.T) {
	assert := assert.New(_t)

	loadLoc := func(tz string) *time.Location {
		loc, _ := time.LoadLocation(tz)
		return loc
	}

	tests := []struct {
		loc      *time.Location
		expected string
	}{
		{loadLoc("Asia/Tokyo"), "JST"},
		{loadLoc("UTC"), "UTC"},
		{loadLoc("America/Los_Angeles"), "PST"},
	}

	for _, t := range tests {
		tm, _ := time.ParseInLocation("2006-01-02", "2020-03-04", t.loc)
		assert.Equal(t.expected, util.TimeZoneName(tm))
	}
}
