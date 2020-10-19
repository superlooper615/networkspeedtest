package main

import (
	"os"

	"github.com/superlooper615/networkspeedtest/cmd"

	_ "github.com/superlooper615/networkspeedtest/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
