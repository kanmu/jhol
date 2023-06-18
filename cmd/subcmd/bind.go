package subcmd

import (
	"io"
	"time"

	"github.com/kanmu/jhol"
)

type Binds struct {
	Client *jhol.ClientWithoutContext
	Out    io.Writer
	Now    func() time.Time
}
