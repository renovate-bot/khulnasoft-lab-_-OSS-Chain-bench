package main

import (
	"os"

	"github.com/khulnasoft-lab/oss-chain-bench/internal/commands"
)

var version = "dev"

func main() {
	if err := commands.Execute(version); err != nil {
		os.Exit(1)
	}
}
