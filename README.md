# jhol

Go language library to check Japanese holidays using [Google Calendar](https://calendar.google.com/calendar/embed?src=ja.japanese%23holiday%40group.v.calendar.google.com).

## Preparation

1. Enable the API.
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

	"github.com/winebarrel/jhol"
)

func parseDate(s string) time.Time {
	d, _ := time.ParseInLocation("2006-01-02", s, time.Local)
	return d
}

func main() {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	client := jhol.NewClient(apiKey).WithoutContext()

	date := parseDate("2023-07-17")
	h, _ := client.Get(date)
	fmt.Println(h) //=> 2023-07-17	海の日

	yes, _ := client.IsHoliday(date)
	fmt.Println(yes) //=> true

	holidays, _ := client.NextN(date, 3)
	fmt.Println(holidays) //=> [2023-07-17	海の日 2023-08-11	山の日 2023-09-18	敬老の日]

	holidays, _ = client.Between(date, parseDate("2023-08-11"))
	fmt.Println(holidays) //=> [2023-07-17	海の日 2023-08-11	山の日]
}
```

## CLI

### Installation

```
brew install winebarrel/jhol/jhol
```

## Usage

```
Usage: main --api-key=STRING <command>

Flags:
  -h, --help              Show context-sensitive help.
      --version
      --api-key=STRING    Google API Key ($GOOGLE_API_KEY)
      --lang="ja"         Calendar language (ja, en).

Commands:
  next --api-key=STRING [<n>]
    Show next holidays.

  is-holiday --api-key=STRING [<date>]
    Check whether the specified date is a holiday.

Run "main <command> --help" for more information on a command.
```

```
$ export GOOGLE_API_KEY=...

$ jhol next
2023-07-17	海の日
2023-08-11	山の日
2023-09-18	敬老の日

$ jhol is-holiday 2023-07-17
true
```
