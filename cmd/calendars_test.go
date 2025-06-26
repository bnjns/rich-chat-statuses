package main

import (
	"context"
	"github.com/bnjns/rich-chat-statuses/calendars/google"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_selectCalendarProvider(t *testing.T) {
	ctx := context.Background()
	secrets := &testSecretsClient{}

	t.Run("an error should be returned if no CALENDAR_TYPE set", func(t *testing.T) {
		calendar, err := selectCalendarProvider(ctx, secrets)

		assert.ErrorIs(t, err, errMissingCalendarType)
		assert.Nil(t, calendar)
	})

	t.Run("an error should be invalid CALENDAR_TYPE set", func(t *testing.T) {
		t.Setenv(envCalendarType, "invalid")

		calendar, err := selectCalendarProvider(ctx, secrets)

		assert.ErrorContains(t, err, "invalid calendar type")
		assert.Nil(t, calendar)
	})

	t.Run("google calendar should be selected when CALENDAR_TYPE is set to google", func(t *testing.T) {
		t.Setenv(envCalendarType, "google")
		t.Setenv("GOOGLE_CREDENTIALS_JSON", `{"type": "service_account"}`)

		calendar, err := selectCalendarProvider(ctx, secrets)

		assert.NoError(t, err)
		assert.IsType(t, &google.CalendarProvider{}, calendar)
	})
}
