/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package aws

import (
	"fmt"

	"github.com/spf13/cobra"
)

// awsCmd represents the aws command
var AwsCmd = &cobra.Command{
	Use:   "aws",
	Short: "aws access keys configuration",
	Long:  "aws access keys configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())

	},
}

func init() {
	AwsCmd.AddCommand(LoginCmd)
}
