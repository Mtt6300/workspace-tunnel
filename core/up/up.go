package up

import (
	"context"
	"log"
	"os"
	"os/signal"
	"slices"
	"sync"
	"syscall"

	"github.com/Mtt6300/workspace-tunnel/types"
	"gopkg.in/yaml.v3"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func findWorkspace(workspaces []types.WorkspaceConf, workspaceName string, onlyFlag []string) []types.KubeResource {
	for _, workspace := range workspaces {
		if workspace.Name == workspaceName {
			result := make([]types.KubeResource, 0, 20)

			for _, service := range workspace.Services {
				if slices.Contains[[]string, string](onlyFlag, service.Name) || len(onlyFlag) == 0 {
					result = append(result, types.KubeResource{
						Name: service.Name,
						Port: types.Port{
							LocalPort:  service.LocalPort,
							RemotePort: service.RemotePort,
						},
						Namespace:  service.Namespace,
						Type:       types.ServiceResource,
						KubeConfig: workspace.KubeConfigPath,
					})
				}
			}

			for _, pod := range workspace.Pods {
				if slices.Contains[[]string, string](onlyFlag, pod.Name) || len(onlyFlag) == 0 {
					result = append(result, types.KubeResource{
						Name: pod.Name,
						Port: types.Port{
							LocalPort:  pod.LocalPort,
							RemotePort: pod.RemotePort,
						},
						Namespace:  pod.Namespace,
						Type:       types.ServiceResource,
						KubeConfig: workspace.KubeConfigPath,
					})
				}
			}

			return result
		}
	}

	return nil
}

func UpCommand(nameFlag string, configFlag string, onlyFlag []string) {
	content, err := os.ReadFile(configFlag)
	if err != nil {
		log.Fatal(err)
	}

	setupConfig := types.Config{}
	err = yaml.Unmarshal(content, &setupConfig)
	if err != nil {
		log.Fatal(err)
	}

	targetWorkspace := findWorkspace(setupConfig.Workspaces, nameFlag, onlyFlag)
	if targetWorkspace == nil || len(targetWorkspace) == 0 {
		log.Fatalf("could not find worksapce with name %s", nameFlag)
	}

	var wg sync.WaitGroup
	wg.Add(len(targetWorkspace))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sigs
		cancel()
	}()

	for _, resource := range targetWorkspace {
		go func(r types.KubeResource) {
			defer wg.Done()

			kubeConfig, err := clientcmd.BuildConfigFromFlags("", r.KubeConfig)
			if err != nil {
				log.Fatalf("error while loading kubeconfig file: %s", err)
			}

			client, err := kubernetes.NewForConfig(kubeConfig)
			if err != nil {
				log.Fatalf("error while creating kubernetes client.: %s", err)
			}

			err = startForwarding(ctx, kubeConfig, r, client)
			if err != nil {
				log.Fatal(err)
			}

			<-ctx.Done()
		}(resource)
	}

	wg.Wait()
}
