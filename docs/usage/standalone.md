---
title: Standalone binary
---

# Standalone binary

If you do not wish to use this programmatically, each release builds a "standalone binary" that you can run instead;
this supports all available calendar providers and chat clients, but can only currently be run on a single calendar and
client at once.

## Configuring the app

You can configure the application using environment variables:

| Environment variable | Required | Default  | Description                                                   |
|:---------------------|:--------:|:---------|:--------------------------------------------------------------|
| `CALENDAR_TYPE`      |    Y     | N/A      | The calendar provider to use (see below).                     |
| `CALENDAR_ID`        |    Y     | N/A      | The ID of the calendar to fetch events from.                  |
| `LOG_FORMAT`         |    N     | `logfmt` | The format to use for the app's logging (`logfmt` or `json`). |
| `LOG_LEVEL`          |    N     | `info`   | The minimum level the app will log.                           |
| `STATUS_PRESETS`     |    N     | N/A      | A JSON-encoded array of any [status presets][3].              |

Calendar provider configuration:

- [Google Calendar](../calendars/google.md#standalone-binary)

Chat client configuration:

- [Slack](../clients/slack.md#standalone-binary)

### Securely storing sensitive values

You may not wish to expose sensitive values (eg, the Slack token) as a raw environment variable; the standalone binary
supports automatically retrieving any environment variable from the following stores:

- AWS Secrets Manager (set the environment variable to the ARN of the secret)
- AWS SSM Parameter Store (set the environment variable to the ARN of the parameter)

It will inherit any permissions from the underlying system, so make sure you have correctly authenticated with the
desired store(s) before running the binary.

## Deployment mechanisms

### AWS Lambda

The binary will automatically run a lambda handler if it detects the AWS Lambda runtime (when the
`AWS_LAMBDA_RUNTIME_API` environment variable is present). You'll need to use [the container runtime][1] to deploy the
binary; you can either use [the pre-built image][2] or build the image yourself.

You will need to ensure that the lambda function has the necessary IAM permissions to access any configuration values
which are stored in Secrets Manager or SSM Parameter Store.

??? example "Example IAM policy"
    ```json5
    {
      "Version": "2012-10-17",
      "Statement": [
        {
          "Sid": "Statement1",
          "Effect": "Allow",
          "Action": ["secretsmanager:GetSecretValue"],
          "Resource": [
            // provide ARNs here
          ]
        },
        {
          "Sid": "Statement2",
          "Effect": "Allow",
          "Action": ["ssm:GetParameter"],
          "Resource": [
            // provide ARNs here
          ]
        }
      ]
    }
    ```

[1]: https://docs.aws.amazon.com/lambda/latest/dg/images-create.html
[2]: #
[3]: ../reference/status-presets.md
