package up

import (
	"context"
	"errors"

	"github.com/Mtt6300/workspace-tunnel/types"
	v1 "k8s.io/api/core/v1"
	api "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func findPodForPortForward(resource types.KubeResource, client *kubernetes.Clientset) (v1.Pod, error) {
	switch resource.Type {
	case types.ServiceResource:
		toFindService, err := client.CoreV1().Services(resource.Namespace).Get(context.Background(), resource.Name, api.GetOptions{})
		if err != nil {
			return v1.Pod{}, err
		}
		servicePodList, err := getPodListFromService(*toFindService, client)
		if err != nil {
			return v1.Pod{}, err
		}

		return servicePodList.Items[0], nil

	case types.PodResource:
		pod, err := client.CoreV1().Pods(resource.Namespace).Get(context.Background(), resource.Name, api.GetOptions{})
		if err != nil {
			return v1.Pod{}, err
		}

		return *pod, nil
	}

	return v1.Pod{}, errors.New("unsupported resource type")
}

func getPodListFromService(
	service v1.Service,
	client *kubernetes.Clientset,
) (*v1.PodList, error) {
	set := labels.Set(service.Spec.Selector)
	return client.CoreV1().Pods(service.Namespace).List(context.Background(), api.ListOptions{
		LabelSelector: set.AsSelector().String(),
	})
}
