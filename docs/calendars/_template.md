---
title: <Calendar provider>
---

# <Calendar provider>

## Requirements

<!--
Describe the requirements of the calendar provider:
- How to set up/install
- Permissions
-->

## Usage

!!! tip

    Where do you get any relevant config from?

### Programmatically

Add the `github.com/bnjns/rich-chat-statuses/calendars/<calendar>` module:

```shell
go get github.com/bnjns/rich-chat-statuses/calendars/<calendar>
```

Then create the calendar provider using the credentials created above, and configure the app with that and the calendar
ID using the provided helper functions:

```go
package main

import (
	"context"
	rcs "github.com/bnjns/rich-chat-statuses"
	"github.com/bnjns/rich-chat-statuses/calendars/<calendar>"
)

func main() {
    ctx := context.Background()

	// provide code to create and configure the calendar provider
	
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

<!--
Is the calendar provider included in the standalone binary? If so, how do you select and configure it?
-->
