---
title: Status Presets
---

Status presets provide a way to configure commonly used settings without the need to repeat the flags or emoji in every
event in your calendar. They can also be used to share configurations for different event titles which may mean the same
thing (eg, you can share the config for "holiday" and "annual leave" events).

The application will use the preset where an entry in the list of events matches or is part of [the active event's][2]
title. This check is case-insensitive and order-dependent; if multiple presets match the event, then the first one in
the list will be used.

## Available settings

Each preset can optionally set the emoji, do not disturb (snooze) and presence.

```go
type StatusPreset struct {
    // Events is a list of event titles or partial titles.
	// The preset will be used if the active event equals or contains one of these entries.
    Events []string `json:"events"`
    
	// Emoji is used to set the status' emoji, if specified.
	Emoji *string `json:"emoji"`
    
	// DoNotDisturb is used to set the do not disturb (snooze) setting, if specified.
	DoNotDisturb *bool `json:"doNotDisturb"`
    
	// Away is used to set the presence (away/online) setting, if specified.
	Away *bool `json:"away"`
}
```

!!! tip

    The preset will always be overridden by any flags or emoji specified in the event title.

## Usage

### Programmatically

When used programmatically these are configured by providing a slice of presets to the `WithStatusPresets` option
function.

```go
package main

import(
    rcs "github.com/bnjns/rich-chat-statuses"
    "github.com/bnjns/rich-chat-statuses/types"
)

func main() {
    emoji := "palm_tree"
    doNotDisturb := true
    away := true

    // ...
    app, _ := rcs.New(
        rcs.WithStatusPresets(
            []types.StatusPreset{
                {
                    Events: []string{"holiday", "vacation", "annual leave"},
                    Emoji: &emoji,
                    DoNotDisturb: &doNotDisturb,
                    Away: &away,
                },
            },
        ),
    )
}


```

### Standalone binary

When used in the standalone binary these are configured using the `STATUS_PRESETS` [environment variable][1] as a
JSON-encoded array.

```json
[
    {
        "events": ["holiday", "vacation", "annual leave"],
        "emoji": "palm_tree",
        "doNotDisturb": true,
        "away": true
    }
]
```

[1]: ../usage/standalone.md#configuring-the-app
[2]: ../index.md#prioritising-an-event
