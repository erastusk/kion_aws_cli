/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package aws

import (
	"brcp/internal"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// LoginCmd represents the login command.
// This command allows users to interactively select an AWS role and configure credentials.
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Choose AWS login profile",
	Long:  `Interactively select an AWS account and role to configure credentials.

This command fetches available AWS roles from Kion, presents them in an
interactive menu, and configures your local AWS credentials file with
temporary access keys for the selected role.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Step 1: Fetch and format available AWS roles
		z := GetUserData()
		
		// Step 2: Present interactive selection menu
		a := GetUserPrompt(z)
		
		// Step 3: Parse the selected option (format: "AccountNumber RoleName ProfileName")
		acc := strings.Split(a, " ")[0]  // Account number
		r := strings.Split(a, " ")[1]    // IAM role name
		pro := strings.Split(a, " ")[2]  // Profile/account name
		
		// Step 4: Retrieve temporary AWS credentials from Kion
		s := internal.GetKeys(acc, r)
		
		// Step 5: Read existing AWS credentials file
		b := internal.GetFile()
		
		// Step 6: Update credentials file with new keys
		err := internal.UpdateFile(b, pro, s)
		if err != nil {
			fmt.Println(err)
		}
	},
}

// init initializes the login command configuration.
func init() {
	// Future: Add flags for non-interactive mode, profile selection, etc.
}

// GetUserData fetches all available AWS roles from Kion and formats them for display.
// Returns a slice of strings, each formatted as "AccountNumber RoleName AccountName".
func GetUserData() []string {
	// Fetch all cloud access roles for the current user
	r := internal.GetRoles()
	s := []string{}
	
	// Format each role with account details
	for i := range r.Data {
		// Get human-readable account name using the account ID
		z := internal.Getlabels(r.Data[i].AccountID)
		
		// Format: "AccountNumber IAMRoleName AccountName"
		str := fmt.Sprintf("%s %s %s", r.Data[i].AccountNumber, r.Data[i].AwsIamRoleName, z.Data.AccountName)
		s = append(s, str)
	}
	return s
}

// GetUserPrompt displays an interactive menu for role selection.
// Takes a slice of formatted role strings and returns the user's selection.
func GetUserPrompt(z []string) string {
	var role string
	
	// Configure the interactive survey question
	qs := []*survey.Question{
		{
			Name: "AWS Roles",
			Prompt: &survey.Select{
				Message: "Select AWS Role:",
				Options: z,
			},
		},
	}
	
	// Present the interactive menu and capture user selection
	err := survey.Ask(qs, &role)
	if err != nil {
		fmt.Println(err)
	}
	return role
}
