package main

import (
	"fmt"
	"io/ioutil"
	"log"

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
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
func LoadWorkspaceConfig(path string) {
	Appconfig.LoadWorkspaceConfig(path)
	fmt.Println("Workspace config loaded: ", path)

}
