package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Sync CUE definitions to application configurations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("who up")
	},
}

func setupUpCommand(cmd *cobra.Command) {
	cmd.AddCommand(upCmd)
}
