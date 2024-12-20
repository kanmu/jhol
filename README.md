# jhol

[![test](https://github.com/kanmu/jhol/actions/workflows/test.yml/badge.svg)](https://github.com/kanmu/jhol/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/kanmu/jhol.svg)](https://pkg.go.dev/github.com/kanmu/jhol)
[![Go Report Card](https://goreportcard.com/badge/github.com/kanmu/jhol)](https://goreportcard.com/report/github.com/kanmu/jhol)

Go language library to check Japanese holidays using [NAOJ Calendar](https://calendar.google.com/calendar/embed?src=2bk907eqjut8imoorgq1qa4olc%40group.calendar.google.com).

```sh
curl -s -H "X-goog-api-key: $JHOL_API_KEY" \
  'https://www.googleapis.com/calendar/v3/calendars/2bk907eqjut8imoorgq1qa4olc%40group.calendar.google.com/events?showDeleted=false&singleEvents=true&orderBy=startTime&timeMin=2019-01-01T00:00:00Z&maxResults=100' \
  | jq -c '.items[] | {start, summary}'
```

## Preparation

1. Enable Google Calendar API.
    * https://developers.google.com/calendar/api/quickstart/go#enable_the_api
1. Create an API Key.
    * https://cloud.google.com/docs/authentication/api-keys#create

## Usage

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kanmu/jhol"
)

func parseDate(s string) time.Time {
	d, _ := time.ParseInLocation("2006-01-02", s, time.Local)
	return d
}

func main() {
	apiKey := os.Getenv("JHOL_API_KEY")
	client := jhol.NewClient(apiKey).WithoutContext()

	date := parseDate("2023-07-17")
	h, _ := client.Get(date)
	fmt.Println(h) //=> 2023-07-17	海の日

	yes, _ := client.IsHoliday(date)
	fmt.Println(yes) //=> true

	yes, _ = client.IsTodayHoliday()
	fmt.Println(yes)

	h, _ = client.Next(date)
	fmt.Println(h) //=> 2023-07-17	海の日

	holidays, _ := client.NextN(date, 3)
	fmt.Println(holidays) //=> [2023-07-17	海の日 2023-08-11	山の日 2023-09-18	敬老の日]

	holidays, _ = client.Between(date, parseDate("2023-08-11"))
	fmt.Println(holidays) //=> [2023-07-17	海の日 2023-08-11	山の日]
}
```

# CLI

## Installation

```
brew install kanmu/tools/jhol
```

## Usage

```
Usage: jhol --api-key=STRING <command>

Flags:
  -h, --help              Show context-sensitive help.
      --version
      --api-key=STRING    Google API Key ($JHOL_API_KEY)

Commands:
  next --api-key=STRING [<n>]
    Show next holidays.

  is-holiday --api-key=STRING [<date>]
    Check whether the specified date is a holiday.

Run "jhol <command> --help" for more information on a command.
```

```
$ export JHOL_API_KEY=...

$ jhol next
2023-07-17	海の日
2023-08-11	山の日
2023-09-18	敬老の日

$ jhol is-holiday 2023-07-17
true

$ jhol is-holiday # today
false

$ jhol next 5 -f '%Y/%m/%d(%a)'
2023/07/17(Mon)	海の日
2023/08/11(Fri)	山の日
2023/09/18(Mon)	敬老の日
2023/09/23(Sat)	秋分の日
2023/10/09(Mon)	スポーツの日
```
