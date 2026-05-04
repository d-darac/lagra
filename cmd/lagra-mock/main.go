package main

import (
	"os"

	"github.com/d-darac/lagra/cmd/lagra-mock/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
