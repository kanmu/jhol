package jhol_test

import (
	"context"
	"testing"
	"time"

	"github.com/kanmu/jhol"
	"github.com/stretchr/testify/assert"
)

func TestClientNextN(_t *testing.T) {
	assert := assert.New(_t)
	client := jhol.NewClient(TestGCalAPIKey)

	type expectedHoliday struct {
		expectedDate string
		expectedName string
	}

	tests := []struct {
		date     string
		holidays []expectedHoliday
	}{
		{"2022-01-01 00:00:00", []expectedHoliday{
			{"2022-01-01", "元日"},
			{"2022-01-10", "成人の日"},
			{"2022-02-11", "建国記念の日"},
		}},
		{"2022-01-01 23:59:59", []expectedHoliday{
			{"2022-01-01", "元日"},
			{"2022-01-10", "成人の日"},
			{"2022-02-11", "建国記念の日"},
		}},
		{"2022-01-02 00:00:00", []expectedHoliday{
			{"2022-01-10", "成人の日"},
			{"2022-02-11", "建国記念の日"},
			{"2022-02-23", "天皇誕生日"},
		}},
		{"2022-05-03 00:00:00", []expectedHoliday{
			{"2022-05-03", "憲法記念日"},
			{"2022-05-04", "みどりの日"},
			{"2022-05-05", "こどもの日"},
		}},
		{"2022-05-03 23:59:59", []expectedHoliday{
			{"2022-05-03", "憲法記念日"},
			{"2022-05-04", "みどりの日"},
			{"2022-05-05", "こどもの日"},
		}},
		{"2022-05-04 00:00:00", []expectedHoliday{
			{"2022-05-04", "みどりの日"},
			{"2022-05-05", "こどもの日"},
			{"2022-07-18", "海の日"},
		}},
		{"2022-11-23 00:00:00", []expectedHoliday{
			{"2022-11-23", "勤労感謝の日"},
			{"2023-01-01", "元日"},
			{"2023-01-02", "休日 元日"},
		}},
		{"2022-11-23 23:59:59", []expectedHoliday{
			{"2022-11-23", "勤労感謝の日"},
			{"2023-01-01", "元日"},
			{"2023-01-02", "休日 元日"},
		}},
		{"2022-11-24 00:00:00", []expectedHoliday{
			{"2023-01-01", "元日"},
			{"2023-01-02", "休日 元日"},
			{"2023-01-09", "成人の日"},
		}},
	}

	for _, t := range tests {
		d, _ := time.ParseInLocation("2006-01-02 15:04:05", t.date, JST)
		holidays, err := client.NextN(context.Background(), d, 3)

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

func TestClientNextN_UTC(_t *testing.T) {
	assert := assert.New(_t)
	client := jhol.NewClient(TestGCalAPIKey)

	type expectedHoliday struct {
		expectedDate string
		expectedName string
	}

	tests := []struct {
		date     string
		holidays []expectedHoliday
	}{
		{"2022-01-01 00:00:00", []expectedHoliday{
			{"2022-01-01", "元日"},
			{"2022-01-10", "成人の日"},
			{"2022-02-11", "建国記念の日"},
		}},
		{"2022-01-01 14:59:59", []expectedHoliday{
			{"2022-01-01", "元日"},
			{"2022-01-10", "成人の日"},
			{"2022-02-11", "建国記念の日"},
		}},
		{"2022-01-01 15:00:00", []expectedHoliday{
			{"2022-01-10", "成人の日"},
			{"2022-02-11", "建国記念の日"},
			{"2022-02-23", "天皇誕生日"},
		}},
	}

	for _, t := range tests {
		d, _ := time.ParseInLocation("2006-01-02 15:04:05", t.date, time.UTC)
		holidays, err := client.NextN(context.Background(), d, 3)

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

func TestClientNext(_t *testing.T) {
	assert := assert.New(_t)
	client := jhol.NewClient(TestGCalAPIKey)

	tests := []struct {
		date         string
		expectedDate string
		expectedName string
	}{
		{"2022-01-01 00:00:00", "2022-01-01", "元日"},
		{"2022-01-01 23:59:59", "2022-01-01", "元日"},
		{"2022-01-02 00:00:00", "2022-01-10", "成人の日"},
		{"2022-05-03 00:00:00", "2022-05-03", "憲法記念日"},
		{"2022-05-03 23:59:59", "2022-05-03", "憲法記念日"},
		{"2022-05-04 00:00:00", "2022-05-04", "みどりの日"},
		{"2022-11-23 00:00:00", "2022-11-23", "勤労感謝の日"},
		{"2022-11-23 23:59:59", "2022-11-23", "勤労感謝の日"},
		{"2022-11-24 00:00:00", "2023-01-01", "元日"},
	}

	for _, t := range tests {
		d, _ := time.ParseInLocation("2006-01-02 15:04:05", t.date, JST)
		h, err := client.Next(context.Background(), d)

		if !assert.NoErrorf(err, "%+v", t) {
			continue
		}

		expectedDate, _ := time.ParseInLocation("2006-01-02", t.expectedDate, JST)
		expected := &jhol.Holiday{Date: expectedDate, Name: t.expectedName}
		assert.Equal(expected, h)
	}
}
