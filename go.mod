module github.com/icereed/klabeler

go 1.12

require (
	github.com/icereed/klabeler/cmd v0.0.0-20190603201259-d36e0d82a881 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.4
	github.com/spf13/viper v1.4.0
	klabeler/cmd v0.0.0-00010101000000-000000000000
)

replace klabeler/cmd => ./cmd
