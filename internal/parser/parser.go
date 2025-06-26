package parser

import (
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/bnjns/rich-chat-statuses/types"
	"regexp"
	"strings"
)

const (
	flagDoNotDisturb = "DND"
	flagAway         = "AWAY"

	emojiDefault      = "calendar"
	emojiDoNotDisturb = "no_entry"
)

var emojiRegex = regexp.MustCompile(`(?i):([a-z0-9_\\-]+):`)
var multipleSpaceRegex = regexp.MustCompile(`\s{2,}`)

type Config struct {
	StatusPresets []types.StatusPreset
}

func Parse(cfg *Config, event *types.CalendarEvent) *types.ParsedEvent {
	parsedEvent := &types.ParsedEvent{
		Title: event.Title,
		Start: event.Start,
		End:   event.End,
	}

	detectPreset(cfg.StatusPresets, parsedEvent)
	detectDoNotDisturb(parsedEvent)
	detectAway(parsedEvent)
	detectEmoji(parsedEvent)
	detectPrioritisation(parsedEvent)

	log.Debugf("CalendarEvent title has been changed to '%s' from '%s'", parsedEvent.Title, event.Title)

	return parsedEvent
}

func cleanSpaces(str string) string {
	return strings.TrimSpace(multipleSpaceRegex.ReplaceAllString(str, " "))
}

func detectFlag(event *types.ParsedEvent, flag string) bool {
	flagText := fmt.Sprintf("[%s]", flag)

	if strings.Contains(event.Title, flagText) {
		log.Debugf("Found flag %s in event title", flag)
		event.Title = cleanSpaces(strings.ReplaceAll(event.Title, flagText, ""))
		return true
	}

	return false
}

func detectDoNotDisturb(event *types.ParsedEvent) {
	event.SetDoNotDisturb = detectFlag(event, flagDoNotDisturb)
}

func detectAway(event *types.ParsedEvent) {
	event.SetAway = detectFlag(event, flagAway)
}

func detectEmoji(event *types.ParsedEvent) {
	matches := emojiRegex.FindAllStringSubmatch(event.Title, -1)

	// The emoji is provided explicitly in the title
	if len(matches) > 0 {
		log.Debugf("Found %d matching emojies in title: %v", len(matches), matches)

		selectedEmoji := matches[0][1]
		event.Emoji = selectedEmoji
		log.Debugf("Selected emoji: %s", selectedEmoji)

		// Clean all matches
		for _, match := range matches {
			event.Title = cleanSpaces(strings.ReplaceAll(event.Title, match[0], ""))
		}
		return
	}

	// If the event is marked as DND, set the emoji to the default DND emoji
	if event.SetDoNotDisturb {
		log.Debugf("CalendarEvent is marked as do not disturb - using emoji %s", emojiDoNotDisturb)
		event.Emoji = emojiDoNotDisturb
		return
	}

	event.Emoji = emojiDefault
}

func detectPrioritisation(event *types.ParsedEvent) {
	event.Prioritise = strings.HasPrefix(event.Title, "!")

	if event.Prioritise {
		event.Title = cleanSpaces(strings.TrimPrefix(event.Title, "!"))
	}
}

func detectPreset(presets []types.StatusPreset, event *types.ParsedEvent) bool {
	for _, preset := range presets {
		for _, presetEvent := range preset.Events { // TODO: rename `presetEvent`
			if strings.Contains(strings.ToLower(event.Title), strings.ToLower(presetEvent)) {
				mappingJson, _ := json.Marshal(preset)

				log.Debugf("Found text '%s' for status preset in event title: %s", presetEvent, mappingJson)

				if preset.Emoji != nil && *preset.Emoji != "" {
					event.Emoji = *preset.Emoji
				}

				if preset.Away != nil {
					event.SetAway = *preset.Away
				}

				if preset.DoNotDisturb != nil {
					event.SetDoNotDisturb = *preset.DoNotDisturb
				}

				return true
			}
		}
	}

	return false
}
