package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type NetworkPolicy struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	// Metadata   struct {
	// 	Name      string `yaml:"name"`
	// 	Namespace string `yaml:"namespace"`
	// } `yaml:"metadata"`
	// Spec struct {
	// 	PodSelector struct {
	// 		MatchLabels struct {
	// 			App string `yaml:"app"`
	// 		} `yaml:"matchLabels"`
	// 	} `yaml:"podSelector"`
	// 	PolicyTypes []string `yaml:"policyTypes"`
	// 	Ingress     []struct {
	// 		From []struct {
	// 			PodSelector struct {
	// 				MatchLabels struct {
	// 					App string `yaml:"app"`
	// 				} `yaml:"matchLabels"`
	// 			} `yaml:"podSelector"`
	// 		} `yaml:"from"`
	// 	} `yaml:"ingress"`
	// } `yaml:"spec"`
}

func main() {
	// Marshal a Person struct to YAML.
	np := &NetworkPolicy{"networking.k8s.io/v1", "NetworkPolicy"}
	y, err := yaml.Marshal(np)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
}
