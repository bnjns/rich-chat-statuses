package rich_chat_statuses

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"github.com/bnjns/rich-chat-statuses/internal/parser"
	"github.com/bnjns/rich-chat-statuses/internal/selector"
	"github.com/bnjns/rich-chat-statuses/types"
	"time"
)

type App struct {
	CalenderProvider types.CalendarProvider
	CalendarId       string

	LogHandler log.Handler
	LogLevel   log.Level

	StatusPresets []types.StatusPreset
}

func (a *App) Execute(ctx context.Context, client types.Client) error {
	activeEvents, err := getActiveEvents(ctx, a.CalendarId, a.CalenderProvider)
	if err != nil {
		return fmt.Errorf("error getting the active events: %w", err)
	}

	if len(activeEvents) > 0 {
		log.Debug("Parsing the events")
		parsedEvents := make([]*types.ParsedEvent, len(activeEvents))
		parserConfig := &parser.Config{StatusPresets: a.StatusPresets}
		for i, event := range activeEvents {
			parsedEvents[i] = parser.Parse(parserConfig, event)
		}

		return setStatus(ctx, client, selector.Select(parsedEvents))
	} else {
		log.Info("There are no active events. Clearing any existing status.")
		return clearStatus(ctx, client)
	}
}

func getActiveEvents(ctx context.Context, calendarId string, calenderProvider types.CalendarProvider) ([]*types.CalendarEvent, error) {
	now := time.Now()
	log.Debugf("Finding events that are active at %s", now.Format(time.RFC3339))

	events, err := calenderProvider.GetEvents(ctx, now, calendarId)
	if err != nil {
		return []*types.CalendarEvent{}, fmt.Errorf("failed to get events from provider %s: %w", calenderProvider.Name(), err)
	}

	log.Infof("Found %d active events", len(events))
	return events, nil
}

func setStatus(ctx context.Context, client types.Client, event *types.ParsedEvent) error {
	if event.SetDoNotDisturb {
		if err := client.EnableDoNotDisturb(ctx, event.Start, event.End); err != nil {
			return fmt.Errorf("error enabling do not disturb: %w", err)
		}
		log.Info("Successfully enabled Do Not Disturb")
	} else {
		if err := client.DisableDoNotDisturb(ctx); err != nil {
			return fmt.Errorf("error disabling do not disturb: %w", err)
		}
		log.Info("Successfully disabled Do Not Disturb")
	}

	if err := client.SetPresence(ctx, event.SetAway); err != nil {
		return fmt.Errorf("error setting presence: %w", err)
	}
	log.Infof("Successfully configured presence to %t", event.SetAway)

	if err := client.SetStatus(ctx, event.Title, event.Emoji, event.End.UTC()); err != nil {
		return fmt.Errorf("error setting status: %w", err)
	}
	log.Info("Successfully set custom status")

	return nil
}

func clearStatus(ctx context.Context, client types.Client) error {
	log.Debug("Disabling do not disturb")
	if err := client.DisableDoNotDisturb(ctx); err != nil {
		return fmt.Errorf("error disabling do not disturb: %w", err)
	}
	log.Info("Successfully disabled Do Not Disturb")

	log.Debug("Clearing any custom status")
	if err := client.ClearStatus(ctx); err != nil {
		return fmt.Errorf("error clearing status: %w", err)
	}

	log.Debug("Clearing any custom presence")
	if err := client.ClearPresence(ctx); err != nil {
		return fmt.Errorf("error clearing presence: %w", err)
	}

	return nil
}
