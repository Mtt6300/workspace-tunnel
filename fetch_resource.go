package main

import (
	"context"
	"os"

	"github.com/olekukonko/tablewriter"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func FetchResource(selectedResource string, client *kubernetes.Clientset) {
	var data [][]string

	switch selectedResource {
	case "service":
		resourceList, err := client.CoreV1().Services("").List(context.Background(), v1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for _, resource := range resourceList.Items {
			_, portString, _ := SelectPodFromService(Service{
				Name:      resource.Name,
				Namespace: resource.Namespace,
			}, client)
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