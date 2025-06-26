package slack

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"time"
)

func (c *Client) SetStatus(ctx context.Context, status string, emoji string, end time.Time) error {
	log.Debugf("Configuring user status to '%s' with emoji '%s' expiring at '%v'", status, emoji, end)
	return c.slack.SetUserCustomStatusContext(ctx, status, fmt.Sprintf(":%s:", emoji), end.Unix())
}

func (c *Client) ClearStatus(ctx context.Context) error {
	return c.slack.UnsetUserCustomStatusContext(ctx)
}
