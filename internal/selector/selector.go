package selector

import (
	"github.com/bnjns/rich-chat-statuses/types"
	"slices"
)

func Select(events []*types.ParsedEvent) *types.ParsedEvent {
	events = sortByMostRecent(events)

	// Find the first prioritised event
	for _, event := range events {
		if event.Prioritise {
			return event
		}
	}

	// Fall back to the first event
	return events[0]
}

func sortByMostRecent(events []*types.ParsedEvent) []*types.ParsedEvent {
	clonedEvents := slices.Clone(events)

	slices.SortStableFunc(clonedEvents, func(a *types.ParsedEvent, b *types.ParsedEvent) int {
		aStart := a.Start.UnixMilli()
		bStart := b.Start.UnixMilli()

		if aStart == bStart {
			return int(a.End.UnixMilli() - b.End.UnixMilli())
		} else {
			return int(bStart - aStart)
		}
	})

	return clonedEvents
}
