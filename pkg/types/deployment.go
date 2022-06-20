package types

type SSHConfig struct {
	Host     string
	Port     uint16
	Type     string
	Username string
	Password string
	KeyFile  string
}

// A deployment is the configuration file that contain all the instructions to deploy your project.
type Deployment struct {
	Name      string
	Variables map[string]string
	Config    struct {
		SSH map[string]SSHConfig
	}
	Jobs []map[string]any
}
