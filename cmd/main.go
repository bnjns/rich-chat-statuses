package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/apex/log"
	jsonfmt "github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/logfmt"
	"github.com/aws/aws-lambda-go/lambda"
	rcs "github.com/bnjns/rich-chat-statuses"
	"github.com/bnjns/rich-chat-statuses/types"
	"os"
	"strings"
)

const (
	envCalendarType = "CALENDAR_TYPE"
	envCalendarId   = "CALENDAR_ID"
	envLogFormat    = "LOG_FORMAT"
	envLogLevel     = "LOG_LEVEL"

	envStatusPresets = "STATUS_PRESETS"

	envGoogleCredentialsJson = "GOOGLE_CREDENTIALS_JSON"
	envSlackToken            = "SLACK_TOKEN"

	envIsAwsLambda = "AWS_LAMBDA_RUNTIME_API"
)

var (
	version string = "dev-master"
	commit  string = "unknown"
)

const versionStringTemplate = `
Rich Chat Statuses standalone binary, version %s (commit %s)

Copyright 2025 Ben Jones

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
`

func main() {
	versionFlag := flag.Bool("version", false, "Displays the version of the CLI")
	flag.Parse()

	if *versionFlag {
		fmt.Printf(strings.TrimSpace(versionStringTemplate)+"\n", version, commit)
		os.Exit(0)
	}

	// Run as AWS Lambda handler
	if _, isAwsLambda := os.LookupEnv(envIsAwsLambda); isAwsLambda {
		log.Debug("Running AWS Lambda handler")
		lambda.Start(func(ctx context.Context) error {
			err := execute(ctx)
			if err != nil {
				log.WithError(err).Error("Error running lambda")
			}
			return err
		})
		return
	}

	// Fall back to running as normal
	log.Debug("Running application")
	if err := execute(context.Background()); err != nil {
		log.WithError(err).Error("Error running application")
		os.Exit(1)
	}
}

func execute(ctx context.Context) error {
	clients, err := newSecretClients(ctx)
	if err != nil {
		return fmt.Errorf("error building aws clients: %w", err)
	}

	calendarId, err := loadEnv(ctx, clients, envCalendarId)
	if err != nil {
		return err
	}

	client, err := selectClient(ctx, clients)
	if err != nil {
		return err
	}

	calendarProvider, err := selectCalendarProvider(ctx, clients)
	if err != nil {
		return err
	}

	app, err := rcs.New(
		rcs.WithCalendarProvider(calendarProvider),
		rcs.WithCalendarId(calendarId),
		withLogFormatFromEnv(ctx, clients),
		withLogLevelFromEnv(ctx, clients),
		withStatusPresetsFromEnv(ctx, clients),
	)
	if err != nil {
		return err
	}

	return app.Execute(ctx, client)
}

func withLogFormatFromEnv(ctx context.Context, secrets secretFetcher) func(*rcs.App) {
	determineLogHandler := func() log.Handler {
		logFormat, _ := loadEnv(ctx, secrets, envLogFormat)
		switch strings.ToLower(logFormat) {
		case "json":
			return jsonfmt.Default
		default:
			return logfmt.Default
		}
	}

	return rcs.WithLogHandler(determineLogHandler())
}

func withLogLevelFromEnv(ctx context.Context, secrets secretFetcher) func(*rcs.App) {
	logLevelStr, _ := loadEnv(ctx, secrets, envLogLevel)
	if logLevelStr == "" {
		logLevelStr = "info"
	}

	return rcs.WithLogLevel(log.MustParseLevel(strings.ToLower(logLevelStr)))
}

func withStatusPresetsFromEnv(ctx context.Context, secrets secretFetcher) func(*rcs.App) {
	return func(app *rcs.App) {
		statusPresetsStr, _ := loadEnv(ctx, secrets, envStatusPresets)
		if statusPresetsStr != "" {
			var statusPresets []types.StatusPreset
			err := json.Unmarshal([]byte(statusPresetsStr), &statusPresets)
			if err != nil {
				log.Warnf("error parsing status presets from %s: %s", envStatusPresets, err.Error())
			} else {
				app.StatusPresets = statusPresets
			}
		}
	}
}
