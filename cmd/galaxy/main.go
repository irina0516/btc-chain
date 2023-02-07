package main

import (
	"fmt"
	"os"

	"github.com/galaxy-digital/galaxychain-chain/cmd/galaxy/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
