package main

import (
	"github.com/alexsuslov/cli"
	"os"
)

func main() {
	if err := cli.
		New("cli", "Example Cli").
		Action(os.Args); err != nil {
		panic(err)
	}
}
