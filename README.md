# klabeler

Apply current git hash to Kubernetes resources.

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

## Good to read

As this is also a project to teach myself golang best practices, I will dump some useful resources here.

- https://blog.chewxy.com/2018/03/18/golang-interfaces/
- https://medium.com/@cep21/aspects-of-a-good-go-library-7082beabb403
- https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
  