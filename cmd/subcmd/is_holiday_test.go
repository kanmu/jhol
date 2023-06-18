package subcmd_test

import (
	"strings"
	"testing"

	"github.com/kanmu/jhol/cmd/subcmd"
	"github.com/stretchr/testify/assert"
)

func TestIsHolidayCmd_True(_t *testing.T) {
	assert := assert.New(_t)

	cmd := &subcmd.IsHoliday{
		Date: "2023-07-17",
	}

	out := &strings.Builder{}

	cmd.Run(&subcmd.Binds{
		Client: Client,
		Out:    out,
	})

	assert.Equal("true\n", out.String())
}

func TestIsHolidayCmd_False(_t *testing.T) {
	assert := assert.New(_t)

	cmd := &subcmd.IsHoliday{
		Date: "2023-07-18",
	}

	out := &strings.Builder{}

	cmd.Run(&subcmd.Binds{
		Client: Client,
		Out:    out,
	})

	assert.Equal("false\n", out.String())
}
