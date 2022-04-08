package main

import (
	"context"
	"errors"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var comparableLabel string = "app.kubernetes.io/name"

func SelectPodFromService(service Service, client *kubernetes.Clientset) (string, string, error) {
	ServiceInfo, err := client.CoreV1().Services(service.Namespace).Get(context.Background(), service.Name, v1.GetOptions{})
	if err != nil {
		return "", "", err
	}
	selectedLabel := ServiceInfo.Labels[comparableLabel]
	podsList, err := client.CoreV1().Pods(service.Namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return "", "", err
	}
	for _, pod := range podsList.Items {
		if pod.Labels[comparableLabel] == selectedLabel {
			var portString string
			for _, container := range pod.Spec.Containers {
				for _, port := range container.Ports {
					portString += fmt.Sprintf(":%d", port.ContainerPort)
				}
			}
			return pod.Name, portString, nil
		}
	}
	return "", "", errors.New("no pod found")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
