# Release Pod-Rotator-Operator

This project heavily relies on Operator SDK and the Operator Framework.
The manifest for release is created via `kustomize` with a simple step.

After the project has been tested and the image built and pushed (see [docs/STEPS.md](./STEPS.md)),
just run the following command to create the release manifest:

```sh
./release/prepare-release.sh
```

which will generate a `release.yaml` file in the `release` directory.
This is ready to be applied to the cluster:

```sh
kubectl apply -f release.yaml
```