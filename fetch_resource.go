package main

import (
	"context"
	"os"

	"github.com/olekukonko/tablewriter"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ShowResourceDetails(selectedResource string, client *kubernetes.Clientset) {
	var data [][]string

	switch selectedResource {
	case string(Service):
		resourceList, err := client.CoreV1().Services("").List(context.Background(), v1.ListOptions{})
		if err != nil {
			app.FatalIfError(err, "Error while fetching "+*resource+" from cluster.")
		}

		for _, resource := range resourceList.Items {
			var portString string = ""
			resourcePodList, err := GetPodListFromService(resource, client)
			if err != nil {
				app.FatalIfError(err, "Error while fetching pods from service.")
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
	case string(Pod):
		resourceList, err := client.CoreV1().Pods("").List(context.Background(), v1.ListOptions{})
		if err != nil {
			app.FatalIfError(err, "Error while fetching "+*resource+" from cluster.")
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
	table.SetHeader([]string{"Namespace", selectedResource, "Ports"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

}
