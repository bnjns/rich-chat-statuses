---
title: Slack
---

# Slack

## Supported features

| Feature      | Supported |
|:-------------|:---------:|
| Status text  |     ✅     |
| Status emoji |     ✅     |
| Snooze       |     ✅     |
| Presence     |     ✅     |

## Requirements

At the moment, each user will need to create their own Slack app within your workspace; simply head
to [api.slack.com/apps](https://api.slack.com/apps), and click `Create New App`. You can use the manifest below to
simply the process.

<details>
<summary>Slack manifest</summary>

```yaml
display_information:
  name: Rich Chat Statuses
oauth_config:
  scopes:
    user:
      - dnd:write
      - users.profile:write
      - users:write
settings:
  org_deploy_enabled: false # Set to true to install to your entire org (Enterprise)
  socket_mode_enabled: false
  token_rotation_enabled: false # Recommended you set this to true for security

```
</details>

Once created, you will need to install it to the workspace; go to `Settings > Install App` and press
`Install to Workspace`. This may require administrator approval, depending on your workspace settings.

## Usage

!!! tip

    The authentication token can be retrieved by going to `OAuth & Permissions` and copying the `User OAuth Token`.

### Programmatically

Add the `github.com/bnjns/rich-chat-statuses/clients/slack` module:

```shell
go get github.com/bnjns/rich-chat-statuses/clients/slack
```

Then create the client and run the application on it:

```go
package main

import (
	"context"
	rcs "github.com/bnjns/rich-chat-statuses"
	"github.com/bnjns/rich-chat-statuses/clients/slack"
)

func main() {
	// ... other set up here

	ctx := context.Background()
	app, _ := rcs.New()
	
	// Ideally you should fetch the token from an environment variable or secret store
	slackClient := slack.New("xoxp-...")
	if err := app.Execute(ctx, slackClient); err != nil {
		panic(err)
    }
}

```

### Standalone binary

The Slack client is automatically included in the standalone binary; to use it set the`SLACK_TOKEN` environment variable
to the value of the `User OAuth Token`. See [Configuring the app][1] for more details.

[1]: ../usage/standalone.md#configuring-the-app
