package setup

import (
	"context"
)

func New(opts Options) core.Server {
	// internal setup already accepted functional-options; we wrap only extras.
	return core.New(core.WithExtraPlugins(opts.ExtraPlugins))
}
