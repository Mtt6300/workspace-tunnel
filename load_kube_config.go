package main

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func LoadKubeConfig(path string) (*rest.Config, *kubernetes.Clientset) {
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		app.FatalIfError(err, "Error while loading kubeconfig file.")
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		app.FatalIfError(err, "Error while creating kubernetes client.")
	}
	fmt.Println("Kube Config loaded: ", path)
	return config, client
}
