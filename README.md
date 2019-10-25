# Kubernetes Admission Webhook for LXCFS

This project shows how to build and deploy an [AdmissionWebhook](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers) for [LXCFS](https://github.com/lxc/lxcfs).

## Prerequisites

Kubernetes 1.9.0 or above with the `admissionregistration.k8s.io/v1beta1` API enabled. Verify that by the following command:
```
kubectl api-versions | grep admissionregistration.k8s.io/v1beta1
```
The result should be:
```
admissionregistration.k8s.io/v1beta1
```

In addition, the `MutatingAdmissionWebhook` and `ValidatingAdmissionWebhook` admission controllers should be added and listed in the correct order in the admission-control flag of kube-apiserver.

## Build docker image

```
make container
```

## Deploy 
 
1. Deploy lxcfs

```
yum install automake autoconf libtool fuse fuse-libs fuse-devel
git clone https://github.com/lxc/lxcfs.git
cd lxcfs
./bootstrap.sh
make && make install
systemctl start lxcfs
```

1. Deply lxcfs-admission-webhook

```
kubectl apply -f deployment/lxcfs-admission-webhook.yaml
```

## Test

1. Deploy the test pod
 
```
kubectl apply -f deployment/pod.yaml
```

2. Inspect the resource inside container

```
$ kubectl exec -it lxcfs-pod bash
[root@lxcfs-pod /]# free -h
              total        used        free      shared  buff/cache   available
Mem:           1.0G        3.3M        1.0G          0B          0B        1.0G
Swap:            0B          0B          0B
```

## Cleanup

1. Uninstall lxcfs-admission-webhook

```
kubectl delete -f deployment/lxcfs-admission-webhook.yaml
kubectl delete mutatingwebhookconfigurations lxcfs-admission-webhook
```

2. Uninstall lxcfs from cluster nodes

```
systemctl stop lxcfs
cd lxcfs 
make uninstall
```
