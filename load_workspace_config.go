package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type service_conf struct {
	Name      string `yaml:"name"`
	Port      int    `yaml:"port"`
	LocalPort int    `yaml:"localPort"`
	Namespace string `yaml:"namespace"`
}

type workspace_conf struct {
	Services []service_conf `yaml:"service"`
	Name     string         `yaml:"name"`
}

type config struct {
	Workspaces []workspace_conf `yaml:"workspace"`
}

var Appconfig config

func (c *config) LoadWorkspaceConfig(path string) *config {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		app.FatalIfError(err, "Error while loading workspace config file.")
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		app.FatalIfError(err, "Error while parsing workspace config file.")
	}
	return c
}
func LoadWorkspaceConfig(path string) {
	Appconfig.LoadWorkspaceConfig(path)
	fmt.Println("Workspace config loaded: ", path)

}
