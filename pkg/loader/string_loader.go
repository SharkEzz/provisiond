package loader

import "github.com/SharkEzz/provisiond/pkg/deployment"

// StringLoader load a deployment configuration from it's content directly.
type StringLoader string

func (s StringLoader) Load() (*deployment.Deployment, error) {
	return parseYAML([]byte(s))
}
