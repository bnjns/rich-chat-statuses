---
title: Home
---

# Rich Chat Statuses

In a world of remote and hybrid working, your status is an important tool to let others know your availability and to
help set expectations on how quickly you might reply to a message. However, manually maintaining an accurate status in a
busy working environment can be challenging. While there are existing integrations for syncing your status to a
calendar, the information they provide in the chat client is quite limited and often just show that you're in a meeting.

With Rich Chat Statuses, you can customise all aspects of your status[^1] using calendar events:

- Status text
- Status emoji
- Do not disturb (snooze) setting
- Presence (away/online) setting

## How it works

Rich Chat Statuses works by reading events from your calendar and parsing the event title (some calendars call this the
summary):

- The status emoji is set by specifying the emoji name in the event title surrounded by `:` (eg, `:no_entry:`). If no
  emoji is found, this defaults to `calendar`.
- Use the `[DND]` prefix to enable Do Not Disturb (snooze). This will also automatically set the emoji to `:no_entry:`,
  unless another emoji is explicitly provided.
- Use the `[AWAY]` prefix to set the presence to `away`, otherwise the presence is set to `auto` which uses your
  activity to mark you as online/away.
- The event title is used to set the status text, with any parsed info (eg, emoji, flags) removed.

As an example, an event title of `[DND] [AWAY] :no_entry: I'm not here` would result in:

- Status: `I'm not here`
- Emoji: `:no_entry:`
- Do not disturb: `Enabled`
- Presence: `Away`

### Prioritising an event

The active event is determined by finding all events that are currently occurring and selecting the event which started
last and then ends first.

If you wish to override the selected event prefix the summary with `!`.

[^1]: What you can configure depends on your chat client
