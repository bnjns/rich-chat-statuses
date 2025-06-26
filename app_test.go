package rich_chat_statuses

import (
	"context"
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	"github.com/bnjns/rich-chat-statuses/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestApp_Execute(t *testing.T) {
	makeApp := func(calendarId string, eventsToReturn []*types.CalendarEvent) *App {
		provider := &mockCalendarProvider{
			ExpectedCalendarId: calendarId,
			EventsToReturn:     eventsToReturn,
		}

		return &App{
			CalenderProvider: provider,
			CalendarId:       calendarId,
			LogHandler:       logfmt.Default,
			LogLevel:         log.DebugLevel,
			StatusPresets:    []types.StatusPreset{},
		}
	}

	ctx := context.Background()
	t.Parallel()
	t.Run("should return error if calendar provider returns error", func(t *testing.T) {
		calendarId := "calendar-123"
		app := &App{
			CalenderProvider: &mockCalendarProvider{
				ExpectedCalendarId: "incorrect",
				EventsToReturn:     []*types.CalendarEvent{},
			},
			CalendarId:    calendarId,
			LogHandler:    logfmt.Default,
			LogLevel:      log.DebugLevel,
			StatusPresets: []types.StatusPreset{},
		}
		client := &mockClient{}

		err := app.Execute(ctx, client)

		assert.ErrorContains(t, err, "error getting the active events")
		assert.Equal(t, 0, client.CallsEnableDoNotDisturb)
		assert.Equal(t, 0, client.CallsDisableDoNotDisturb)
		assert.Equal(t, 0, client.CallsSetPresence)
		assert.Equal(t, 0, client.CallsClearPresence)
		assert.Equal(t, 0, client.CallsSetStatus)
		assert.Equal(t, 0, client.CallsClearStatus)
	})

	t.Run("if no active events the status should be cleared", func(t *testing.T) {
		app := makeApp("calendar-123", []*types.CalendarEvent{})
		client := &mockClient{}

		err := app.Execute(ctx, client)

		assert.NoError(t, err)
		assert.Equal(t, 0, client.CallsEnableDoNotDisturb)
		assert.Equal(t, 1, client.CallsDisableDoNotDisturb)
		assert.Equal(t, 0, client.CallsSetPresence)
		assert.Equal(t, 1, client.CallsClearPresence)
		assert.Equal(t, 0, client.CallsSetStatus)
		assert.Equal(t, 1, client.CallsClearStatus)
	})

	t.Run("status should be set to active event", func(t *testing.T) {
		app := makeApp("calendar-123", []*types.CalendarEvent{
			{Title: "Example", Start: time.Now(), End: time.Now().Add(time.Hour)},
		})
		client := &mockClient{}

		err := app.Execute(ctx, client)

		assert.NoError(t, err)
		assert.Equal(t, 0, client.CallsEnableDoNotDisturb)
		assert.Equal(t, 1, client.CallsDisableDoNotDisturb)
		assert.Equal(t, 1, client.CallsSetPresence)
		assert.Equal(t, 0, client.CallsClearPresence)
		assert.Equal(t, 1, client.CallsSetStatus)
		assert.Equal(t, 0, client.CallsClearStatus)
	})

	t.Run("should set do not disturb if event has the flag", func(t *testing.T) {
		app := makeApp("calendar-123", []*types.CalendarEvent{
			{Title: "[DND] Example", Start: time.Now(), End: time.Now().Add(time.Hour)},
		})
		client := &mockClient{}

		err := app.Execute(ctx, client)

		assert.NoError(t, err)
		assert.Equal(t, 1, client.CallsEnableDoNotDisturb)
	})
}
