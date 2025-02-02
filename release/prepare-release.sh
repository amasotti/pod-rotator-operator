export IMG=katom92/pod-rotator:v0.0.3

# Generate the complete manifest
make manifests
cd config/manager && kustomize edit set image controller=$IMG && cd ../..
kustomize build config/default > release/release.yaml