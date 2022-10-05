/* This file is auto-generated, manual edits in this file will be overwritten! */
package sentry

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/nightborn-be/blink/blink-demo/config"
	"github.com/samber/lo"
)

func InitialiseSentry(dsn string, environment config.Environment, version *string) error {
	// Default options
	options := sentry.ClientOptions{
		Dsn:         dsn,
		Environment: string(environment),
	}
	// Add version
	if version != nil {
		options.Release = *version
	}
	// Ignore errors
	if len(ignoreErrors) > 0 {
		options.BeforeSend = func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			// Ignore unwanted errors
			if event.Exception != nil && lo.Contains(ignoreErrors, event.Exception[0].Value) {
				return nil
			}
			return event
		}
	}

	// Init sentry
	err := sentry.Init(options)
	if err != nil {
		return fmt.Errorf("sentry.Init: %s", err)
	}

	// Notifies Sentry that the application started
	sentry.CaptureMessage("Application started")
	return nil
}
