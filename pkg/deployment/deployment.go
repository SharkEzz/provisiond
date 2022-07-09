package deployment

// The SSHConfig struct hold a SSH configuration for a single host
type SSHConfig struct {
	Host       string
	Port       uint16
	Type       string
	Username   string
	Password   string
	KeyFile    string
	KeyContent string
	KeyPass    string
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
