package subcmd_test

import (
	"strings"
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"github.com/kanmu/jhol/cmd/subcmd"
	"github.com/stretchr/testify/assert"
)

func TestNextCmd_3(_t *testing.T) {
	assert := assert.New(_t)
	out := &strings.Builder{}

	cmd := &subcmd.Next{
		N: 3,
	}

	err := cmd.Run(&subcmd.Binds{
		Client: Client,
		Out:    out,
		Now:    func() time.Time { return dateparse.MustParse("2023-07-17") },
	})

	if !assert.NoError(err) {
		return
	}

	assert.Equal(`2023-07-17	海の日
2023-08-11	山の日
2023-09-18	敬老の日
`, out.String())
}

func TestNextCmd_5(_t *testing.T) {
	assert := assert.New(_t)
	out := &strings.Builder{}

	cmd := &subcmd.Next{
		N: 5,
	}

	err := cmd.Run(&subcmd.Binds{
		Client: Client,
		Out:    out,
		Now:    func() time.Time { return dateparse.MustParse("2023-07-17") },
	})

	if !assert.NoError(err) {
		return
	}

	assert.Equal(`2023-07-17	海の日
2023-08-11	山の日
2023-09-18	敬老の日
2023-09-23	秋分の日
2023-10-09	スポーツの日
`, out.String())
}
