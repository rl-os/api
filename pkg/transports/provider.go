package transports

import (
	"github.com/google/wire"
	"github.com/rl-os/api/pkg/transports/http"
)

var ProviderSet = wire.NewSet(http.ProviderSet)
