---
title: Programmatically
---

The preferred way to use this SDK is programmatically; this consists of 4 steps:

1. Configure the calendar provider
2. Configure the chat client
3. Configure the application, with the calendar provider, calendar ID, and other settings
4. Run the application

```go
package main

import (
	"context"
	rcs "github.com/bnjns/rich-chat-statuses"
	"github.com/bnjns/rich-chat-statuses/calendars/google"
	"github.com/bnjns/rich-chat-statuses/clients/slack"
	"os"
)

func main() {
	ctx := context.Background()
	
	// 1. Create the calendar provider, here we use Google
	credentialsJson := os.Getenv("GOOGLE_CREDENTIALS_JSON")
	calendar, err := google.New(ctx, google.WithCredentialsJson(credentialsJson))
	if err != nil {
		panic(err)
    }
	
	// 2. Configure the chat client, here we use Slack
	slackClient := slack.New(os.Getenv("SLACK_TOKEN"))
	
	// 3. Create the application
	app, err := rcs.New(
		rcs.WithCalendarProvider(calendar),
		rcs.WithCalendarId(os.Getenv("CALENDAR_ID")),
    )
	
	// 4. Run the application
	if err := app.Execute(ctx, slackClient); err != nil {
		panic(err)
    }
}
```

## Available settings

- `WithCalendarProvider`: configures [the calendar provider](../calendars/index.md) used to fetch the active events
- `WithCalendarId`: configures the ID of the calendar the events are fetched from
- `WithLogHandler`: configures [the log handler](https://github.com/apex/log#handlers) for the app, clients and calendar provider (defaults to `logfmt`)
- `WithLogLevel`: configures the log level for the app, clients and calendar provider (defaults to `INFO`)
- `WithStatusPresets`: configures the [status presets](../reference/status-presets.md)

## Multiple clients

You can run the same application on multiple chat clients if you wish. At the moment, this needs to be manually executed
sequentially.

```go
func main() {
	// Steps 1 and 3 omitted for brevity
	
	// Create clients
	firstClient := slack.New(os.Getenv("SLACK_TOKEN_FIRST"))
	secondClient := slack.New(os.Getenv("SLACK_TOKEN_SECOND"))
	
	// Run the application on each client in turn
	if err := app.Execute(ctx, firstClient); err != nil {
		panic(err)
    }
    if err := app.Execute(ctx, secondClient); err != nil {
        panic(err)
    }
}
```
