package subcmd

import (
	"fmt"
	"time"

	"github.com/kanmu/jhol"
)

type Next struct {
	N int `arg:"" default:"3" help:"Number to output."`
}

func (cmd *Next) Run(client *jhol.ClientWithoutContext) error {
	holidays, err := client.NextN(time.Now(), cmd.N)

	if err != nil {
		return err
	}

	for _, h := range holidays {
		fmt.Println(h)
	}

	return nil
}
