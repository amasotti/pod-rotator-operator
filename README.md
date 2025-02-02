# Pod Rotator Operator 

[![CI](https://github.com/amasotti/pod-rotator-operator/actions/workflows/ci.yml/badge.svg)](https://github.com/amasotti/pod-rotator-operator/actions/workflows/ci.yml)
[![Release](https://github.com/amasotti/pod-rotator-operator/actions/workflows/release.yml/badge.svg)](https://github.com/amasotti/pod-rotator-operator/actions/workflows/release.yml)

A Kubernetes operator that automatically rotates pods in deployments on a configurable schedule. 
Useful for applications that need periodic restarts to prevent memory leaks or maintain optimal performance.

## Features

- 🕒 Configurable rotation schedule using cron syntax
- 🔄 Automatic pod rotation with zero-downtime
- 🎯 Deployment-specific targeting
- 🔍 Built-in monitoring and health checks
- 🛡️ Secure by default with RBAC policies

## Installation

### Quick Start

```bash
# Install the operator
kubectl apply -f https://github.com/amasotti/pod-rotator-operator/releases/latest/download/release.yaml

# Create a pod rotation schedule (example)
kubectl apply -f - <<EOF
apiVersion: apps.tonihacks.com/v1alpha1
kind: CustomPodRotator
metadata:
  name: my-app-rotator
spec:
  targetDeployment: my-app
  schedule: "0 */6 * * *"  # Every 6 hours
EOF
```

### Usage Example

see example [helm chart](./examples/hello-rotator) in `examples/hello-rotator`

## How it Works

The operator:
1. Watches for CustomPodRotator resources
2. On schedule, triggers a rolling update of the target deployment
3. Ensures zero-downtime by respecting deployment strategies
4. Updates status with last rotation time with a custom annotation

## Why a custom operator

The answer is simple, this is a learning project and my first attempt at creating a custom operator.
There are other options, all with their pros and cons that would perfectly fit this use case and are in some cases a better choice.

- [ArgoCD Rollout](https://argoproj.github.io/rollouts/) - ideal if you are already using ArgoCD, but it's a bit overkill for this use case.
- [Kured](https://kured.dev/) - a great tool for node reboots, but it's not designed for pod rotation.
- [Kube-monkey](https://github.com/asobti/kube-monkey) - a chaos engineering tool for Kubernetes, but it's not designed for pod rotation, it reaches however a similar goal by randomly deleting pods.
- [Kubernetes CronJob](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/) - a simple solution, but it lacks the ability to handle edge cases like deployment not found.
- Sidecar Containers: you can also use a sidecar container with e.g. `kubectl` on it to delete pods, but it's not as flexible as an operator.


## Development

### Prerequisites

- Docker
- Go 1.23+
- Kubernetes cluster (local or remote)
- Operator SDK v1.31.0+

### Local Development

see [Documentation](./docs/STEPS.md)

## Docker Images

Official images are available on Docker Hub:
- [katom92/pod-rotator-operator](https://hub.docker.com/r/katom92/pod-rotator-operator/tags)

## Contributing

This is a small side-project with the main goal of improving my understanding of k8s custom operators.
At the moment it is unlikely that the project will grow further, but I would be happy to get in touch and also very
happy if you fork this repo or use as basis for your own side projects.

## License

[Apache 2.0](LICENSE)
