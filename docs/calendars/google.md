---
title: Google Calendar
---

# Google Calendar

## Requirements

To use this calendar provider, you will need to create a GCP service account:

- In the [GCP console](https://console.cloud.google.com) select the correct project
- Go to `APIs and services` > `Credentials`
- Click `+ Create credentials` and select `Service account`
- Enter a suitable ID (and optional name and description) for the service account and click `Done`
- Select the newly-created service account, then under `Keys` click `Add key` and `Create new key` (use JSON)
- Save this key in a suitable location, depending on how you use the application

You will need to grant this service account read access to the desired calendar:

- Find the calendar at [calendar.google.com](https://calendar.google.com)
- Click the 3 dots next to the calendar and select `Settings and sharing`
- Scroll down to `Shared with` and click `+ Add people and groups`
- Enter the email of the service account created above with the `See all event details` permission and click `Send`

## Usage

!!! tip

    The calendar ID can be retrieved by clicking the 3 dots next to the calendar, selecting `Settings and sharing` and
    scrolling down to the `Integrate calendar` section.

### Programmatically

Add the `github.com/bnjns/rich-chat-statuses/calendars/google` module:

```shell
go get github.com/bnjns/rich-chat-statuses/calendars/google
```

Then create the calendar provider using the credentials created above, and configure the app with that and the calendar
ID using the provided helper functions:

```go
package main

import (
	"context"
	rcs "github.com/bnjns/rich-chat-statuses"
	"github.com/bnjns/rich-chat-statuses/calendars/google"
)

func main() {
    ctx := context.Background()

	// Ideally you should fetch the credentials from an environment variable or secret store
	calendar, err := google.New(ctx, google.WithCredentialsJson(`...`))
	if err != nil {
		panic(err)
    }
	
	app, err := rcs.New(
		rcs.WithCalendarProvider(calendar),
		rcs.WithCalendarId("calendar ID"),
    )
	if err != nil {
		panic(err)
    }
	
	if err := app.Execute(ctx, ...); err != nil {
		panic(err)
    }
}
```

### Standalone binary

The Google Calendar provider is automatically included in the standalone binary; to use it set the `CALENDAR_TYPE`
environment variable to `google`.

You should then provide the calendar ID in the `CALENDAR_ID` environment variable and the service account key in the
`GOOGLE_CREDENTIALS_JSON` environment variable. See [Configuring the app][1] for more details.

[1]: ../usage/standalone.md#configuring-the-app
