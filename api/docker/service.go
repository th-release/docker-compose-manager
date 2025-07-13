package docker

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// DockerCompose represents the structure of a docker-compose.yml file
type DockerCompose struct {
	Version  string              `json:"version" yaml:"version"`
	Services map[string]*Service `json:"services" yaml:"services"`
	Networks map[string]*Network `json:"networks,omitempty" yaml:"networks,omitempty"`
	Volumes  map[string]*Volume  `json:"volumes,omitempty" yaml:"volumes,omitempty"` // 포인터로 변경
}

// Service represents a service in docker-compose
type Service struct {
	Image         string            `json:"image,omitempty" yaml:"image,omitempty"`
	ContainerName string            `json:"container_name,omitempty" yaml:"container_name,omitempty"`
	Ports         []string          `json:"ports,omitempty" yaml:"ports,omitempty"`
	Environment   map[string]string `json:"environment,omitempty" yaml:"environment,omitempty"`
	Volumes       []string          `json:"volumes,omitempty" yaml:"volumes,omitempty"`
	DependsOn     []string          `json:"depends_on,omitempty" yaml:"depends_on,omitempty"`
	Networks      []string          `json:"networks,omitempty" yaml:"networks,omitempty"`
	Restart       string            `json:"restart,omitempty" yaml:"restart,omitempty"`
	Command       string            `json:"command,omitempty" yaml:"command,omitempty"`
	WorkingDir    string            `json:"working_dir,omitempty" yaml:"working_dir,omitempty"`
	User          string            `json:"user,omitempty" yaml:"user,omitempty"`
	Expose        []string          `json:"expose,omitempty" yaml:"expose,omitempty"`
	Logging       *LoggingConfig    `json:"logging,omitempty" yaml:"logging,omitempty"`
	Devices       []string          `json:"devices,omitempty" yaml:"devices,omitempty"`
	CapAdd        []string          `json:"cap_add,omitempty" yaml:"cap_add,omitempty"`
	CapDrop       []string          `json:"cap_drop,omitempty" yaml:"cap_drop,omitempty"`
	Privileged    bool              `json:"privileged,omitempty" yaml:"privileged,omitempty"`
	SecurityOpt   []string          `json:"security_opt,omitempty" yaml:"security_opt,omitempty"`
	Sysctls       []string          `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`
	Ulimits       map[string]string `json:"ulimits,omitempty" yaml:"ulimits,omitempty"`
	Hostname      string            `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Labels        map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	HealthCheck   *HealthCheck      `json:"healthcheck,omitempty" yaml:"healthcheck,omitempty"`
	Deploy        *Deploy           `json:"deploy,omitempty" yaml:"deploy,omitempty"`
}

// LoggingConfig represents logging configuration
type LoggingConfig struct {
	Driver  string            `json:"driver,omitempty" yaml:"driver,omitempty"`
	Options map[string]string `json:"options,omitempty" yaml:"options,omitempty"`
}

// HealthCheck represents health check configuration
type HealthCheck struct {
	Test        []string `json:"test,omitempty" yaml:"test,omitempty"`
	Interval    string   `json:"interval,omitempty" yaml:"interval,omitempty"`
	Timeout     string   `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Retries     int      `json:"retries,omitempty" yaml:"retries,omitempty"`
	StartPeriod string   `json:"start_period,omitempty" yaml:"start_period,omitempty"`
}

// Deploy represents deploy configuration for swarm mode
type Deploy struct {
	Replicas      int               `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	RestartPolicy *RestartPolicy    `json:"restart_policy,omitempty" yaml:"restart_policy,omitempty"`
	Resources     *Resources        `json:"resources,omitempty" yaml:"resources,omitempty"`
	Labels        map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}

// RestartPolicy represents restart policy configuration
type RestartPolicy struct {
	Condition   string `json:"condition,omitempty" yaml:"condition,omitempty"`
	Delay       string `json:"delay,omitempty" yaml:"delay,omitempty"`
	MaxAttempts int    `json:"max_attempts,omitempty" yaml:"max_attempts,omitempty"`
	Window      string `json:"window,omitempty" yaml:"window,omitempty"`
}

// Resources represents resource limits and reservations
type Resources struct {
	Limits       *ResourceSpec `json:"limits,omitempty" yaml:"limits,omitempty"`
	Reservations *ResourceSpec `json:"reservations,omitempty" yaml:"reservations,omitempty"`
}

// ResourceSpec represents CPU and memory specifications
type ResourceSpec struct {
	CPUs   string `json:"cpus,omitempty" yaml:"cpus,omitempty"`
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`
}

// Network represents a network configuration
type Network struct {
	Driver   string `json:"driver,omitempty" yaml:"driver,omitempty"`
	Name     string `json:"name,omitempty" yaml:"name,omitempty"`
	External bool   `json:"external,omitempty" yaml:"external,omitempty"`
}

