# klabeler

> Still only a POC. Don't use it in production or even staging!

Apply labels to k8s resources.

## Install

```shell
go get github.com/icereed/klabeler
```

## Usage

```shell
cat deployment.yml | klabeler | kubectl apply -f -
```
