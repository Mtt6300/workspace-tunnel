package main

import "k8s.io/cli-runtime/pkg/genericclioptions"

var ResourceList []ResourceType = []ResourceType{Service, Pod}

type ResourceType string

const (
	Service ResourceType = "service"
	Pod     ResourceType = "pod"
)

type Port struct {
	LocalPort  int32
	RemotePort int32
}

type KubeResource struct {
	Name      string
	Port      Port
	Namespace string
	Streams   genericclioptions.IOStreams
	StopCh    <-chan struct{}
	ReadyCh   chan struct{}
	Type      ResourceType
}
