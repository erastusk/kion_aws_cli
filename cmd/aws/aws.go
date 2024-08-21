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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// awsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// awsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
