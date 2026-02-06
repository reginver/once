package main

import (
	"os"

	"github.com/basecamp/once/internal/command"
)

func main() {
	if err := command.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
