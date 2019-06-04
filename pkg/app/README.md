# KLabeler
```go
import "github.com/icereed/klabeler/pkg/app"
```

## Usage

#### type KLabeler

```go
type KLabeler interface {
	ApplyCurrentGitHash() KLabeler
	ApplyLabel(key string, value string) KLabeler
	SetLabelPrefix(prefix string) KLabeler
	GetJSON() string
	GetYAML() string
}
```

KLabeler represents the main use-cases for the app

#### func  NewKLabeler

```go
func NewKLabeler(jsonOrYaml string) (KLabeler, error)
```
NewKLabeler creates a new instance from JSON or YAML input

#### func  NewKLabelerWithGitHashProvider

```go
func NewKLabelerWithGitHashProvider(jsonOrYaml string, gitHashProvider GitHashProvider) (KLabeler, error)
```
NewKLabelerWithGitHashProvider creates a new instance from JSON or YAML input
with the given GitHashProvider

#### type GitHashProvider

```go
type GitHashProvider interface {
	// contains filtered or unexported methods
}
```
---

> README generated with [godocdown](https://github.com/robertkrimen/godocdown).
