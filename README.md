<p align="center">
  <img src="logo.png" alt="eksup"/>
</p>

eksup is a tool to help DevOps Engineers/Platform Engineers/SREs (you name it!) upgrade EKS clusters.

## What is it?

With eksup you can list and check for newer versions of addons and Kubernetes versions supported by EKS.

## Getting started

- Download the latest release from the [releases page](https://github.com/pequi-io/eksup/releases) and extract the binary.
- Run `./eksup init` to create a new config file and then configure it with your AWS credentials or SSO profile name. You can't use both at the same time.

### Listing addons

```bash
$ eksup addons list

Listing installed add-ons for cluster: my-cluster
coredns
kube-proxy
vpc-cni
```

### Check for newer addons versions

```bash
$ eksup addons check

Listing add-ons for cluster: my-cluster
coredns is running version: v1.9.3-eksbuild.2 and can be upgraded to version: v1.9.3-eksbuild.5
kube-proxy is running version: v1.24.9-eksbuild.1 and can be upgraded to version: v1.24.10-eksbuild.2
vpc-cni is running version: v1.12.5-eksbuild.1 and can be upgraded to version: v1.13.0-eksbuild.1
```

### List EKS clusters

```bash
$ eksup addons check

List of clusters:
my-cluster
my-cluster-02
```

### Check for newer Kubernetes versions

```bash
$ eksup addons check

List of clusters:
my-cluster is running version: 1.24 and can be upgraded to version: 1.27
my-cluster-02 is running version: 1.24 and can be upgraded to version: 1.27
```

## What it does not do

Currently, eksup does not support any upgrade operations. It only lists the available versions for addons and Kubernetes. 

## Contributing

Contributions are welcome! Please check the [CONTRIBUTING.md](CONTRIBUTING.md) file for more information.
