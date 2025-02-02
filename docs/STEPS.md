# Steps for a custom operator

1. Create a new directory and cd inside

2. Initialize a new operator project using the Operator SDK

```sh
operator-sdk init --domain=example.com --repo=github.com/example/pod-rotator-operator
```

3. Create a new API for the operator

```sh
operator-sdk create api --group=apps --version=v1 --kind=PodRotator --resource=true --controller=true
```