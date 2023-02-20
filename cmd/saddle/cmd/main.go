/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// Main adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Main() int {
	setupRootCommand()

	err := rootCmd.Execute()
	if err != nil {
		return 1
	}

	return 0
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "saddle",
	Short: "Code scaffolding for the long haul",
	Long: `saddle manages application configuration in a single language, CUE.
Users supply configuration data provided by libraries and validated against
schemas. saddle will then generate the files in the correct format.

saddle makes it simple to manage configuration for multiple tools and apply
them consistently across one to many repositories.

For more information on using saddle to manage code see saddle.dev.`,
}

func addSubCommandPalettes() {
	setupVersionCommand(rootCmd)
}

func setupRootCommand() {
	addSubCommandPalettes()
}
