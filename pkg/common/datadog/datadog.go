package datadog

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

type Tags map[string]string

var (
	mu         sync.RWMutex
	wg         sync.WaitGroup
	prefix     string
	client     *statsd.Client = nil
	globalTags Tags           = Tags{}
)

func InitializeClient() {
	c := NewClient()
	if c != nil {
		mu.Lock()
		client = c
		mu.Unlock()
	}
}

func NewClient() *statsd.Client {
	c, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "stat/datadog").
			Msg("Could not create a statsd client.")
		return nil
	}

	log.Debug().
		Str("service", "stat/datadog").
		Msg("Created statsd client.")

	return c
}

func AddTag(name, value string) {
	globalTags[name] = value
}

func SetPrefix(value string) {
	prefix = value
}

// RunGaugeTask every duration and send value as Gauge
func RunGaugeTask(name string, duration time.Duration, tags Tags, exec func() (float64, error)) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for range time.Tick(duration) {
			value, err := exec()
			if err != nil {
				log.Fatal().
					Str("service", "stat/datadog").
					Str("name", name).
					Err(err).
					Send()
				break
			}

			Gauge(name, value, tags)
		}
	}()
}

func Start() {
	go func() {
		for range time.Tick(time.Minute) {
			c := NewClient()
			if c != nil {
				mu.Lock()
				client = c
				mu.Unlock()
			}
		}
	}()

	wg.Wait()
}
