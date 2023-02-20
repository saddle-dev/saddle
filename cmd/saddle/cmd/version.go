/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints CLI version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.0.0-dev")
	},
}

func setupVersionCommand(cmd *cobra.Command) {
	cmd.AddCommand(VersionCmd)
}
