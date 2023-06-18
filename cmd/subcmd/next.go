package subcmd

import (
	"fmt"

	"github.com/lestrrat-go/strftime"
)

type Next struct {
	N      int    `arg:"" default:"3" help:"Number to output."`
	Format string `short:"f" help:"Date format"`
}

func (cmd *Next) Run(binds *Binds) error {
	client := binds.Client
	holidays, err := client.NextN(binds.Now(), cmd.N)

	if err != nil {
		return err
	}

	for _, h := range holidays {
		if cmd.Format == "" {
			fmt.Fprintln(binds.Out, h)
		} else {
			date, err := strftime.Format(cmd.Format, h.Date)

			if err != nil {
				return err
			}

			fmt.Fprintf(binds.Out, "%s\t%s\n", date, h.Name)
		}
	}

	return nil
}
