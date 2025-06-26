---
title: <Client name>
---

# <Client>

## Supported features

| Feature      | Supported |
|:-------------|:---------:|
| Status text  |    ✅/❌    |
| Status emoji |    ✅/❌    |
| Snooze       |    ✅/❌    |
| Presence     |    ✅/❌    |

## Requirements

<!--
Describe the requirements of the client:
- How to set up/install
- Permissions
-->

## Usage

!!! tip

    Where do you get any relevant config from?

### Programmatically

Add the `github.com/bnjns/rich-chat-statuses/clients/<client>` module:

```shell
go get github.com/bnjns/rich-chat-statuses/clients/<client>
```

Then create the client and run the application on it:

```go
package main

import (
	"context"
	rcs "github.com/bnjns/rich-chat-statuses"
	"github.com/bnjns/rich-chat-statuses/clients/<client>"
)

func main() {
	// ... other set up here

	ctx := context.Background()
	app, _ := rcs.New()
	
	// provide code to create client and run app
}

```

### Standalone binary

<!--
Is the client included in the standalone binary? If so, how do you select and configure it?
-->
