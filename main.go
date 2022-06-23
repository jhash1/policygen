package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type Spec struct {
	PodSelector NetworkPolicySpecPodSelector `yaml:"podSelector"`
}

type NetworkPolicySpecPodSelector struct {
	MatchLabels map[string]string `yaml:"matchLabels"`
	PolicyTypes []string          `yaml:"policyTypes"`
}

type Ingress struct {
}

type NetworkPolicy struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

func main() {
	np := NetworkPolicy{
		ApiVersion: "networking.k8s.io/v1",
		Kind:       "NetworkPolicy",
		Metadata: Metadata{
			Name:      "allow-ingress",
			Namespace: "default",
		},
		Spec: Spec{
			PodSelector: NetworkPolicySpecPodSelector{
				MatchLabels: map[string]string{
					"env": "prod"},
				PolicyTypes: []string{"Ingress", "Egress"},
			},
		},
	}

	yamlData, _ := yaml.Marshal(&np)
	fmt.Println(string(yamlData))
}
