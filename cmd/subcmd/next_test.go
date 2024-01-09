package subcmd_test

import (
	"strings"
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"github.com/kanmu/jhol/cmd/subcmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNextCmd_3(_t *testing.T) {
	assert := assert.New(_t)
	require := require.New(_t)
	out := &strings.Builder{}

	cmd := &subcmd.Next{
		N: 3,
	}

	err := cmd.Run(&subcmd.Binds{
		Client: TestClient,
		Out:    out,
		Now:    func() time.Time { return dateparse.MustParse("2024-07-15") },
	})

	require.NoError(err)

	assert.Equal(`2024-07-15	海の日
2024-08-11	山の日
2024-08-12	休日 山の日
`, out.String())
}

func TestNextCmd_5(_t *testing.T) {
	assert := assert.New(_t)
	require := require.New(_t)
	out := &strings.Builder{}

	cmd := &subcmd.Next{
		N: 5,
	}

	err := cmd.Run(&subcmd.Binds{
		Client: TestClient,
		Out:    out,
		Now:    func() time.Time { return dateparse.MustParse("2024-07-15") },
	})

	require.NoError(err)

	assert.Equal(`2024-07-15	海の日
2024-08-11	山の日
2024-08-12	休日 山の日
2024-09-16	敬老の日
2024-09-22	秋分の日
`, out.String())
}

func TestNextCmd_Format(_t *testing.T) {
	assert := assert.New(_t)
	require := require.New(_t)
	out := &strings.Builder{}

	cmd := &subcmd.Next{
		N:      3,
		Format: "%Y/%m/%d(%a)",
	}

	err := cmd.Run(&subcmd.Binds{
		Client: TestClient,
		Out:    out,
		Now:    func() time.Time { return dateparse.MustParse("2024-07-15") },
	})

	require.NoError(err)

	assert.Equal(`2024/07/15(Mon)	海の日
2024/08/11(Sun)	山の日
2024/08/12(Mon)	休日 山の日
`, out.String())
}
