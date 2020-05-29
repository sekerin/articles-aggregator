package main

import (
	"github.com/sekerin/artticles-aggregator/cmd/parser"
	"github.com/sekerin/artticles-aggregator/cmd/search"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(parser.Cmd, search.Cmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
