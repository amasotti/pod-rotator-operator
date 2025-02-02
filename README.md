# Pod Rotator Operator 

[![CI](https://github.com/amasotti/pod-rotator-operator/actions/workflows/ci.yml/badge.svg)](https://github.com/amasotti/pod-rotator-operator/actions/workflows/ci.yml)

This repo contains a custom Kubernetes operator using the Operator SDK. 
The operator rotates pods in a deployment by deleting the oldest pod and creating a new one on a schedule. It also 
provides a custom resource definition (CRD) to configure the rotation schedule and takes care of edge cases, like deployment not found.

### Why a custom operator

The answer is simple, this is a learning project and my first attempt at creating a custom operator.
There are other options, all with their pros and cons that would perfectly fit this use case and are in some cases a better choice.

- [ArgoCD Rollout](https://argoproj.github.io/rollouts/) - ideal if you are already using ArgoCD, but it's a bit overkill for this use case.
- [Kured](https://kured.dev/) - a great tool for node reboots, but it's not designed for pod rotation.
- [Kube-monkey](https://github.com/asobti/kube-monkey) - a chaos engineering tool for Kubernetes, but it's not designed for pod rotation, it reaches however a similar goal by randomly deleting pods.
- [Kubernetes CronJob](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/) - a simple solution, but it lacks the ability to handle edge cases like deployment not found.
- Sidecar Containers: you can also use a sidecar container with e.g. `kubectl` on it to delete pods, but it's not as flexible as an operator.

### How it works

_tbd_

## Tooling

- [Operator SDK](https://sdk.operatorframework.io/) - a framework that uses the controller-runtime library to make writing operators easier.
- [Kubebuilder](https://book.kubebuilder.io/) - a Kubernetes controller building SDK that uses the controller-runtime library.
- [OLM](https://olm.operatorframework.io/) - Operator Lifecycle Manager, a tool to install, manage, and upgrade operators and their dependencies in a cluster.

## Development

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Operator SDK](https://sdk.operatorframework.io/docs/installation/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Golang](https://go.dev/)

### Steps

see [STEPS.md](docs/STEPS.md)

## Docker Images

- [katom92/pod-rotator-operator:](https://hub.docker.com/r/katom92/pod-rotator-operator/tags)

## License

[Mozilla Public License 2.0](LICENSE)
