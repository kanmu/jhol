package subcmd

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	"github.com/winebarrel/jhol"
)

type IsHoliday struct {
	Date string `arg:"" default:"" help:"Target date."`
}

func (cmd *IsHoliday) Run(client *jhol.ClientWithoutContext) error {
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

	fmt.Println(isHoliday)

	return nil
}
