package jhol_test

import (
	"os"
	"testing"
	"time"
)

var (
	TestGCalAPIKey string
	JST            *time.Location
)

func TestMain(m *testing.M) {
	TestGCalAPIKey = os.Getenv("TEST_GCAL_API_KEY")

	if TestGCalAPIKey == "" {
		panic("$TEST_GCAL_API_KEY is empty")
	}

	var err error
	JST, err = time.LoadLocation("Asia/Tokyo")

	if err != nil {
		panic(err)
	}

	m.Run()
}
