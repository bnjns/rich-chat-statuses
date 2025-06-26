package main

import (
	"context"
	"github.com/bnjns/rich-chat-statuses/clients/slack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_selectClient(t *testing.T) {
	ctx := context.Background()
	secrets := &testSecretsClient{}

	t.Run("an error should be returned if no env variables set", func(t *testing.T) {
		client, err := selectClient(ctx, secrets)

		assert.ErrorIs(t, err, errInvalidClient)
		assert.Nil(t, client)
	})

	t.Run("the slack client should be returned if SLACK_TOKEN env variable is set", func(t *testing.T) {
		t.Setenv(envSlackToken, "xoxp-example")

		client, err := selectClient(ctx, secrets)

		assert.NoError(t, err)
		assert.IsType(t, &slack.Client{}, client)
	})
}
