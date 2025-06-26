package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/bnjns/rich-chat-statuses/calendars/google"
	"github.com/bnjns/rich-chat-statuses/types"
)

var errMissingCalendarType = errors.New("CALENDAR_TYPE env variable not set")

func selectCalendarProvider(ctx context.Context, secrets secretFetcher) (types.CalendarProvider, error) {
	calendarType, err := loadEnv(ctx, secrets, envCalendarType)
	if err != nil {
		return nil, errMissingCalendarType
	}

	switch calendarType {
	case "google":
		googleCredentialsJson, err := loadEnv(ctx, secrets, envGoogleCredentialsJson)
		if err != nil {
			return nil, fmt.Errorf("error creating google calendar provider: %w", err)
		}

		return google.New(ctx, google.WithCredentialsJson(googleCredentialsJson))
	default:
		return nil, fmt.Errorf("invalid calendar type: %s", calendarType)
	}
}
