package get

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	apiv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func getPodListFromService(
	service v1.Service,
	client *kubernetes.Clientset,
) (*v1.PodList, error) {

	set := labels.Set(service.Spec.Selector)
	pod, err := client.CoreV1().Pods(service.Namespace).List(context.Background(), apiv1.ListOptions{
		LabelSelector: set.AsSelector().String(),
	})
	if err != nil {
		return nil, err
	}
	return pod, nil
}

func generatePortsStringFormat(container []v1.Container) string {
	var message string
	for _, c := range container {
		for _, p := range c.Ports {
			message += fmt.Sprintf("%s:%d ", p.Name, p.ContainerPort)
		}
	}
	return message
}
