package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type NetworkPolicy struct {
	ApiVersion  string            `yaml:"apiVersion"`
	Kind        string            `yaml:"kind"`
	Metadata    map[string]string `yaml:"metadata"`
	Spec        map[string]string `yaml:"spec"`
	PolicyTypes []string
}

func main() {
	// Marshal a NP struct to YAML.
	np := &NetworkPolicy{"networking.k8s.io/v1",
		"NetworkPolicy", map[string]string{"name": "allowingresspolicy", "namespace": "insights-agent"},
		map[string]string{"podSelector": "", "matchLabels": "app:redis\n"},
		},
		PolicyTypes: []string{"Ingress"},
	}

	y, err := yaml.Marshal(&np)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(y))
}
