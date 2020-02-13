package datadog

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type Event struct {
	// Title of the event.  Required.
	Title string
	// Text is the description of the event.  Required.
	Text string
	// AggregationKey groups this event with others of the same key.
	AggregationKey string
	// Priority of the event.  Can be statsd.Low or statsd.Normal.
	Priority statsd.EventPriority
	// SourceTypeName is a source type for the event.
	SourceTypeName string
	// AlertType can be statsd.Info, statsd.Error, statsd.Warning, or statsd.Success.
	// If absent, the default value applied by the dogstatsd server is Info.
	AlertType statsd.EventAlertType
}

// SendEvent represents any record of activity noteworthy for engineers (devs, ops, and security).
func SendEvent(event Event, tags ...Tags) {
	mu.RLock()
	defer mu.RUnlock()
	if client == nil {
		log.Warn().
			Str("service", "stat/datadog").
			Msg("Not initialized")
	}

	host, _ := os.Hostname()
	if err := client.Event(&statsd.Event{
		Title:          event.Title,
		Text:           event.Text,
		Timestamp:      time.Now(),
		Hostname:       host,
		AggregationKey: event.AggregationKey,
		Priority:       event.Priority,
		SourceTypeName: event.SourceTypeName,
		AlertType:      event.AlertType,
		Tags:           mergeTags(tags).StringSlice(),
	}); err != nil {
		log.Error().
			Err(err).
			Msg("Send event failed")
		return
	}

	log.Debug().
		Msg("Sent event")
}
