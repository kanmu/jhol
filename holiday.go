package jhol

import (
	"fmt"
	"time"

	"google.golang.org/api/calendar/v3"
)

var (
	HolidayFormat           = "2006-01-02\t%s"
	JapaneseHolidayLocation *time.Location
)

func init() {
	var err error
	JapaneseHolidayLocation, err = time.LoadLocation("Asia/Tokyo")

	if err != nil {
		panic(err)
	}
}

type Holiday struct {
	Date time.Time
	Name string
}

func newHoliday(event *calendar.Event) (*Holiday, error) {
	date := event.Start.DateTime

	if date == "" {
		date = event.Start.Date
	}

	d, err := time.ParseInLocation("2006-01-02", date, JapaneseHolidayLocation)

	if err != nil {
		return nil, err
	}

	return &Holiday{
		Date: d,
		Name: event.Summary,
	}, nil
}

func (h *Holiday) String() string {
	holFmt := h.Date.Format(HolidayFormat)
	return fmt.Sprintf(holFmt, h.Name)
}
