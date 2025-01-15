package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Digger",
	Short: "Digger is an experimental JS runtime.",
	Long:  `A minimal JS runtime built on top of V8GO, with a focus on creating bindings`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
