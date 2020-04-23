package bancho

import (
	"github.com/deissh/osu-lazer/ayako/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {

	tests := []struct {
		name string
		cfg  *config.Config
	}{
		{
			"create client",
			&config.Config{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Init(tt.cfg)
			assert.NotNil(t, got)
		})
	}
}
