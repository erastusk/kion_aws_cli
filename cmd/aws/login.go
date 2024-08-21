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

// loginCmd represents the login command
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Choose AWS login Profile",
	Long:  "Choose AWS login Profile",
	Run: func(cmd *cobra.Command, args []string) {
		z := GetUserData()
		a := GetUserPrompt(z)
		acc := strings.Split(a, " ")[0]
		r := strings.Split(a, " ")[1]
		pro := strings.Split(a, " ")[2]
		s := internal.GetKeys(acc, r)
		b := internal.GetFile()
		err := internal.UpdateFile(b, pro, s)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func GetUserData() []string {
	r := internal.GetRoles()
	s := []string{}
	for i := range r.Data {
		z := internal.Getlabels(r.Data[i].AccountID)
		str := fmt.Sprintf("%s %s %s", r.Data[i].AccountNumber, r.Data[i].AwsIamRoleName, z.Data.AccountName)
		s = append(s, str)
	}
	return s
}

func GetUserPrompt(z []string) string {

	var role string
	qs := []*survey.Question{
		{
			Name: "AWS Roles",
			Prompt: &survey.Select{
				Message: "Select AWS Role:",
				Options: z,
			},
		},
	}
	err := survey.Ask(qs, &role)
	if err != nil {
		fmt.Println(err)
	}
	return role
}
