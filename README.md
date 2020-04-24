# Firewall Controller


## Initial Setup

1. download kubebuilder
1. download kustomize from https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv3.5.4/kustomize_v3.5.4_linux_amd64.tar.gz
1. init project and run kubebuilder

```bash
kubebuilder init --domain metal-stack.io
kubebuilder create api --group firewall --version v1 --kind Network
```

1. run test

```bash
export KUBEBUILDER_ASSETS=~/dev/kubebuilder_2.3.1_linux_amd64/bin
make test
```

## Testing locally

```bash
# start kind cluster
kind create cluster

# start node-exporter
node_exporter

# deploy manifests and sample crd
k apply -f config/crd/bases
k apply -f config/samples

# start the controller
bin/firewall-controller

# watch results
k describe networktraffic
```
