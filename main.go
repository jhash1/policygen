package main

import (
	"flag"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type Spec struct {
	PodSelector NetworkPolicySpecPodSelector `yaml:"podSelector"`
	PolicyTypes []string                     `yaml:"policyTypes"`
}

type NetworkPolicySpecPodSelector struct {
	MatchLabels map[string]string `yaml:"matchLabels"`
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
			},
			PolicyTypes: []string{"Ingress", "Egress"},
		},
	}
	flag.StringVar(&np.Metadata.Name, "name", "", "Specify Name for the Network Policy")
	flag.StringVar(&np.Metadata.Namespace, "namespace", "", "Specify Namespace for the Network Policy")
	flag.Parse()
	yamlData, _ := yaml.Marshal(&np)
	fmt.Println(string(yamlData))
}
