package jhol_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/jhol"
)

func TestClientIsHoliday(_t *testing.T) {
	assert := assert.New(_t)
	client := jhol.NewClient(TestGCalAPIKey)

	tests := []struct {
		date string
	}{
		{"2022-01-01"},
		{"2022-01-10"},
		{"2022-02-11"},
		{"2022-02-23"},
		{"2022-03-21"},
		{"2022-04-29"},
		{"2022-05-03"},
		{"2022-05-04"},
		{"2022-05-05"},
		{"2022-07-18"},
		{"2022-08-11"},
		{"2022-09-19"},
		{"2022-09-23"},
		{"2022-10-10"},
		{"2022-11-03"},
		{"2022-11-23"},
	}

	for _, t := range tests {
		d, _ := time.ParseInLocation("2006-01-02", t.date, JST)
		ok, err := client.IsHoliday(context.Background(), d)

		if !assert.NoErrorf(err, "%+v", t) {
			continue
		}

		assert.Truef(ok, "%+v", t)
	}
}

func TestClientIsNotHoliday(t *testing.T) {
	assert := assert.New(t)
	client := jhol.NewClient(TestGCalAPIKey)

	tests := []struct {
		date string
	}{
		{"2022-01-02"}, // +1
		{"2022-01-09"}, // -1
		{"2022-02-12"}, // +1
		{"2022-02-22"}, // -1
		{"2022-03-22"}, // +1
		{"2022-04-28"}, // -1
		{"2022-05-13"}, // +10
		{"2022-05-14"}, // +10
		{"2022-05-15"}, // +10
		{"2022-07-17"}, // -1
		{"2022-08-12"}, // +1
		{"2022-09-18"}, // -1
		{"2022-09-24"}, // +1
		{"2022-10-09"}, // -1
		{"2022-11-04"}, // +1
		{"2022-11-22"}, // -1
	}

	for _, t := range tests {
		d, _ := time.ParseInLocation("2006-01-02", t.date, JST)
		ok, err := client.IsHoliday(context.Background(), d)

		if !assert.NoErrorf(err, "%+v", t) {
			continue
		}

		assert.Falsef(ok, "%+v", t)
	}
}
