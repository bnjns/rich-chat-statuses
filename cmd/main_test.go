package main

import (
	"context"
	"github.com/apex/log"
	jsonfmt "github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/logfmt"
	"github.com/aws/smithy-go/ptr"
	rcs "github.com/bnjns/rich-chat-statuses"
	"github.com/bnjns/rich-chat-statuses/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_withLogFormatFromEnv(t *testing.T) {
	ctx := context.Background()

	t.Run("the log format handler should be set to logfmt by default", func(t *testing.T) {
		app := &rcs.App{}
		fn := withLogFormatFromEnv(ctx, &testSecretsClient{})

		fn(app)

		assert.IsType(t, logfmt.Default, app.LogHandler)
	})

	t.Run("the log format handler should be set to json when the LOG_FORMAT env variable is set", func(t *testing.T) {
		t.Setenv(envLogFormat, "json")
		app := &rcs.App{}
		fn := withLogFormatFromEnv(ctx, &testSecretsClient{})

		fn(app)

		assert.IsType(t, jsonfmt.Default, app.LogHandler)
	})
}

func Test_withLogLevelFromEnv(t *testing.T) {
	ctx := context.Background()

	t.Run("the log level should default to info", func(t *testing.T) {
		app := &rcs.App{}
		fn := withLogLevelFromEnv(ctx, &testSecretsClient{})

		fn(app)

		assert.Equal(t, log.InfoLevel, app.LogLevel)
	})

	t.Run("the log level should be set to the contents of the LOG_LEVEL env variable", func(t *testing.T) {
		t.Setenv(envLogLevel, "DEBUG")
		app := &rcs.App{}
		fn := withLogLevelFromEnv(ctx, &testSecretsClient{})

		fn(app)

		assert.Equal(t, log.DebugLevel, app.LogLevel)
	})
}

func Test_withStatusPresetsFromEnv(t *testing.T) {
	ctx := context.Background()

	t.Run("the status presets should not be set by default", func(t *testing.T) {
		app := &rcs.App{}
		fn := withStatusPresetsFromEnv(ctx, &testSecretsClient{})

		fn(app)

		assert.Empty(t, app.StatusPresets)
	})

	t.Run("the status presets should not be set if STATUS_PRESETS contains invalid json", func(t *testing.T) {
		t.Setenv(envStatusPresets, `{"invalid":"structure"}`)
		app := &rcs.App{}
		fn := withStatusPresetsFromEnv(ctx, &testSecretsClient{})

		fn(app)

		assert.Empty(t, app.StatusPresets)
	})

	t.Run("the status presets should be set to the STATUS_PRESETS env variable", func(t *testing.T) {
		t.Setenv(envStatusPresets, `[
	{
		"events": ["first"],
		"emoji": "example",
		"doNotDisturb": true,
		"away": true
	},
	{
		"events": ["second"]
	}
]`)
		app := &rcs.App{}
		fn := withStatusPresetsFromEnv(ctx, &testSecretsClient{})

		fn(app)

		assert.Len(t, app.StatusPresets, 2)
		assert.Equal(t, types.StatusPreset{
			Events:       []string{"first"},
			Emoji:        ptr.String("example"),
			DoNotDisturb: ptr.Bool(true),
			Away:         ptr.Bool(true),
		}, app.StatusPresets[0])
		assert.Equal(t, types.StatusPreset{
			Events:       []string{"second"},
			Emoji:        nil,
			DoNotDisturb: nil,
			Away:         nil,
		}, app.StatusPresets[1])
	})
}
