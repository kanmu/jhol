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

	yes, _ = client.IsTodayHoliday()
	fmt.Println(yes)

	h, _ = client.Next(date)
	fmt.Println(h) //=> 2023-07-17	海の日

	holidays, _ := client.NextN(date, 3)
	fmt.Println(holidays) //=> [2023-07-17	海の日 2023-08-11	山の日 2023-09-18	敬老の日]

	holidays, _ = client.Between(date, parseDate("2023-08-11"))
	fmt.Println(holidays) //=> [2023-07-17	海の日 2023-08-11	山の日]
}
