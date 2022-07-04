package loader

import (
	"os"

	"github.com/SharkEzz/provisiond/pkg/deployment"
	"github.com/SharkEzz/provisiond/pkg/utils"
)

// FileLoader load a deployment configuration from a file.
type FileLoader string

func (f FileLoader) Load() (*deployment.Deployment, error) {
	content, err := os.ReadFile(string(f))
	if err != nil {
		return nil, err
	}

	return utils.ParseYAML(content)
}
