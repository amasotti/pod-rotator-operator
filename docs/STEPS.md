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
