package subcmd_test

import (
	"os"
	"testing"

	"github.com/kanmu/jhol"
)

var (
	TestClient *jhol.ClientWithoutContext
)

func TestMain(m *testing.M) {
	apiKey := os.Getenv("TEST_GCAL_API_KEY")

	if apiKey == "" {
		panic("$TEST_GCAL_API_KEY is empty")
	}

	TestClient = jhol.NewClient(apiKey).WithoutContext()

	m.Run()
}
