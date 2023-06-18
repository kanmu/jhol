package main

import (
	"log"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/kanmu/jhol"
	"github.com/kanmu/jhol/cmd/subcmd"
)

var version string

var cli struct {
	Version   kong.VersionFlag
	APIKey    string           `required:"" env:"JHOL_API_KEY" help:"Google API Key"`
	Lang      string           `enum:"ja,en" default:"ja" help:"Calendar language (ja, en)."`
	Next      subcmd.Next      `cmd:"" help:"Show next holidays."`
	IsHoliday subcmd.IsHoliday `cmd:"" help:"Check whether the specified date is a holiday."`
}

func init() {
	log.SetFlags(0)
}

func main() {
	ctx := kong.Parse(
		&cli,
		kong.Vars{"version": version},
	)

	var calendarID string

	switch cli.Lang {
	case "ja":
		calendarID = jhol.JaJapaneseHolidayCalendar
	case "en":
		calendarID = jhol.EnJapaneseHolidayCalendar
	}

	client := jhol.NewClientWithCalendar(cli.APIKey, calendarID).WithoutContext()
	err := ctx.Run(&subcmd.Binds{
		Client: client,
		Out:    os.Stdout,
		Now:    time.Now,
	})

	if err != nil {
		log.Fatalln(err)
	}
}
