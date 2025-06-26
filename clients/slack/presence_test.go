package slack

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Presence(t *testing.T) {
	ctx := context.Background()

	t.Run("should set the presence to auto clearing presence", func(t *testing.T) {
		slack := &testSlackClient{
			ExpectedPresence: presenceAuto,
		}
		client := &Client{slack}

		err := client.ClearPresence(ctx)

		assert.NoError(t, err)
		assert.Equal(t, 1, slack.CallsSetPresence)
	})

	t.Run("should set the presence to auto when isaway is false", func(t *testing.T) {
		slack := &testSlackClient{
			ExpectedPresence: presenceAuto,
		}
		client := &Client{slack}

		err := client.SetPresence(ctx, false)

		assert.NoError(t, err)
		assert.Equal(t, 1, slack.CallsSetPresence)
	})

	t.Run("should set the presence to away when isaway is true", func(t *testing.T) {
		slack := &testSlackClient{
			ExpectedPresence: presenceAway,
		}
		client := &Client{slack}

		err := client.SetPresence(ctx, true)

		assert.NoError(t, err)
		assert.Equal(t, 1, slack.CallsSetPresence)
	})
}
