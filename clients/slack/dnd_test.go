package slack

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClient_DoNotDisturb(t *testing.T) {
	ctx := context.Background()

	t.Run("should set snooze with correct duration", func(t *testing.T) {
		slack := &testSlackClient{
			ExpectedSnoozeMins: 30,
		}
		client := &Client{slack}

		start := time.Now().Add(-30 * time.Minute)
		end := time.Now().Add(30 * time.Minute)

		err := client.EnableDoNotDisturb(ctx, start, end)

		assert.NoError(t, err)
		assert.Equal(t, 1, slack.CallsSetSnooze)
	})

	t.Run("should end snooze when enable is false", func(t *testing.T) {
		slack := &testSlackClient{}
		client := &Client{slack}

		err := client.DisableDoNotDisturb(ctx)

		assert.NoError(t, err)
		assert.Equal(t, 1, slack.CallsEndSnooze)
	})
}
