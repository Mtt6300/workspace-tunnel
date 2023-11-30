package up

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Mtt6300/workspace-tunnel/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

func startForwarding(ctx context.Context, config *rest.Config, resource types.KubeResource, client *kubernetes.Clientset) error {
	servicePort, err := findPodForPortForward(resource, client)
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
		ctx.Done(),
		make(chan struct{}),
		os.Stdout,
		os.Stderr,
	)
	if err != nil {
		return err
	}
	return fw.ForwardPorts()
}
