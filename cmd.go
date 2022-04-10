package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app                    = kingpin.New("workspace-tunnel", "")
	selectedKubeConfigPath = app.Flag("kubeconfig", "Kube config path").Default("~/.kube/config").String()

	portforwardCMD          = app.Command("port-forward", "")
	selectedWorkspace       = portforwardCMD.Arg("name", "Workspace name").Required().String()
	selectedWorkspaceConfig = portforwardCMD.Arg("config", "Workspace config file").Required().String()

	getCMD   = app.Command("get", "")
	resource = getCMD.Arg("resource", "").Required().String()
)

func init() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
