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

## Architecture

> This app tries to follow the [Clean Architecture guidelines](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

### Entities

Because the label logic with JSON is the main _business logic_ here, the code is located in `internal/pkg/entities`.