module github.com/icereed/klabeler/internal/pkg/app

go 1.12

require (
	github.com/Jeffail/gabs v1.4.0
	github.com/ghodss/yaml v1.0.0
	github.com/icereed/klabeler/internal/pkg/entities v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.3.0
	gopkg.in/src-d/go-git.v4 v4.11.0
	gopkg.in/yaml.v2 v2.2.2 // indirect
	sigs.k8s.io/yaml v1.1.0
)

replace github.com/icereed/klabeler/internal/pkg/entities => ../entities
