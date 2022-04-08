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
	Services []Service
}

func StartWorkspace(config *rest.Config, Workspace Workspace, stopCh chan struct{}, client *kubernetes.Clientset) {
	var wg sync.WaitGroup
	wg.Add(len(Workspace.Services))

	terminationSignal := make(chan os.Signal, 1)
	signal.Notify(terminationSignal, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-terminationSignal
		fmt.Println("Stop...")
		close(stopCh)
		wg.Done()
	}()

	for _, service := range Workspace.Services {
		go func(service Service) {
			defer wg.Done()
			err := StartForwarding(config, service, client)
			if err != nil {
				panic(err)
			}

			select {
			case <-service.ReadyCh:
				break
			}
		}(service)
	}

	wg.Wait()
}
