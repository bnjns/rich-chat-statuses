package slack

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"math"
	"time"
)

func (c *Client) EnableDoNotDisturb(ctx context.Context, start time.Time, end time.Time) error {
	numMinutes := time.Until(end).Minutes()
	log.Debugf("Enabling do not disturb (snooze) for %f minutes", numMinutes)

	if _, err := c.slack.SetSnoozeContext(ctx, int(math.Ceil(numMinutes))); err != nil {
		return err
	}

	return nil
}

func (c *Client) DisableDoNotDisturb(ctx context.Context) error {
	if _, err := c.slack.EndSnoozeContext(ctx); err != nil {
		return fmt.Errorf("could not end snooze: %w", err)
	}

	return nil
}
