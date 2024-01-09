package jhol_test

import (
	"context"
	"testing"
	"time"

	"github.com/kanmu/jhol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientNextN(_t *testing.T) {
	assert := assert.New(_t)
	require := require.New(_t)
	client := jhol.NewClient(TestGCalAPIKey)

	type expectedHoliday struct {
		expectedDate string
		expectedName string
	}

	tests := []struct {
		date     string
		holidays []expectedHoliday
	}{
		{"2024-01-01 00:00:00", []expectedHoliday{
			{"2024-01-01", "元日"},
			{"2024-01-08", "成人の日"},
			{"2024-02-11", "建国記念の日"},
		}},
		{"2024-01-01 23:59:59", []expectedHoliday{
			{"2024-01-01", "元日"},
			{"2024-01-08", "成人の日"},
			{"2024-02-11", "建国記念の日"},
		}},
		{"2024-01-02 00:00:00", []expectedHoliday{
			{"2024-01-08", "成人の日"},
			{"2024-02-11", "建国記念の日"},
			{"2024-02-12", "建国記念の日 振替休日"},
		}},
		{"2024-05-03 00:00:00", []expectedHoliday{
			{"2024-05-03", "憲法記念日"},
			{"2024-05-04", "みどりの日"},
			{"2024-05-05", "こどもの日"},
		}},
		{"2024-05-03 23:59:59", []expectedHoliday{
			{"2024-05-03", "憲法記念日"},
			{"2024-05-04", "みどりの日"},
			{"2024-05-05", "こどもの日"},
		}},
		{"2024-05-04 00:00:00", []expectedHoliday{
			{"2024-05-04", "みどりの日"},
			{"2024-05-05", "こどもの日"},
			{"2024-05-06", "こどもの日 振替休日"},
		}},
		{"2024-11-23 00:00:00", []expectedHoliday{
			{"2024-11-23", "勤労感謝の日"},
			{"2025-01-01", "元日"},
			{"2025-01-13", "成人の日"},
		}},
		{"2024-11-23 23:59:59", []expectedHoliday{
			{"2024-11-23", "勤労感謝の日"},
			{"2025-01-01", "元日"},
			{"2025-01-13", "成人の日"},
		}},
		{"2024-11-24 00:00:00", []expectedHoliday{
			{"2025-01-01", "元日"},
			{"2025-01-13", "成人の日"},
			{"2025-02-11", "建国記念の日"},
		}},
	}

	for _, t := range tests {
		d, _ := time.ParseInLocation("2006-01-02 15:04:05", t.date, JST)
		holidays, err := client.NextN(context.Background(), d, 3)
		require.NoErrorf(err, t.date)
		expected := []*jhol.Holiday{}

		for _, h := range t.holidays {
			expectedDate, _ := time.ParseInLocation("2006-01-02", h.expectedDate, JST)
			expected = append(expected, &jhol.Holiday{Date: expectedDate, Name: h.expectedName})
		}

		assert.Equalf(expected, holidays, "%v\n!= %v\n", expected, holidays)
	}
}

func TestClientNextN_UTC(_t *testing.T) {
	assert := assert.New(_t)
	require := require.New(_t)
	client := jhol.NewClient(TestGCalAPIKey)

	type expectedHoliday struct {
		expectedDate string
		expectedName string
	}

	tests := []struct {
		date     string
		holidays []expectedHoliday
	}{
		{"2024-01-01 00:00:00", []expectedHoliday{
			{"2024-01-01", "元日"},
			{"2024-01-08", "成人の日"},
			{"2024-02-11", "建国記念の日"},
		}},
		{"2024-01-01 14:59:59", []expectedHoliday{
			{"2024-01-01", "元日"},
			{"2024-01-08", "成人の日"},
			{"2024-02-11", "建国記念の日"},
		}},
		{"2024-01-01 15:00:00", []expectedHoliday{
			{"2024-01-08", "成人の日"},
			{"2024-02-11", "建国記念の日"},
			{"2024-02-12", "建国記念の日 振替休日"},
		}},
	}

	for _, t := range tests {
		d, _ := time.ParseInLocation("2006-01-02 15:04:05", t.date, time.UTC)
		holidays, err := client.NextN(context.Background(), d, 3)
		require.NoErrorf(err, t.date)
		expected := []*jhol.Holiday{}

		for _, h := range t.holidays {
			expectedDate, _ := time.ParseInLocation("2006-01-02", h.expectedDate, JST)
			expected = append(expected, &jhol.Holiday{Date: expectedDate, Name: h.expectedName})
		}

		assert.Equalf(expected, holidays, "%v\n!= %v\n", expected, holidays)
	}
}

func TestClientNext(_t *testing.T) {
	assert := assert.New(_t)
	require := require.New(_t)
	client := jhol.NewClient(TestGCalAPIKey)

	tests := []struct {
		date         string
		expectedDate string
		expectedName string
	}{
		{"2024-01-01 00:00:00", "2024-01-01", "元日"},
		{"2024-01-01 23:59:59", "2024-01-01", "元日"},
		{"2024-01-02 00:00:00", "2024-01-08", "成人の日"},
		{"2024-05-03 00:00:00", "2024-05-03", "憲法記念日"},
		{"2024-05-03 23:59:59", "2024-05-03", "憲法記念日"},
		{"2024-05-04 00:00:00", "2024-05-04", "みどりの日"},
		{"2024-11-23 00:00:00", "2024-11-23", "勤労感謝の日"},
		{"2024-11-23 23:59:59", "2024-11-23", "勤労感謝の日"},
		{"2024-11-24 00:00:00", "2025-01-01", "元日"},
	}

	for _, t := range tests {
		d, _ := time.ParseInLocation("2006-01-02 15:04:05", t.date, JST)
		h, err := client.Next(context.Background(), d)
		require.NoErrorf(err, t.date)

		expectedDate, _ := time.ParseInLocation("2006-01-02", t.expectedDate, JST)
		expected := &jhol.Holiday{Date: expectedDate, Name: t.expectedName}
		assert.Equalf(expected, h, "%v\n!= %v\n", expected, h)
	}
}
