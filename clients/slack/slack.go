package slack

import (
	"context"
	"github.com/bnjns/rich-chat-statuses/types"
	"github.com/slack-go/slack"
)

type slackClient interface {
	SetSnoozeContext(ctx context.Context, minutes int) (*slack.DNDStatus, error)
	EndSnoozeContext(ctx context.Context) (*slack.DNDStatus, error)
	SetUserPresenceContext(ctx context.Context, presence string) error
	SetUserCustomStatusContext(ctx context.Context, statusText, statusEmoji string, statusExpiration int64) error
	UnsetUserCustomStatusContext(ctx context.Context) error
}

type Client struct {
	slack slackClient
}

// Ensure the client satisfies the interface
var _ types.Client = &Client{}

func New(token string) types.Client {
	return &Client{
		slack: slack.New(token),
	}
}