// Volume represents a volume configuration
type Volume struct {
	Driver string `json:"driver,omitempty" yaml:"driver,omitempty"`
}

// Docker Compose 재시작 함수
func RestartDockerCompose(composePath string) error {
	composeDir := filepath.Dir(composePath)

	// docker compose down 실행
	downCmd := exec.Command("docker", "compose", "down")
	downCmd.Dir = composeDir

	if output, err := downCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("docker compose down failed: %v, output: %s", err, string(output))
	}

	// docker compose up -d 실행
	upCmd := exec.Command("docker", "compose", "up", "-d")
	upCmd.Dir = composeDir

	if output, err := upCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("docker compose up -d failed: %v, output: %s", err, string(output))
	}

	return nil
}

// LoadDockerCompose loads a docker-compose.yml file
func LoadDockerCompose(filepath string) (*DockerCompose, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var compose DockerCompose
	err = yaml.Unmarshal(data, &compose)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %v", err)
	}

	// Initialize maps if they're nil
	if compose.Services == nil {
		compose.Services = make(map[string]*Service)
	}
	if compose.Networks == nil {
		compose.Networks = make(map[string]*Network)
	}
	if compose.Volumes == nil {
		compose.Volumes = make(map[string]*Volume)
	}

	return &compose, nil
}

func (v *Volume) MarshalYAML() (interface{}, error) {
	if v == nil {
		return "", nil // 빈 문자열 반환
	}
	return struct {
		Driver string `yaml:"driver,omitempty"`
	}{
		Driver: v.Driver,
	}, nil
}

// SaveDockerCompose saves a DockerCompose struct to file
func SaveDockerCompose(compose *DockerCompose, filepath string) error {
	// 빈 Volume 구조체 생성
	emptyVolume := &Volume{}

	// 모든 nil 포인터를 빈 구조체로 교체
	if compose.Volumes != nil {
		for k, v := range compose.Volumes {
			if v == nil {
				compose.Volumes[k] = emptyVolume
			}
		}
	}

	data, err := yaml.Marshal(compose)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %v", err)
	}

	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

// CreateNewDockerCompose creates a new empty docker-compose structure
func CreateNewDockerCompose(version string) *DockerCompose {
	if version == "" {
		version = "3.8"
	}
	return &DockerCompose{
		Version:  version,
		Services: make(map[string]*Service),
		Networks: make(map[string]*Network),
		Volumes:  make(map[string]*Volume),
	}
}

// AddService adds a service to the compose
func AddService(compose *DockerCompose, name string, service *Service) {
	if compose.Services == nil {
		compose.Services = make(map[string]*Service)
	}
	compose.Services[name] = service
}

// UpdateService updates an existing service
func UpdateService(compose *DockerCompose, name string, service *Service) error {
	if _, exists := compose.Services[name]; !exists {
		return fmt.Errorf("service '%s' does not exist", name)
	}
	compose.Services[name] = service
	return nil
}

// DeleteService removes a service from the compose
func DeleteService(compose *DockerCompose, name string) error {
	if _, exists := compose.Services[name]; !exists {
		return fmt.Errorf("service '%s' does not exist", name)
	}
	delete(compose.Services, name)
	return nil
}

// GetService retrieves a service by name
func GetService(compose *DockerCompose, name string) (*Service, error) {
	service, exists := compose.Services[name]
	if !exists {
		return nil, fmt.Errorf("service '%s' does not exist", name)
	}
	return service, nil
}

// ListServices returns all service names
func ListServices(compose *DockerCompose) []string {
	var services []string
	for name := range compose.Services {
		services = append(services, name)
	}
	return services
}

// AddNetwork adds a network to the compose
func AddNetwork(compose *DockerCompose, name string, network *Network) {
	if compose.Networks == nil {
		compose.Networks = make(map[string]*Network)
	}
	compose.Networks[name] = network
}

// DeleteNetwork removes a network from the compose
func DeleteNetwork(compose *DockerCompose, name string) error {
	if _, exists := compose.Networks[name]; !exists {
		return fmt.Errorf("network '%s' does not exist", name)
	}
	delete(compose.Networks, name)
	return nil
}

// AddVolume adds a volume to the compose
func AddVolume(compose *DockerCompose, name string, volume *Volume) {
	if compose.Volumes == nil {
		compose.Volumes = nil
	}
	compose.Volumes[name] = volume
}

// DeleteVolume removes a volume from the compose
func DeleteVolume(compose *DockerCompose, name string) error {
	if _, exists := compose.Volumes[name]; !exists {
		return fmt.Errorf("volume '%s' does not exist", name)
	}
	delete(compose.Volumes, name)
	return nil
}

// PrintCompose prints the compose structure as YAML
func PrintCompose(compose *DockerCompose) {
	data, err := yaml.Marshal(compose)
	if err != nil {
		fmt.Printf("Error marshaling YAML: %v\n", err)
		return
	}
	fmt.Println(string(data))
}
