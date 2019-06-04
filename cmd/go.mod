module github.com/icereed/klabeler/cmd

go 1.12

require (
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/icereed/klabeler/pkg/app v0.0.0-00010101000000-000000000000
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.4
	github.com/spf13/viper v1.4.0
	golang.org/x/sys v0.0.0-20190312061237-fead79001313 // indirect
	golang.org/x/text v0.3.1-0.20181227161524-e6919f6577db // indirect
)

replace github.com/icereed/klabeler/internal/pkg/entities => ../internal/pkg/entities

replace github.com/icereed/klabeler/pkg/app => ../pkg/app
