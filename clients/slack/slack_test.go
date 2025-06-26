package slack

import (
	"context"
	"fmt"
	"github.com/slack-go/slack"
)

var _ slackClient = &testSlackClient{}

type testSlackClient struct {
	CallsSetSnooze   int
	CallsEndSnooze   int
	CallsSetPresence int
	CallsSetStatus   int
	CallsUnsetStatus int

	ExpectedSnoozeMins       int
	ExpectedPresence         string
	ExpectedStatusText       string
	ExpectedStatusEmoji      string
	ExpectedStatusExpiration int64
}

func (c *testSlackClient) SetSnoozeContext(ctx context.Context, minutes int) (*slack.DNDStatus, error) {
	c.CallsSetSnooze++

	if c.ExpectedSnoozeMins != minutes {
		return nil, fmt.Errorf("incorrect value %d for parameter minutes, expected %d", minutes, c.ExpectedSnoozeMins)
	}

	return nil, nil
}

func (c *testSlackClient) EndSnoozeContext(ctx context.Context) (*slack.DNDStatus, error) {
	c.CallsEndSnooze++
	return nil, nil
}

func (c *testSlackClient) SetUserPresenceContext(ctx context.Context, presence string) error {
	c.CallsSetPresence++

	if presence != c.ExpectedPresence {
		return fmt.Errorf("incorrect value %s for parameter presence, expected %s", presence, c.ExpectedPresence)
	}

	return nil
}

func (c *testSlackClient) SetUserCustomStatusContext(ctx context.Context, statusText, statusEmoji string, statusExpiration int64) error {
	c.CallsSetStatus++

	if statusText != c.ExpectedStatusText {
		return fmt.Errorf("incorrect value %s for parameter statusText, expected %s", statusText, c.ExpectedStatusText)
	}

	if statusEmoji != c.ExpectedStatusEmoji {
		return fmt.Errorf("incorrect value %s for parameter statusEmoji, expected %s", statusEmoji, c.ExpectedStatusEmoji)
	}

	if statusExpiration != c.ExpectedStatusExpiration {
		return fmt.Errorf("incorrect value %d for parameter statusExpiration, expected %d", statusExpiration, c.ExpectedStatusExpiration)
	}

	return nil
}

func (c *testSlackClient) UnsetUserCustomStatusContext(ctx context.Context) error {
	c.CallsUnsetStatus++
	return nil
}
