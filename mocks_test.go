package rich_chat_statuses

import (
	"context"
	"errors"
	"github.com/bnjns/rich-chat-statuses/types"
	"time"
)

type mockClient struct {
	CallsEnableDoNotDisturb  int
	CallsDisableDoNotDisturb int
	CallsSetPresence         int
	CallsClearPresence       int
	CallsSetStatus           int
	CallsClearStatus         int
}

func (c *mockClient) EnableDoNotDisturb(ctx context.Context, start time.Time, end time.Time) error {
	c.CallsEnableDoNotDisturb++
	return nil
}

func (c *mockClient) DisableDoNotDisturb(ctx context.Context) error {
	c.CallsDisableDoNotDisturb++
	return nil
}

func (c *mockClient) SetPresence(ctx context.Context, isAway bool) error {
	c.CallsSetPresence++
	return nil
}

func (c *mockClient) ClearPresence(ctx context.Context) error {
	c.CallsClearPresence++
	return nil
}

func (c *mockClient) SetStatus(ctx context.Context, status string, emoji string, end time.Time) error {
	c.CallsSetStatus++
	return nil
}

func (c *mockClient) ClearStatus(ctx context.Context) error {
	c.CallsClearStatus++
	return nil
}

var errInvalidCalendarId = errors.New("invalid calendar id")

type mockCalendarProvider struct {
	CallsGetEvents     int
	ExpectedCalendarId string
	EventsToReturn     []*types.CalendarEvent
}

func (p *mockCalendarProvider) Name() string {
	return "mock"
}

func (p *mockCalendarProvider) GetEvents(ctx context.Context, date time.Time, calendarId string) ([]*types.CalendarEvent, error) {
	p.CallsGetEvents++

	if p.ExpectedCalendarId != "" && calendarId != p.ExpectedCalendarId {
		return []*types.CalendarEvent{}, errInvalidCalendarId
	}

	return p.EventsToReturn, nil
}
