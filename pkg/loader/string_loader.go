package loader

import "github.com/SharkEzz/provisiond/pkg/types"

// StringLoader load a deployment configuration from it's content directly.
type StringLoader string

func (s StringLoader) Load() (*types.Deployment, error) {
	return parseYAML([]byte(s))
}
