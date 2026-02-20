/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"brcp/cmd/aws"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
// This is the entry point for all CLI operations.
var rootCmd = &cobra.Command{
	Use:   "brcp",
	Short: "Broadridge Cloud Platform AWS CLI Tool",
	Long: `BRCP is a command-line interface tool for managing AWS credentials through Kion.

This application simplifies the process of obtaining temporary AWS credentials
by integrating with the Kion cloud access management platform. It allows users
to select from available AWS roles and automatically configures their local
AWS credentials file.

Example usage:
  brcp aws login    # Interactive login to select and configure AWS credentials`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init initializes the root command by adding subcommands and defining flags.
func init() {
	// Register the aws subcommand
	rootCmd.AddCommand(aws.AwsCmd)
	
	// Define persistent flags (available to all subcommands)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
