package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case portforwardCMD.FullCommand():
		LoadWorkspaceConfig(*selectedWorkspaceConfig)
		fmt.Println()

		stopChannel := make(chan struct{}, 1)
		stream := genericclioptions.IOStreams{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		}

		var myWorkspace Workspace
		for _, workspace := range Appconfig.Workspaces {
			if workspace.Name == *selectedWorkspace {
				for _, service := range workspace.Services {
					myWorkspace.Resource = append(myWorkspace.Resource, insertResourceToWorkspace(service, Service, stream, stopChannel, workspace.KubeConfigPath))
				}
				for _, pod := range workspace.Pods {
					myWorkspace.Resource = append(myWorkspace.Resource, insertResourceToWorkspace(pod, Pod, stream, stopChannel, workspace.KubeConfigPath))
				}
			}
		}
		if len(myWorkspace.Resource) == 0 {
			app.FatalIfError(fmt.Errorf("no workspace or resource found"), "")
		}

		StartWorkspace(myWorkspace, stopChannel)

	case getCMD.FullCommand():
		if contains(ResourceList, ResourceType(*resource)) {
			_, client := LoadKubeConfig(*selectedKubeConfigPath)
			ShowResourceDetails(*resource, client)
		} else {
			app.Errorf("Resource not found")
		}
	}

}

func insertResourceToWorkspace(
	r resource_conf,
	resourceType ResourceType,
	stream genericclioptions.IOStreams,
	stopChannel chan struct{},
	config string,
) KubeResource {
	return KubeResource{
		Name: r.Name,
		Port: Port{
			LocalPort:  r.LocalPort,
			RemotePort: r.RemotePort,
		},
		Namespace:  r.Namespace,
		Streams:    stream,
		StopCh:     stopChannel,
		ReadyCh:    make(chan struct{}),
		Type:       resourceType,
		KubeConfig: config,
	}
}
