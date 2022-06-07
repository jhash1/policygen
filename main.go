package main

import (
	"fmt"
)

//Create hardcoded vars for NP
const apiVersion string = "networking.k8s.io/v1"
const kind string = "NetworkPolicy"

type NetworkPolicy struct {
	name       string
	namespace  string
	labelName  string
	labelValue string
	policyType string
}

func main() {
	fmt.Println("YAML Here")
}
