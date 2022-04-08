package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	config, client := LoadKubeConfig(*selectedKubeConfigPath)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case portforwardCMD.FullCommand():
		LoadWorkspaceConfig(*selectedWorkspaceConfig)

		stopChannel := make(chan struct{}, 1)
		stream := genericclioptions.IOStreams{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		}
		var Ws Workspace
		for _, workspace := range Appconfig.Workspaces {
			if workspace.Name == *selectedWorkspace {
				for _, service := range workspace.Services {
					Ws.Services = append(Ws.Services, Service{
						Name:      service.Name,
						Port:      service.Port,
						LocalPort: service.LocalPort,
						Namespace: service.Namespace,
						Streams:   stream,
						StopCh:    stopChannel,
						ReadyCh:   make(chan struct{}),
					})
				}
			}
		}
		StartWorkspace(config, Ws, stopChannel, client)

	case getCMD.FullCommand():
		if contains(ResourceList, *resource) {
			ShowResourceDetails(*resource, client)
		} else {
			fmt.Println("Resource not found")
		}
	}

}
