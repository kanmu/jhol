package subcmd

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

type IsHoliday struct {
	Date string `arg:"" default:"" help:"Target date (default: today)."`
}

func (cmd *IsHoliday) Run(binds *Binds) error {
	client := binds.Client
	var date time.Time

	if cmd.Date != "" {
		var err error
		date, err = dateparse.ParseAny(cmd.Date)

		if err != nil {
			return err
		}
	} else {
		date = time.Now()
	}

	isHoliday, err := client.IsHoliday(date)

	if err != nil {
		return err
	}

	fmt.Fprintln(binds.Out, isHoliday)

	return nil
}
