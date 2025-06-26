package slack

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClient_Status(t *testing.T) {
	ctx := context.Background()

	t.Run("clearing the status should call the expected method", func(t *testing.T) {
		slack := &testSlackClient{}
		client := &Client{slack}

		err := client.ClearStatus(ctx)

		assert.NoError(t, err)
		assert.Equal(t, 1, slack.CallsUnsetStatus)
	})

	t.Run("providing the parameters should set the user's status", func(t *testing.T) {
		slack := &testSlackClient{
			ExpectedStatusText:       "The status",
			ExpectedStatusEmoji:      ":calendar:",
			ExpectedStatusExpiration: 1704067200,
		}
		client := &Client{slack}

		err := client.SetStatus(ctx, "The status", "calendar", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))

		assert.NoError(t, err)
		assert.Equal(t, 1, slack.CallsSetStatus)
	})
}
