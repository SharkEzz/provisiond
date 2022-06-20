package loader

import (
	"os"

	"github.com/SharkEzz/provisiond/pkg/types"
)

// FileLoader load a deployment configuration from a file.
type FileLoader string

func (f FileLoader) Load() (*types.Deployment, error) {
	content, err := os.ReadFile(string(f))
	if err != nil {
		return nil, err
	}

	return parseYAML(content)
}
