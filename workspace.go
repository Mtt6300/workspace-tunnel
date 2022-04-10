package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Workspace struct {
	Resource []KubeResource
}

func StartWorkspace(config *rest.Config, Workspace Workspace, stopCh chan struct{}, client *kubernetes.Clientset) {
	var wg sync.WaitGroup
	wg.Add(len(Workspace.Resource))

	terminationSignal := make(chan os.Signal, 1)
	signal.Notify(terminationSignal, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-terminationSignal
		fmt.Println("Stop...")
		close(stopCh)
		wg.Done()
	}()

	for _, r := range Workspace.Resource {
		go func(r KubeResource) {
			defer wg.Done()
			err := StartForwarding(config, r, client)
			if err != nil {
				app.FatalIfError(err, "Error while starting port forwarding.")
			}
			select {
			case <-r.ReadyCh:
				break
			}
		}(r)
	}

	wg.Wait()
}
