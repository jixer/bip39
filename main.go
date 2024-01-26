package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("bip39", "Command line utility written in Go for generating and parsing bip-39 seed phrases")

	// Create subparser
	help := parser.Flag("h", "help", &argparse.Options{Required: false, Help: "Show help information for this command"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *help {
		fmt.Print(parser.Usage(nil))
	}
}
