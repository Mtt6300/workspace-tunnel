package main

import (
	"os"

	"github.com/alecthomas/kingpin/v2"
)

var (
	app = kingpin.New("workspace-tunnel", "")

	portforwardCMD          = app.Command("up", "")
	selectedWorkspace       = portforwardCMD.Arg("name", "Workspace name").Required().String()
	selectedWorkspaceConfig = portforwardCMD.Arg("config", "Workspace config file").Required().String()

	getCMD                 = app.Command("get", "")
	resource               = getCMD.Arg("resource", "").Required().String()
	selectedKubeConfigPath = getCMD.Flag("kube-config", "Kube config path").Default("~/.kube/config").String()
)

func init() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
