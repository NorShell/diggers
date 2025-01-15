package cmd

import (
	"github.com/NorShell/diggers/runtime"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a js script",
	Run: func(cmd *cobra.Command, args []string) {
		runtime.Execute()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
