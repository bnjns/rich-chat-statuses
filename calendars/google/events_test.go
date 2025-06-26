package google

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/calendar/v3"
	"testing"
	"time"
)

func TestCalendarProvider_GetEvents(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("getEvents should be called with the expected parameters", func(t *testing.T) {
		expectedDate := time.Now()
		expectedCalendar := "example"
		expectedMaxResults := int64(10)
		expectedOrderBy := "startTime"

		provider := &CalendarProvider{
			getEvents: func(ctx context.Context, date time.Time, calendarId string, opts *getEventsOpts) (*calendar.Events, error) {
				switch {
				case date != expectedDate:
					return nil, fmt.Errorf("expected date '%s', got '%s'", expectedDate, date)
				case calendarId != expectedCalendar:
					return nil, fmt.Errorf("expected calendarId '%s', got '%s'", expectedCalendar, calendarId)
				case opts.maxResults != expectedMaxResults:
					return nil, fmt.Errorf("expected maxResults '%d', got '%d'", expectedMaxResults, opts.maxResults)
				case opts.orderBy != expectedOrderBy:
					return nil, fmt.Errorf("expected orderBy '%s', got '%s'", expectedOrderBy, opts.orderBy)
				default:
					return &calendar.Events{Items: []*calendar.Event{}}, nil
				}
			},
		}

		_, err := provider.GetEvents(ctx, expectedDate, expectedCalendar)

		assert.NoError(t, err)
	})

	t.Run("an error fetching the events should be returned", func(t *testing.T) {
		expectedErr := errors.New("something went wrong")
		provider := &CalendarProvider{
			getEvents: func(ctx context.Context, date time.Time, calendarId string, opts *getEventsOpts) (*calendar.Events, error) {
				return nil, expectedErr
			},
		}

		events, err := provider.GetEvents(ctx, time.Now(), "example")

		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, events)
	})

	t.Run("no events should be returned as an empty slice", func(t *testing.T) {
		provider := &CalendarProvider{
			getEvents: func(ctx context.Context, date time.Time, calendarId string, opts *getEventsOpts) (*calendar.Events, error) {
				return &calendar.Events{Items: []*calendar.Event{}}, nil
			},
		}

		events, err := provider.GetEvents(ctx, time.Now(), "example")

		assert.NoError(t, err)
		assert.Empty(t, events)
	})

	t.Run("events should be mapped correctly", func(t *testing.T) {
		provider := &CalendarProvider{
			getEvents: func(ctx context.Context, date time.Time, calendarId string, opts *getEventsOpts) (*calendar.Events, error) {
				return &calendar.Events{Items: []*calendar.Event{
					{
						Summary: "An example event",
						Status:  "confirmed",
						Start:   &calendar.EventDateTime{Date: "2025-01-01"},
						End:     &calendar.EventDateTime{Date: "2025-01-02"},
					},
				}}, nil
			},
		}

		events, err := provider.GetEvents(ctx, time.Now(), "example")

		assert.NoError(t, err)
		if assert.Len(t, events, 1) {
			event := events[0]

			assert.Equal(t, "An example event", event.Title)
			assert.Equal(t, time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), event.Start)
			assert.Equal(t, time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC), event.End)
		}
	})

	t.Run("cancelled events should be ignored", func(t *testing.T) {
		provider := &CalendarProvider{
			getEvents: func(ctx context.Context, date time.Time, calendarId string, opts *getEventsOpts) (*calendar.Events, error) {
				return &calendar.Events{Items: []*calendar.Event{
					{
						Summary: "A cancelled event",
						Status:  "cancelled",
						Start:   &calendar.EventDateTime{Date: "2025-01-01"},
						End:     &calendar.EventDateTime{Date: "2025-01-02"},
					},
				}}, nil
			},
		}

		events, err := provider.GetEvents(ctx, time.Now(), "example")

		assert.NoError(t, err)
		assert.Empty(t, events)
	})
}

func TestParseDateTime(t *testing.T) {
	t.Parallel()

	t.Run("an empty datetime should return an error", func(t *testing.T) {
		dateTime := &calendar.EventDateTime{}

		_, err := parseDateTime(dateTime)

		assert.ErrorIs(t, err, ErrMissingDateTime)
	})

	t.Run("a datetime in an invalid format should return an error", func(t *testing.T) {
		dateTime := &calendar.EventDateTime{
			Date: "invalid",
		}

		_, err := parseDateTime(dateTime)

		assert.Error(t, err)
	})

	t.Run("a valid date should be parsed", func(t *testing.T) {
		dateTime := &calendar.EventDateTime{
			Date: "2025-01-02",
		}

		parsedTime, err := parseDateTime(dateTime)

		assert.NoError(t, err)
		assert.Equal(t, time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC), parsedTime)
	})

	t.Run("a valid datetime should be parsed", func(t *testing.T) {
		dateTime := &calendar.EventDateTime{
			DateTime: "2025-01-02T15:04:05Z",
		}

		parsedTime, err := parseDateTime(dateTime)

		assert.NoError(t, err)
		assert.Equal(t, time.Date(2025, 1, 2, 15, 4, 5, 0, time.UTC), parsedTime)
	})
}
