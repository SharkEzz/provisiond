package loader

import (
	"github.com/SharkEzz/provisiond/pkg/deployment"
	"github.com/SharkEzz/provisiond/pkg/utils"
)

// StringLoader load a deployment configuration from it's content directly.
type StringLoader string

func (s StringLoader) Load() (*deployment.Deployment, error) {
	return utils.ParseYAML([]byte(s))
}
