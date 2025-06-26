package types

import (
	"context"
	"time"
)

type Client interface {
	EnableDoNotDisturb(ctx context.Context, start time.Time, end time.Time) error
	DisableDoNotDisturb(ctx context.Context) error

	SetPresence(ctx context.Context, isAway bool) error
	ClearPresence(ctx context.Context) error

	SetStatus(ctx context.Context, status string, emoji string, end time.Time) error
	ClearStatus(ctx context.Context) error
}
