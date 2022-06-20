package loader

import (
	"os"
	"strings"

	"github.com/SharkEzz/provisiond/pkg/types"
	"gopkg.in/yaml.v3"
)

// The Loader interface represent a custom loader for loading a deployment configuration.
type Loader interface {
	Load() (*types.Deployment, error)
}

// GetLoader return the correct type of Loader regarding what type of data is passed to it.
func GetLoader(name string) Loader {
	// If the name is multiline, treat it as a string
	if len(strings.Split(name, "\n")) > 1 {
		return StringLoader(name)
	}

	return FileLoader(name)
}

// parseYaml expand all the variables found then parse the entire configuration file.
func parseYAML(content []byte) (*types.Deployment, error) {
	type variables struct {
		Variables map[string]string
	}
	vars := &variables{}
	err := yaml.Unmarshal(content, vars)
	if err != nil {
		return nil, err
	}

	for name, value := range vars.Variables {
		err = os.Setenv(name, value)
		if err != nil {
			return nil, err
		}
	}
	content = []byte(os.ExpandEnv(string(content)))

	// Prevent leaking of environment variables by unsetting them after expanding
	for name := range vars.Variables {
		err = os.Unsetenv(name)
		if err != nil {
			return nil, err
		}
	}

	deployment := &types.Deployment{}
	err = yaml.Unmarshal(content, deployment)
	if err != nil {
		return nil, err
	}

	return deployment, err
}
