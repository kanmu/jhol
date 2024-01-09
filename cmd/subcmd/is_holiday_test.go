package subcmd_test

import (
	"strings"
	"testing"

	"github.com/kanmu/jhol/cmd/subcmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsHolidayCmd_True(_t *testing.T) {
	assert := assert.New(_t)
	require := require.New(_t)
	out := &strings.Builder{}

	cmd := &subcmd.IsHoliday{
		Date: "2024-07-15",
	}

	err := cmd.Run(&subcmd.Binds{
		Client: TestClient,
		Out:    out,
	})

	require.NoError(err)
	assert.Equal("true\n", out.String())
}

func TestIsHolidayCmd_False(_t *testing.T) {
	assert := assert.New(_t)
	require := require.New(_t)
	out := &strings.Builder{}

	cmd := &subcmd.IsHoliday{
		Date: "2023-07-19",
	}

	err := cmd.Run(&subcmd.Binds{
		Client: TestClient,
		Out:    out,
	})

	require.NoError(err)
	assert.Equal("false\n", out.String())
}
