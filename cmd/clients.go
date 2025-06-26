package main

import (
	"context"
	"errors"
	"github.com/bnjns/rich-chat-statuses/clients/slack"
	"github.com/bnjns/rich-chat-statuses/types"
)

var errInvalidClient = errors.New("unable to determine client to set status on")

func selectClient(ctx context.Context, secrets secretFetcher) (types.Client, error) {
	if slackToken, _ := loadEnv(ctx, secrets, envSlackToken); slackToken != "" {
		return slack.New(slackToken), nil
	}

	return nil, errInvalidClient
}
