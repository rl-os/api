package datadog

import "github.com/rs/zerolog/log"

// Gauge Stored as a GAUGE type in DataDog. Each value in the stored time-series is the
// last gauge value submitted for the metric during the StatsD flush period.
func Gauge(name string, value float64, tags ...Tags) {
	mu.RLock()
	defer mu.RUnlock()
	if client == nil {
		log.Warn().
			Str("service", "stat/datadog").
			Msg("Not initialized")
	}

	name = formatName(name)
	if err := client.Gauge(name, value, mergeTags(tags).StringSlice(), 1); err != nil {
		log.Error().
			Err(err).
			Str("name", name).
			Msg("Send gauge failed")
		return
	}

	log.Debug().
		Str("name", name).
		Float64("value", value).
		Msg("Sent gauge")
}
