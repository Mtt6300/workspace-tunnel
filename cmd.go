package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app                    = kingpin.New("chat", "A command-line chat application.")
	selectedKubeConfigPath = app.Flag("kubeconfig", "kube config path").Default("~/.kube/config").String()

	portforwardCMD          = app.Command("port-forward", "open port for your workspace")
	selectedWorkspace       = portforwardCMD.Arg("name", "Your workspace name").Required().String()
	selectedWorkspaceConfig = portforwardCMD.Arg("workspace", "workspace yaml config").Required().String()

	getCMD   = app.Command("get", "get resource list")
	resource = getCMD.Arg("resource", "your resource.").Required().String()
)

func init() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
