package slack

import (
	"context"
)

const (
	presenceAuto = "auto"
	presenceAway = "away"
)

func (c *Client) SetPresence(ctx context.Context, isAway bool) error {
	presence := presenceAuto
	if isAway {
		presence = presenceAway
	}

	return c.slack.SetUserPresenceContext(ctx, presence)
}

func (c *Client) ClearPresence(ctx context.Context) error {
	return c.slack.SetUserPresenceContext(ctx, presenceAuto)
}
