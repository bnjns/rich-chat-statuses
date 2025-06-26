package parser

import (
	"github.com/aws/smithy-go/ptr"
	"github.com/bnjns/rich-chat-statuses/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	t.Parallel()

	t.Run("should parse flags", func(t *testing.T) {
		event := &types.CalendarEvent{
			Title: "![DND] [AWAY] Event name",
			Start: time.Now(),
			End:   time.Now().Add(time.Hour),
		}

		parsedEvent := Parse(&Config{}, event)

		assert.Equal(t, "Event name", parsedEvent.Title)
		assert.True(t, parsedEvent.SetAway)
		assert.True(t, parsedEvent.SetDoNotDisturb)
		assert.True(t, parsedEvent.Prioritise)
	})

	t.Run("a status preset should be overridden by the event title", func(t *testing.T) {
		cfg := &Config{
			StatusPresets: []types.StatusPreset{
				{
					Events:       []string{"Event name"},
					Emoji:        ptr.String("example"),
					Away:         ptr.Bool(false),
					DoNotDisturb: ptr.Bool(false),
				},
			},
		}
		event := &types.CalendarEvent{
			Title: "![DND] [AWAY] :emoji: Event name",
			Start: time.Now(),
			End:   time.Now().Add(time.Hour),
		}

		parsedEvent := Parse(cfg, event)

		assert.Equal(t, "Event name", parsedEvent.Title)
		assert.Equal(t, "emoji", parsedEvent.Emoji)
		assert.True(t, parsedEvent.SetAway)
		assert.True(t, parsedEvent.SetDoNotDisturb)
		assert.True(t, parsedEvent.Prioritise)
	})
}

func TestDetectDoNotDisturb(t *testing.T) {
	t.Parallel()

	t.Run("an event without the [DND] flag should not have SetDoNotDisturb set", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "Event name",
		}

		detectDoNotDisturb(event)

		assert.False(t, event.SetDoNotDisturb)
		assert.Equal(t, "Event name", event.Title)
	})

	t.Run("an event with the [DND] flag at the start should have SetDoNotDisturb set", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "[DND] Event name",
		}

		detectDoNotDisturb(event)

		assert.True(t, event.SetDoNotDisturb)
		assert.Equal(t, "Event name", event.Title)
	})

	t.Run("an event with the [DND] flag anywhere in the title should have SetDoNotDisturb set", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "Event [DND] name",
		}

		detectDoNotDisturb(event)

		assert.True(t, event.SetDoNotDisturb)
		assert.Equal(t, "Event name", event.Title)
	})
}

func TestDetectAway(t *testing.T) {
	t.Parallel()

	t.Run("an event without the [AWAY] flag should not have SetAway set", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "Event name",
		}

		detectAway(event)

		assert.False(t, event.SetAway)
		assert.Equal(t, "Event name", event.Title)
	})

	t.Run("an event with the [AWAY] flag at the start should have SetAway set", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "[AWAY] Event name",
		}

		detectAway(event)

		assert.True(t, event.SetAway)
		assert.Equal(t, "Event name", event.Title)
	})

	t.Run("an event with the [AWAY] flag anywhere in the title should have SetAway set", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "Event [AWAY] name",
		}

		detectAway(event)

		assert.True(t, event.SetAway)
		assert.Equal(t, "Event name", event.Title)
	})
}

func TestDetectEmoji(t *testing.T) {
	t.Parallel()

	t.Run("an event without an emoji should have the default set", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "Event name",
		}

		detectEmoji(event)

		assert.Equal(t, emojiDefault, event.Emoji)
	})

	t.Run("an event without an emoji but with DND should have that emoji set", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title:           "Event name",
			SetDoNotDisturb: true,
		}

		detectEmoji(event)

		assert.Equal(t, emojiDoNotDisturb, event.Emoji)
	})

	t.Run("an event with an emoji in the title should have that emoji set without the colons", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: ":test: Event name",
		}

		detectEmoji(event)

		assert.Equal(t, "test", event.Emoji)
		assert.Equal(t, "Event name", event.Title)
	})

	t.Run("if the event has multiple emojis, the first one should be used", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: ":test: Event :test2: name",
		}

		detectEmoji(event)

		assert.Equal(t, "test", event.Emoji)
		assert.Equal(t, "Event name", event.Title)
	})
}

func TestDetectPrioritisation(t *testing.T) {
	t.Parallel()

	t.Run("an event with the ! prefix should be prioritised", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "!Prioritised event",
		}

		detectPrioritisation(event)

		assert.True(t, event.Prioritise)
		assert.Equal(t, "Prioritised event", event.Title)
	})

	t.Run("an event with ! in the title should not be prioritised", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "Super awesome event!",
		}

		detectPrioritisation(event)

		assert.False(t, event.Prioritise)
		assert.Equal(t, "Super awesome event!", event.Title)
	})
}

func TestDetectPreset(t *testing.T) {
	t.Parallel()

	t.Run("no presets should not be matched", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "Event name",
		}
		presets := make([]types.StatusPreset, 0)

		detected := detectPreset(presets, event)

		assert.False(t, detected)
		assert.Equal(t, "Event name", event.Title)
	})

	t.Run("no matching presets should not be matched", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "Event name",
		}
		presets := []types.StatusPreset{
			{
				Events:       []string{"test"},
				Emoji:        ptr.String("emoji"),
				Away:         ptr.Bool(true),
				DoNotDisturb: ptr.Bool(true),
			},
		}

		detected := detectPreset(presets, event)

		assert.False(t, detected)
		assert.Equal(t, "Event name", event.Title)
	})

	t.Run("a matching preset should be used", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title: "Test event",
		}
		presets := []types.StatusPreset{
			{
				Events:       []string{"test"},
				Emoji:        ptr.String("emoji"),
				Away:         ptr.Bool(true),
				DoNotDisturb: ptr.Bool(true),
			},
		}

		detected := detectPreset(presets, event)

		assert.True(t, detected)
		assert.Equal(t, "Test event", event.Title)
		assert.Equal(t, "emoji", event.Emoji)
		assert.True(t, event.SetAway)
		assert.True(t, event.SetDoNotDisturb)
		assert.False(t, event.Prioritise)
	})

	t.Run("a matching mapping should not set properties which are nil", func(t *testing.T) {
		event := &types.ParsedEvent{
			Title:           "Test event",
			Emoji:           "calendar",
			SetAway:         true,
			SetDoNotDisturb: true,
		}
		presets := []types.StatusPreset{
			{
				Events: []string{"test"},
			},
		}

		detected := detectPreset(presets, event)

		assert.True(t, detected)
		assert.Equal(t, "Test event", event.Title)
		assert.Equal(t, "calendar", event.Emoji)
		assert.True(t, event.SetAway)
		assert.True(t, event.SetDoNotDisturb)
	})
}
