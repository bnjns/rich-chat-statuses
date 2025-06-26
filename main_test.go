package rich_chat_statuses

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	"github.com/aws/smithy-go/ptr"
	"github.com/bnjns/rich-chat-statuses/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithCalendarProvider(t *testing.T) {
	t.Run("should set the calendar provider", func(t *testing.T) {
		app := &App{}
		fn := WithCalendarProvider(&mockCalendarProvider{})

		fn(app)

		assert.IsType(t, &mockCalendarProvider{}, app.CalenderProvider)
	})
}

func TestWithCalendarId(t *testing.T) {
	t.Run("should set the calendar id", func(t *testing.T) {
		calendarId := "calendar-123"
		app := &App{}
		fn := WithCalendarId(calendarId)

		fn(app)

		assert.Equal(t, calendarId, app.CalendarId)
	})
}

func TestWithStatusPresets(t *testing.T) {
	t.Run("should set the status presets", func(t *testing.T) {
		statusPresets := []types.StatusPreset{
			{Events: []string{"first"}, Emoji: ptr.String("emoji"), DoNotDisturb: ptr.Bool(true), Away: ptr.Bool(true)},
		}
		app := &App{}
		fn := WithStatusPresets(statusPresets)

		fn(app)

		if assert.Len(t, app.StatusPresets, 1) {
			assert.Equal(t, "emoji", *app.StatusPresets[0].Emoji)
			assert.Equal(t, true, *app.StatusPresets[0].DoNotDisturb)
			assert.Equal(t, true, *app.StatusPresets[0].Away)
		}
	})
}

func TestWithLogHandler(t *testing.T) {
	t.Run("should set the log handler", func(t *testing.T) {
		app := &App{}
		fn := WithLogHandler(logfmt.Default)

		fn(app)

		assert.IsType(t, logfmt.Default, app.LogHandler)
	})
}

func TestWithLogLevel(t *testing.T) {
	t.Run("should set the log level", func(t *testing.T) {
		app := &App{}
		fn := WithLogLevel(log.DebugLevel)

		fn(app)

		assert.Equal(t, log.DebugLevel, app.LogLevel)
	})
}
