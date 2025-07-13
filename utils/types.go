package utils

type Config struct {
	DockerPath string `json:"DOCKER_PATH"`
	Port       int    `json:"PORT"`
}

// DockerCompose represents the structure of a docker-compose.yml file
type DockerCompose struct {
	Services map[string]*Service      `yaml:"services"`
	Volumes  []string                 `yaml:"volumes,omitempty,flow"`
	Networks map[string]NetworkConfig `yaml:"networks,omitempty"`
}

// Service represents a service in the docker-compose file
type Service struct {
	Image       string            `yaml:"image"`
	Restart     string            `yaml:"restart"`
	Ports       []string          `yaml:"ports"`
	CapAdd      []string          `yaml:"cap_add"`
	Sysctls     []string          `yaml:"sysctls"`
	Devices     []string          `yaml:"devices"`
	Environment map[string]string `yaml:"environment"`
	Volumes     []string          `yaml:"volumes"`
	Logging     LoggingConfig     `yaml:"logging"`
	Networks    []string          `yaml:"networks"`
}

// LoggingConfig represents the logging configuration
type LoggingConfig struct {
	Options map[string]string `yaml:"options"`
}

// NetworkConfig represents the network configuration
type NetworkConfig struct {
	Name     string `yaml:"name"`
	External bool   `yaml:"external"`
}

// BasicResponse represents a standard API response
type BasicResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
