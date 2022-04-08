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

func StartForwarding(config *rest.Config, service Service, client *kubernetes.Clientset) error {
	servicePod, _, err := SelectPodFromService(service, client)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward",
		service.Namespace, servicePod)
	transport, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		return err
	}

	dialer := spdy.NewDialer(upgrader,
		&http.Client{Transport: transport},
		http.MethodPost,
		&url.URL{Scheme: "https",
			Path: path,
			Host: strings.TrimLeft(config.Host, "https://"),
		})
	fw, err := portforward.New(dialer,
		[]string{fmt.Sprintf("%d:%d", service.LocalPort, service.Port)},
		service.StopCh,
		service.ReadyCh,
		service.Streams.Out,
		service.Streams.ErrOut,
	)
	if err != nil {
		return err
	}
	return fw.ForwardPorts()
}
