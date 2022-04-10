package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

func StartForwarding(config *rest.Config, resource KubeResource, client *kubernetes.Clientset) error {
	servicePort, err := FindPodForPortForward(resource, client)
	if err != nil {
		return err
	}
	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward",
		resource.Namespace, servicePort.Name)
	transport, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		return err
	}
	dialer := spdy.NewDialer(upgrader,
		&http.Client{Transport: transport},
		http.MethodPost,
		&url.URL{Scheme: "https",
			Path: path,
			Host: strings.TrimLeft(config.Host, "htps:/"),
		})
	fw, err := portforward.New(dialer,
		[]string{fmt.Sprintf("%d:%d", resource.Port.LocalPort, resource.Port.RemotePort)},
		resource.StopCh,
		resource.ReadyCh,
		resource.Streams.Out,
		resource.Streams.ErrOut,
	)
	if err != nil {
		return err
	}
	return fw.ForwardPorts()
}
