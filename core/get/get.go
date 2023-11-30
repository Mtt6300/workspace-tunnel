package get

import (
	"context"
	"log"
	"os"

	"github.com/Mtt6300/workspace-tunnel/types"
	"github.com/olekukonko/tablewriter"
	apiv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetCommand(resourceFlag string, kubeConfigFlag string) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigFlag)
	if err != nil {
		log.Fatalf("error while loading kubeconfig file: %s", err)
	}

	client, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalf("error while creating kubernetes client.: %s", err)
	}

	var data [][]string

	switch resourceFlag {
	case string(types.ServiceResource):
		resourceList, err := client.CoreV1().Services("").List(context.Background(), apiv1.ListOptions{})
		if err != nil {
			log.Fatalf("error while fetching %s from cluster: %s", resourceFlag, err)
		}

		for _, resource := range resourceList.Items {
			var portString string = ""
			resourcePodList, err := getPodListFromService(resource, client)
			if err != nil {
				log.Fatalf("error while fetching pods from service: %s", err)
			}

			for _, pod := range resourcePodList.Items {
				portString = generatePortsStringFormat(pod.Spec.Containers)
			}

			data = append(data, []string{
				resource.Namespace,
				resource.Name,
				portString,
			})
		}

	case string(types.PodResource):
		resourceList, err := client.CoreV1().Pods("").List(context.Background(), apiv1.ListOptions{})
		if err != nil {
			log.Fatalf("error while fetching %s from cluster: %s", resourceFlag, err)
		}

		for _, resource := range resourceList.Items {
			var portString string = ""
			portString = generatePortsStringFormat(resource.Spec.Containers)

			data = append(data, []string{
				resource.Namespace,
				resource.Name,
				portString,
			})
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Namespace", resourceFlag, "Ports"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
