package utils

import (
	"os"

	"github.com/SharkEzz/provisiond/pkg/deployment"
	"gopkg.in/yaml.v3"
)

// parseYaml expand all the variables found, parse the entire configuration file and return the corresponding deployment.
func ParseYAML(content []byte) (*deployment.Deployment, error) {
	// TODO: better way to process environment variable?

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

	deployment := &deployment.Deployment{}
	err = yaml.Unmarshal(content, deployment)
	if err != nil {
		return nil, err
	}

	return deployment, err
}
