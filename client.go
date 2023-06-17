package jhol

import (
	"context"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	EnJapaneseHolidayCalendar = "en.japanese#holiday@group.v.calendar.google.com"
	JaJapaneseHolidayCalendar = "ja.japanese#holiday@group.v.calendar.google.com"
)

type Client struct {
	APIKey     string
	CalendarID string
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		CalendarID: JaJapaneseHolidayCalendar,
	}
}

func NewClientWithCalendar(apiKey string, calendarID string) *Client {
	return &Client{
		APIKey:     apiKey,
		CalendarID: calendarID,
	}
}

type ClientWithoutContext struct {
	client *Client
}

func (client *Client) WithoutContext() *ClientWithoutContext {
	return &ClientWithoutContext{client: client}
}

func (client *Client) newEventsListCall(ctx context.Context) (*calendar.EventsListCall, error) {
	srv, err := calendar.NewService(ctx, option.WithAPIKey(client.APIKey))

	if err != nil {
		return nil, err
	}

	eventsListCall := srv.Events.List(client.CalendarID).
		ShowDeleted(false).
		SingleEvents(true).
		TimeZone(JapaneseHolidayLocation.String()).
		OrderBy("startTime")

	return eventsListCall, nil
}

///////////////////////////////////////////////////////////////////////

func (client *Client) Get(ctx context.Context, t time.Time) (*Holiday, error) {
	eventsListCall, err := client.newEventsListCall(ctx)

	if err != nil {
		return nil, err
	}

	t = t.In(JapaneseHolidayLocation)
	min := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, JapaneseHolidayLocation)
	max := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, JapaneseHolidayLocation)
	events, err := eventsListCall.TimeMin(min.Format(time.RFC3339)).TimeMax(max.Format(time.RFC3339)).MaxResults(1).Do()

	if err != nil {
		return nil, err
	}

	if len(events.Items) == 0 {
		return nil, nil
	}

	h, err := newHoliday(events.Items[0])

	if err != nil {
		return nil, err
	}

	return h, nil
}

func (c *ClientWithoutContext) Get(t time.Time) (*Holiday, error) {
	return c.client.Get(context.Background(), t)
}

///////////////////////////////////////////////////////////////////////

func (client *Client) NextN(ctx context.Context, t time.Time, n int) ([]*Holiday, error) {
	eventsListCall, err := client.newEventsListCall(ctx)

	if err != nil {
		return nil, err
	}

	t = t.In(JapaneseHolidayLocation)
	min := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, JapaneseHolidayLocation)
	events, err := eventsListCall.TimeMin(min.Format(time.RFC3339)).MaxResults(int64(n)).Do()

	if err != nil {
		return nil, err
	}

	holidays := []*Holiday{}

	for _, item := range events.Items {
		h, err := newHoliday(item)

		if err != nil {
			return nil, err
		}

		holidays = append(holidays, h)
	}

	return holidays, nil
}

func (client *Client) Next(ctx context.Context, t time.Time) (*Holiday, error) {
	holidays, err := client.NextN(ctx, t, 1)

	if err != nil {
		return nil, err
	}

	if len(holidays) == 0 {
		return nil, nil
	}

	return holidays[0], nil
}

func (c *ClientWithoutContext) NextN(t time.Time, n int) ([]*Holiday, error) {
	return c.client.NextN(context.Background(), t, n)
}

func (c *ClientWithoutContext) Next(t time.Time) (*Holiday, error) {
	return c.client.Next(context.Background(), t)
}

///////////////////////////////////////////////////////////////////////

func (client *Client) IsHoliday(ctx context.Context, t time.Time) (bool, error) {
	h, err := client.Get(ctx, t)

	if err != nil {
		return false, err
	}

	return h != nil, nil
}

func (c *ClientWithoutContext) IsHoliday(t time.Time) (bool, error) {
	return c.client.IsHoliday(context.Background(), t)
}

///////////////////////////////////////////////////////////////////////

func (client *Client) Between(ctx context.Context, from time.Time, to time.Time) ([]*Holiday, error) {
	eventsListCall, err := client.newEventsListCall(ctx)

	if err != nil {
		return nil, err
	}

	from = from.In(JapaneseHolidayLocation)
	to = to.In(JapaneseHolidayLocation)
	min := time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, JapaneseHolidayLocation)
	max := time.Date(to.Year(), to.Month(), to.Day(), 23, 59, 59, 0, JapaneseHolidayLocation)
	events, err := eventsListCall.TimeMin(min.Format(time.RFC3339)).TimeMax(max.Format(time.RFC3339)).Do()

	if err != nil {
		return nil, err
	}

	holidays := []*Holiday{}

	for _, item := range events.Items {
		h, err := newHoliday(item)

		if err != nil {
			return nil, err
		}

		holidays = append(holidays, h)
	}

	return holidays, nil
}

func (c *ClientWithoutContext) Between(from time.Time, to time.Time) ([]*Holiday, error) {
	return c.client.Between(context.Background(), from, to)
}
