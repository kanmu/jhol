package subcmd

import (
	"fmt"
)

type Next struct {
	N int `arg:"" default:"3" help:"Number to output."`
}

func (cmd *Next) Run(binds *Binds) error {
	client := binds.Client
	holidays, err := client.NextN(binds.Now(), cmd.N)

	if err != nil {
		return err
	}

	for _, h := range holidays {
		fmt.Fprintln(binds.Out, h)
	}

	return nil
}
