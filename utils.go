package main

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	api "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func GetListOfPodFromService(
	service v1.Service,
	client *kubernetes.Clientset,
) *v1.PodList {

	set := labels.Set(service.Spec.Selector)
	myPod, err := client.CoreV1().Pods(service.Namespace).List(context.Background(), api.ListOptions{
		LabelSelector: set.AsSelector().String(),
	})
	if err != nil {
		fmt.Println(err)
	}
	return myPod
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
