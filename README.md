<h3 align="center">Rich Chat Statuses</h3>

<div align="center">

  [![Status](https://img.shields.io/github/actions/workflow/status/bnjns/rich-chat-statuses/test.yml?branch=main&style=flat-square)](https://github.com/bnjns/rich-chat-statuses/actions/workflows/test.yml) 
  [![GitHub Issues](https://img.shields.io/github/issues/bnjns/rich-chat-statuses?style=flat-square)](https://github.com/bnjns/rich-chat-statuses/issues)
  [![GitHub Pull Requests](https://img.shields.io/github/issues-pr/bnjns/rich-chat-statuses?style=flat-square)](https://github.com/bnjns/rich-chat-statuses/pulls)
  [![License](https://img.shields.io/github/license/bnjns/rich-chat-statuses?style=flat-square)](LICENSE)

</div>

---

<p align="center">
  Allows you to automatically configure rich statuses in your chat client, including the emoji, text, and do not disturb and away settings.
</p>

## 🧐 About

In a world of remote and hybrid working, your status is an important tool to let others know your availability and to
help set expectations on how quickly you might reply to a message. Manually maintaining an accurate status in a
busy working environment can be challenging; with Rich Chat Statuses, you can customise all aspects of your status using
calendar events:

- Status text
- Status emoji
- Do not disturb (snooze) setting
- Presence (away/online) setting

For more information see [the full documentation](https://rich-chat-statuses.bnjns.uk/).

## 🏁 Contributing

### Prerequisites

- [asdf](https://asdf-vm.com/guide/getting-started.html)
- [just](https://github.com/casey/just#installation)

### Installing

Clone the repository:

```sh
git clone git@github.com:bnjns/rich-chat-statuses.git
```

Install the asdf plugins and mkdocs:

```sh
just install
```

### Running the tests

To run all tests:

```sh
just test-all
```

To run specific tests:

- Application tests
  ```sh
  just test-app
  ```
- Standalone binary tests
  ```sh
  just test-cmd
  ```
- Calendar tests
  ```sh
  just test-calendars
  ```
- Client tests
  ```sh
  just test-clients
  ```

### Running manually

To run the standalone binary manually:

```sh
just run-cmd
```

## ⛏️ Built Using

- [apex/log](https://github.com/apex/log)
- [Google Calendar API](https://github.com/googleapis/google-api-go-client)
- [Slack Web API](https://github.com/slack-go/slack)
- [AWS SDK v2](https://github.com/aws/aws-sdk-go-v2)

## ✍️ Authors

- [@bnjns](https://github.com/bnjns)
