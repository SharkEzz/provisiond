package executor

import (
	"os"

	"github.com/SharkEzz/provisiond/pkg/logging"
	"gopkg.in/yaml.v3"
)

// A config file provide additional information about how provisiond should run.
type Config struct {
	JobTimeout        uint32 `yaml:"job_timeout"`
	DeploymentTimeout uint32 `yaml:"deployment_timeout"`
	AllowFailure      bool   `yaml:"allow_failure"`
}

// Parse the config.yaml file (if existing) and returb a Config struct.
func LoadConfig() (*Config, error) {
	if _, err := os.Stat("./config.yaml"); err != nil {
		return nil, nil
	}

	config := &Config{}

	content, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}
	logging.LogOut("Loaded config from ./config.yaml")

	return config, nil
}
