package main

import (
	"context"

	v1 "k8s.io/api/core/v1"
	api "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func GetPodListFromService(
	service v1.Service,
	client *kubernetes.Clientset,
) (*v1.PodList, error) {

	set := labels.Set(service.Spec.Selector)
	pod, err := client.CoreV1().Pods(service.Namespace).List(context.Background(), api.ListOptions{
		LabelSelector: set.AsSelector().String(),
	})
	if err != nil {
		return nil, err
	}
	return pod, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
