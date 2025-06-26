package rich_chat_statuses

import (
	"errors"
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	"github.com/bnjns/rich-chat-statuses/types"
)

var (
	ErrMissingCalendarProvider = errors.New("no calendar provider set")
	ErrMissingCalendarId       = errors.New("no calendar id set")
)

func New(optFns ...func(*App)) (*App, error) {
	app := &App{
		LogHandler: logfmt.Default,
		LogLevel:   log.InfoLevel,
	}

	for _, optFn := range optFns {
		optFn(app)
	}

	log.SetHandler(app.LogHandler)
	log.SetLevel(app.LogLevel)

	if app.CalenderProvider == nil {
		return nil, ErrMissingCalendarProvider
	}

	if app.CalendarId == "" {
		return nil, ErrMissingCalendarId
	}

	return app, nil
}

func WithCalendarProvider(calendar types.CalendarProvider) func(*App) {
	return func(app *App) {
		app.CalenderProvider = calendar
	}
}

func WithCalendarId(calendarId string) func(*App) {
	return func(app *App) {
		app.CalendarId = calendarId
	}
}

func WithStatusPresets(presets []types.StatusPreset) func(*App) {
	return func(app *App) {
		app.StatusPresets = presets
	}
}

func WithLogHandler(handler log.Handler) func(*App) {
	return func(app *App) {
		app.LogHandler = handler
	}
}

func WithLogLevel(level log.Level) func(*App) {
	return func(app *App) {
		app.LogLevel = level
	}
}
