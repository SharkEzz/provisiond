package plugin

import (
	"context"

	"github.com/SharkEzz/provisiond/pkg/plugin/internal"
)

// The Plugin interface define one method,
// which execute the content of the associated content in the deployment file.
type Plugin interface {
	Execute(ctx context.Context, data any) error
}

// The Plugins map contain all the registered plugins,
// whoses names must be the same as the ones used in the deployment file.
var Plugins = map[string]Plugin{
	"shell": &internal.Shell{},
}
