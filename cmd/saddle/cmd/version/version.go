/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package version

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

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
