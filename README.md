# Policygen
Network Policy Generator for Kubernetes

Policy generator is a CLI utility for creating YAML output for a Kubernetes network policy.
Current support includes by label. This utility is desinged to make creating a network policy simple
and easy.


Sample:

go run policygen --name appnetworkpolicy --namespace testing --labels app=nginx


apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: john
  namespace: testing
spec:
  podSelector:
    matchLabels:
      app: sql
  policyTypes:
  - Ingress
  - Egress
