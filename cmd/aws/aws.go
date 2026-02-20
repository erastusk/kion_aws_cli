/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package aws

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AwsCmd represents the aws command group.
// This command serves as a parent for AWS-related subcommands.
var AwsCmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS access keys configuration",
	Long:  `Manage AWS access keys and credentials through Kion integration.

This command group provides functionality to authenticate with AWS accounts
and configure temporary credentials for secure access.`,
	Run: func(cmd *cobra.Command, args []string) {
		// When called without subcommands, display help information
		fmt.Println(cmd.Help())
	},
}

// init registers all subcommands under the aws command.
func init() {
	// Register the login subcommand
	AwsCmd.AddCommand(LoginCmd)
}
