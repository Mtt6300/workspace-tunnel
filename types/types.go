package types

type ResourceConf struct {
	Name       string `yaml:"name"`
	RemotePort int32  `yaml:"port"`
	LocalPort  int32  `yaml:"localPort"`
	Namespace  string `yaml:"namespace"`
}

type WorkspaceConf struct {
	Services       []ResourceConf `yaml:"service"`
	Pods           []ResourceConf `yaml:"pod"`
	Name           string         `yaml:"name"`
	KubeConfigPath string         `yaml:"kubeConfigPath"`
}

type Config struct {
	Workspaces []WorkspaceConf `yaml:"workspace"`
}

type ResourceType string

const (
	ServiceResource ResourceType = "service"
	PodResource     ResourceType = "pod"
)

type Port struct {
	LocalPort  int32
	RemotePort int32
}

type KubeResource struct {
	Name       string
	Port       Port
	Namespace  string
	Type       ResourceType
	KubeConfig string
}
