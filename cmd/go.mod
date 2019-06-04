module github.com/icereed/klabeler/cmd

go 1.12

require (
	github.com/Jeffail/gabs v1.4.0
	github.com/icereed/klabeler/internal/pkg/entities v0.0.0-00010101000000-000000000000
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.4
	github.com/spf13/viper v1.4.0
	gopkg.in/src-d/go-git.v4 v4.11.0
	k8s.io/apimachinery v0.0.0-20190602183612-63a6072eb563
	sigs.k8s.io/yaml v1.1.0
)

replace github.com/icereed/klabeler/internal/pkg/entities => ../internal/pkg/entities
