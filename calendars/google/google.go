package google

import (
	"context"
	"errors"
	"fmt"
	"github.com/bnjns/rich-chat-statuses/types"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"time"
)

const providerName = "google"
const statusCancelled = "cancelled"

// Ensure the provider satisfies the interface
var _ types.CalendarProvider = &CalendarProvider{}

var ErrMissingDateTime = errors.New("event missing date/time")

type CalendarProvider struct {
	getEvents getEventsFunc
}

type Options struct {
	CredentialsJson string
}

func New(ctx context.Context, optFuncs ...func(opt *Options)) (types.CalendarProvider, error) {
	options := &Options{}
	for _, optFunc := range optFuncs {
		optFunc(options)
	}

	jwt, err := google.JWTConfigFromJSON([]byte(options.CredentialsJson), calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("error creating google jwt config: %w", err)
	}

	service, err := calendar.NewService(
		ctx,
		option.WithScopes(calendar.CalendarReadonlyScope),
		option.WithHTTPClient(jwt.Client(ctx)),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating google calendar service: %w", err)
	}

	return &CalendarProvider{
		getEvents: func(ctx context.Context, date time.Time, calendarId string, opts *getEventsOpts) (*calendar.Events, error) {
			return service.Events.List(calendarId).
				Context(ctx).
				MaxResults(opts.maxResults).
				OrderBy(opts.orderBy).
				SingleEvents(true).
				TimeMin(date.Format(time.RFC3339)).
				TimeMax(date.Add(time.Minute).Format(time.RFC3339)).
				Do()
		},
	}, nil
}

func WithCredentialsJson(credentialsJson string) func(opt *Options) {
	return func(opt *Options) {
		opt.CredentialsJson = credentialsJson
	}
}

func (p *CalendarProvider) Name() string {
	return providerName
}
