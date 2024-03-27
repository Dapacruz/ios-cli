package main

import (
	"github.com/Dapacruz/ios-cli/cmd"
	_ "github.com/Dapacruz/ios-cli/cmd/config"
	_ "github.com/Dapacruz/ios-cli/cmd/device"
)

func main() {
	cmd.Execute()
}
