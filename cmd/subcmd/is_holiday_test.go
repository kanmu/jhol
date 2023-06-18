package subcmd_test

import (
	"strings"
	"testing"

	"github.com/kanmu/jhol/cmd/subcmd"
	"github.com/stretchr/testify/assert"
)

func TestIsHolidayCmd_True(_t *testing.T) {
	assert := assert.New(_t)
	out := &strings.Builder{}

	cmd := &subcmd.IsHoliday{
		Date: "2023-07-17",
	}

	err := cmd.Run(&subcmd.Binds{
		Client: Client,
		Out:    out,
	})

	if !assert.NoError(err) {
		return
	}

	assert.Equal("true\n", out.String())
}

func TestIsHolidayCmd_False(_t *testing.T) {
	assert := assert.New(_t)
	out := &strings.Builder{}

	cmd := &subcmd.IsHoliday{
		Date: "2023-07-18",
	}

	err := cmd.Run(&subcmd.Binds{
		Client: Client,
		Out:    out,
	})

	if !assert.NoError(err) {
		return
	}

	assert.Equal("false\n", out.String())
}
