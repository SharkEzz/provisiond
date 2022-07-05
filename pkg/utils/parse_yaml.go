package utils

import (
	"github.com/SharkEzz/provisiond/pkg/deployment"
	"gopkg.in/yaml.v3"
)

// parseYaml parse the entire configuration file and return the corresponding deployment.
func ParseYAML(content []byte) (*deployment.Deployment, error) {
	deployment := &deployment.Deployment{}
	err := yaml.Unmarshal(content, deployment)
	if err != nil {
		return nil, err
	}

	return deployment, err
}
