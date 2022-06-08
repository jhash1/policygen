package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type NetworkPolicy struct {
	ApiVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
}

func main() {
	// Marshal a Person struct to YAML.
	np := &NetworkPolicy{"networking.k8s.io/v1",
		"NetworkPolicy",
		map[string]string{"name": "allowingresspolicy", "namespace": "insights-agent"},
	}
	y, err := yaml.Marshal(np)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
}
