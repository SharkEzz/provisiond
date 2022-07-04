package loader

import (
	"strings"

	"github.com/SharkEzz/provisiond/internal/loader"
	"github.com/SharkEzz/provisiond/pkg/deployment"
)

// The Loader interface represent a custom loader for loading a deployment configuration.
type Loader interface {
	Load() (*deployment.Deployment, error)
}

// GetLoader return the correct type of Loader regarding what type of data is passed to it.
func GetLoader(name string) Loader {
	// If the name is multiline, treat it as a string
	if len(strings.Split(name, "\n")) > 1 {
		return loader.StringLoader(name)
	}

	return loader.FileLoader(name)
}
