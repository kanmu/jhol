package subcmd_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/kanmu/jhol"
)

var (
	Client *jhol.ClientWithoutContext
)

func TestMain(m *testing.M) {
	fmt.Print("xx")
	apiKey := os.Getenv("TEST_GCAL_API_KEY")

	if apiKey == "" {
		panic("$TEST_GCAL_API_KEY is empty")
	}

	Client = jhol.NewClient(apiKey).WithoutContext()

	m.Run()
}
