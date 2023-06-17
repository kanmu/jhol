package jhol_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/jhol"
)

func TestClientBetween(_t *testing.T) {
	assert := assert.New(_t)
	client := jhol.NewClient(TestGCalAPIKey)

	type expectedHoliday struct {
		expectedDate string
		expectedName string
	}

	tests := []struct {
		from     string
		to       string
		holidays []expectedHoliday
	}{
		{"2022-01-01 00:00:00", "2022-12-31 23:59:59", []expectedHoliday{
			{"2022-01-01", "元日"},
			{"2022-01-10", "成人の日"},
			{"2022-02-11", "建国記念の日"},
			{"2022-02-23", "天皇誕生日"},
			{"2022-03-21", "春分の日"},
			{"2022-04-29", "昭和の日"},
			{"2022-05-03", "憲法記念日"},
			{"2022-05-04", "みどりの日"},
			{"2022-05-05", "こどもの日"},
			{"2022-07-18", "海の日"},
			{"2022-08-11", "山の日"},
			{"2022-09-19", "敬老の日"},
			{"2022-09-23", "秋分の日"},
			{"2022-10-10", "スポーツの日"},
			{"2022-11-03", "文化の日"},
			{"2022-11-23", "勤労感謝の日"},
		}},
		{"2022-01-10 00:00:00", "2022-11-03 23:59:59", []expectedHoliday{
			{"2022-01-10", "成人の日"},
			{"2022-02-11", "建国記念の日"},
			{"2022-02-23", "天皇誕生日"},
			{"2022-03-21", "春分の日"},
			{"2022-04-29", "昭和の日"},
			{"2022-05-03", "憲法記念日"},
			{"2022-05-04", "みどりの日"},
			{"2022-05-05", "こどもの日"},
			{"2022-07-18", "海の日"},
			{"2022-08-11", "山の日"},
			{"2022-09-19", "敬老の日"},
			{"2022-09-23", "秋分の日"},
			{"2022-10-10", "スポーツの日"},
			{"2022-11-03", "文化の日"},
		}},
	}

	for _, t := range tests {
		from, _ := time.ParseInLocation("2006-01-02 15:04:05", t.from, JST)
		to, _ := time.ParseInLocation("2006-01-02 15:04:05", t.to, JST)
		holidays, err := client.Between(context.Background(), from, to)

		if !assert.NoErrorf(err, "%+v", t) {
			continue
		}

		expected := []*jhol.Holiday{}

		for _, h := range t.holidays {
			expectedDate, _ := time.ParseInLocation("2006-01-02", h.expectedDate, JST)
			expected = append(expected, &jhol.Holiday{Date: expectedDate, Name: h.expectedName})
		}

		assert.Equal(expected, holidays)
	}
}

func TestClientBetween_UTC(_t *testing.T) {
	assert := assert.New(_t)
	client := jhol.NewClient(TestGCalAPIKey)

	type expectedHoliday struct {
		expectedDate string
		expectedName string
	}

	tests := []struct {
		from     string
		to       string
		holidays []expectedHoliday
	}{
		{"2022-01-01 00:00:00", "2022-12-31 23:59:59", []expectedHoliday{
			{"2022-01-01", "元日"},
			{"2022-01-10", "成人の日"},
			{"2022-02-11", "建国記念の日"},
			{"2022-02-23", "天皇誕生日"},
			{"2022-03-21", "春分の日"},
			{"2022-04-29", "昭和の日"},
			{"2022-05-03", "憲法記念日"},
			{"2022-05-04", "みどりの日"},
			{"2022-05-05", "こどもの日"},
			{"2022-07-18", "海の日"},
			{"2022-08-11", "山の日"},
			{"2022-09-19", "敬老の日"},
			{"2022-09-23", "秋分の日"},
			{"2022-10-10", "スポーツの日"},
			{"2022-11-03", "文化の日"},
			{"2022-11-23", "勤労感謝の日"},
			{"2023-01-01", "元日"},
		}},
	}

	for _, t := range tests {
		from, _ := time.ParseInLocation("2006-01-02 15:04:05", t.from, time.UTC)
		to, _ := time.ParseInLocation("2006-01-02 15:04:05", t.to, time.UTC)
		holidays, err := client.Between(context.Background(), from, to)

		if !assert.NoErrorf(err, "%+v", t) {
			continue
		}

		expected := []*jhol.Holiday{}

		for _, h := range t.holidays {
			expectedDate, _ := time.ParseInLocation("2006-01-02", h.expectedDate, JST)
			expected = append(expected, &jhol.Holiday{Date: expectedDate, Name: h.expectedName})
		}

		assert.Equal(expected, holidays)
	}
}
