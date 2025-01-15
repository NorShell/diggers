package cmd

import (
	"fmt"

	"github.com/NorShell/diggers/runtime"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [file]",
	Short: "Execute a JS script",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		fmt.Printf(" Running %s \n", file)
		runtime.Execute(file)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
