package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type NetworkPolicy struct {
	ApiVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Spec       map[string]string `yaml:"spec"`
}

//next part is coming up with the whole YAML such as the spec file,
//creating the framework for Cobra CLI and avail options. Ill add to spreadsheet
//add user input

func main() {
	// Marshal a NP struct to YAML.
	np := &NetworkPolicy{"networking.k8s.io/v1",
		"NetworkPolicy",
		map[string]string{"name": "allowingresspolicy", "namespace": "insights-agent"},
		map[string]string{"podSelector": "", "matchLabels": "{}"},
	}
	y, err := yaml.Marshal(np)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
}
