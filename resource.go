package main

import "k8s.io/cli-runtime/pkg/genericclioptions"

type Port struct {
	LocalPort  int32
	RemotePort int32
}

type Service struct {
	Name      string
	Port      Port
	Namespace string
	Streams   genericclioptions.IOStreams
	StopCh    <-chan struct{}
	ReadyCh   chan struct{}
}
