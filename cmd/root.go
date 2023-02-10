/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/saddle-dev/saddle/cmd/version"
)

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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(version.VersionCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.saddle.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommandPalettes()
}
