package types

import (
	"context"
	"time"
)

type CalendarProvider interface {
	Name() string
	GetEvents(ctx context.Context, date time.Time, calendarId string) ([]*CalendarEvent, error)
}

type CalendarEvent struct {
	Title string
	Start time.Time
	End   time.Time
}

type ParsedEvent struct {
	Title string
	Start time.Time
	End   time.Time

	Emoji           string
	SetDoNotDisturb bool
	SetAway         bool
	Prioritise      bool
}
