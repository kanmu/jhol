package jhol_test

import (
	"testing"
	"time"

	"github.com/kanmu/jhol"
	"github.com/stretchr/testify/assert"
)

func TestHolidayString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		date     string
		name     string
		expected string
	}{
		{"2022-03-21", "春分の日", "2022-03-21\t春分の日"},
		{"2022-04-29", "昭和の日", "2022-04-29\t昭和の日"},
	}

	loc, _ := time.LoadLocation("Asia/Tokyo")

	for _, t := range tests {
		d, _ := time.ParseInLocation("2006-01-02", t.date, loc)
		h := &jhol.Holiday{Date: d, Name: t.name}
		assert.Equal(t.expected, h.String())
	}
}
