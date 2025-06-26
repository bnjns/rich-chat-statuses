package google

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"github.com/bnjns/rich-chat-statuses/types"
	"google.golang.org/api/calendar/v3"
	"time"
)

type getEventsOpts struct {
	maxResults int64
	orderBy    string
}
type getEventsFunc func(ctx context.Context, date time.Time, calendarId string, opts *getEventsOpts) (*calendar.Events, error)

func (p *CalendarProvider) GetEvents(ctx context.Context, date time.Time, calendarId string) ([]*types.CalendarEvent, error) {
	googleEvents, err := p.getEvents(ctx, date, calendarId, &getEventsOpts{
		orderBy:    "startTime",
		maxResults: 10,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting events from google calendar %s: %w", calendarId, err)
	}

	events := make([]*types.CalendarEvent, 0)
	for _, event := range googleEvents.Items {
		if event.Status != statusCancelled {
			startTime, err := parseDateTime(event.Start)
			if err != nil {
				log.Warnf("error parsing event start %+v: %s", event.Start, err.Error())
			}

			endTime, err := parseDateTime(event.End)
			if err != nil {
				log.Warnf("error parsing event end %+v: %s", event.End, err.Error())
			}

			events = append(events, &types.CalendarEvent{
				Title: event.Summary,
				Start: startTime,
				End:   endTime,
			})
		}
	}

	return events, nil
}

func parseDateTime(dateTime *calendar.EventDateTime) (time.Time, error) {
	if dateTime.DateTime != "" {
		return time.Parse(time.RFC3339, dateTime.DateTime)
	} else if dateTime.Date != "" {
		return time.Parse(time.DateOnly, dateTime.Date)
	} else {
		return time.Time{}, ErrMissingDateTime
	}
}
