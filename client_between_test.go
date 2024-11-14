package jhol_test

import (
	"context"
	"testing"
	"time"

	"github.com/kanmu/jhol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientBetween(_t *testing.T) {
	assert := assert.New(_t)
	require := require.New(_t)
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
		{"2024-01-01 00:00:00", "2024-12-31 23:59:59", []expectedHoliday{
			{"2024-01-01", "元日"},
			{"2024-01-08", "成人の日"},
			{"2024-02-11", "建国記念の日"},
			{"2024-02-12", "休日"},
			{"2024-02-23", "天皇誕生日"},
			{"2024-03-20", "春分の日"},
			{"2024-04-29", "昭和の日"},
			{"2024-05-03", "憲法記念日"},
			{"2024-05-04", "みどりの日"},
			{"2024-05-05", "こどもの日"},
			{"2024-05-06", "休日"},
			{"2024-07-15", "海の日"},
			{"2024-08-11", "山の日"},
			{"2024-08-12", "休日"},
			{"2024-09-16", "敬老の日"},
			{"2024-09-22", "秋分の日"},
			{"2024-09-23", "秋分の日 振替休日"},
			{"2024-10-14", "スポーツの日"},
			{"2024-11-03", "文化の日"},
			{"2024-11-04", "文化の日 振替休日"},
			{"2024-11-23", "勤労感謝の日"},
		}},
		{"2024-01-08 00:00:00", "2024-11-04 23:59:59", []expectedHoliday{
			{"2024-01-08", "成人の日"},
			{"2024-02-11", "建国記念の日"},
			{"2024-02-12", "休日"},
			{"2024-02-23", "天皇誕生日"},
			{"2024-03-20", "春分の日"},
			{"2024-04-29", "昭和の日"},
			{"2024-05-03", "憲法記念日"},
			{"2024-05-04", "みどりの日"},
			{"2024-05-05", "こどもの日"},
			{"2024-05-06", "休日"},
			{"2024-07-15", "海の日"},
			{"2024-08-11", "山の日"},
			{"2024-08-12", "休日"},
			{"2024-09-16", "敬老の日"},
			{"2024-09-22", "秋分の日"},
			{"2024-09-23", "秋分の日 振替休日"},
			{"2024-10-14", "スポーツの日"},
			{"2024-11-03", "文化の日"},
			{"2024-11-04", "文化の日 振替休日"},
		}},
	}

	for _, t := range tests {
		from, _ := time.ParseInLocation("2006-01-02 15:04:05", t.from, JST)
		to, _ := time.ParseInLocation("2006-01-02 15:04:05", t.to, JST)
		holidays, err := client.Between(context.Background(), from, to)
		require.NoErrorf(err, "%s - %s", t.from, t.to)
		expected := []*jhol.Holiday{}

		for _, h := range t.holidays {
			expectedDate, _ := time.ParseInLocation("2006-01-02", h.expectedDate, JST)
			expected = append(expected, &jhol.Holiday{Date: expectedDate, Name: h.expectedName})
		}

		assert.Equalf(expected, holidays, "%v\n!= %v\n", expected, holidays)
	}
}

func TestClientBetween_UTC(_t *testing.T) {
	assert := assert.New(_t)
	require := require.New(_t)
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
		{"2024-01-01 00:00:00", "2024-12-31 23:59:59", []expectedHoliday{
			{"2024-01-01", "元日"},
			{"2024-01-08", "成人の日"},
			{"2024-02-11", "建国記念の日"},
			{"2024-02-12", "休日"},
			{"2024-02-23", "天皇誕生日"},
			{"2024-03-20", "春分の日"},
			{"2024-04-29", "昭和の日"},
			{"2024-05-03", "憲法記念日"},
			{"2024-05-04", "みどりの日"},
			{"2024-05-05", "こどもの日"},
			{"2024-05-06", "休日"},
			{"2024-07-15", "海の日"},
			{"2024-08-11", "山の日"},
			{"2024-08-12", "休日"},
			{"2024-09-16", "敬老の日"},
			{"2024-09-22", "秋分の日"},
			{"2024-09-23", "秋分の日 振替休日"},
			{"2024-10-14", "スポーツの日"},
			{"2024-11-03", "文化の日"},
			{"2024-11-04", "文化の日 振替休日"},
			{"2024-11-23", "勤労感謝の日"},
			{"2025-01-01", "元日"},
		}},
	}

	for _, t := range tests {
		from, _ := time.ParseInLocation("2006-01-02 15:04:05", t.from, time.UTC)
		to, _ := time.ParseInLocation("2006-01-02 15:04:05", t.to, time.UTC)
		holidays, err := client.Between(context.Background(), from, to)
		require.NoErrorf(err, "%s - %s", t.from, t.to)
		expected := []*jhol.Holiday{}

		for _, h := range t.holidays {
			expectedDate, _ := time.ParseInLocation("2006-01-02", h.expectedDate, JST)
			expected = append(expected, &jhol.Holiday{Date: expectedDate, Name: h.expectedName})
		}

		assert.Equalf(expected, holidays, "%v\n!= %v\n", expected, holidays)
	}
}
