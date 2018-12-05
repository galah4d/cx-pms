package main

import (
	"fmt"
	"github.com/galah4d/cx-pms/cli/commands"
	"os"
)

func main() {
	if _, ok := os.LookupEnv("GOPATH"); !ok {
		fmt.Println("Error: GOPATH not set, exiting...")
		os.Exit(0)
	}
	commands.Execute()
}
