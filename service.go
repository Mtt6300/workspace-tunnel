package main

import "k8s.io/cli-runtime/pkg/genericclioptions"

type Service struct {
	Name      string
	Port      int
	LocalPort int
	Namespace string
	Streams   genericclioptions.IOStreams
	StopCh    <-chan struct{}
	ReadyCh   chan struct{}
}
