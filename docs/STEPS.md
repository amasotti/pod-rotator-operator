# Steps for a custom operator

1. Create a new directory and cd inside

2. Initialize a new operator project using the Operator SDK

```sh
operator-sdk init --domain=example.com --repo=github.com/example/pod-rotator-operator
```

This command generates quite a bunch of files and directories. The most important ones are:
- `Dockerfile` - the Dockerfile for the operator
- `Makefile` - the Makefile for building and deploying the operator
- `config/` - the directory containing the CRD and RBAC manifests
- `bin/` - the directory containing the operator binary
- `cmd/` - the directory containing the main.go file
- `internal/` - the directory containing the controller logic

3. Create a new API for the operator

```sh
operator-sdk create api --group=apps --version=v1 --kind=CustomPodRotator --resource=true --controller=true
```

The comand above generator adds the following files:

- `api/v1/custompodrotator_types.go` - the API type definition
- `controllers/custompodrotator_controller.go` - the controller logic
- `controllers/suite_test.go` - the test suite for the controller
- `controllers/custompodrotator_controller_test.go` - the controller tests
- `config/crd/bases/apps.example.com_custompodrotators.yaml` - the CRD manifest
- `config/samples/apps_v1_custompodrotator.yaml` - the sample CR manifest
- `config/rbac/role.yaml` - the RBAC role manifest


## Important steps in the implementation

1. Implement the reconcile loop in the controller

The central place where you will be spending most of the time is the controller logic. Under `internal/controllers` the 
operator-sdk has scaffolded a controller file for you. The main function in this file is the `Reconcile` function. 

Reconciliation is a key operation in Kubernetes controllers. It is the process of bringing the current state of the
cluster closer to the desired state. In the context of the operator, the desired state is defined by the Custom Resource
Definition (CRD) that you have created.

See [custompodrotator_controller.go (line 49ff)](../internal/controller/custompodrotator_controller.go) for the implemented logic.


2. Run the tests

The operator-sdk has scaffolded a test suite for you. 
Implement your own tests and run them using the following command:

```sh
make test
```

3. Build and push

Build the operator image and push it to a container registry:

```sh
make docker-build docker-push IMG=<registry>/<your-image-name>:<tag>
```

## A couple of words on some standard make recipes

- `make install/uninstall`: These handle the CRD installation/removal in your cluster. They only manage the Custom Resource Definitions.
- `make deploy/undeploy`: These handle the full operator deployment, including the CRDs, RBAC rules, and the controller deployment itself.