// Package internal provides core functionality for interacting with the Kion API.
// This file contains HTTP client functions for authentication and credential retrieval.
package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// token holds the Bearer token for Kion API authentication.
// TODO: This should be moved to environment variables or a secure configuration file.
// Set this to your Kion API token before use.
const token = "Bearer <YOUR API TOKEN>"

// Getlabels retrieves detailed account information from Kion for a given account ID.
// This includes the account name, number, and other metadata.
// Parameters:
//   - id: The numeric account ID in Kion
// Returns:
//   - projectLabels: Struct containing account details
func Getlabels(id int) projectLabels {
	// Convert account ID to string for URL construction
	i := strconv.Itoa(id)
	url := fmt.Sprintf("https://login.mcp.nasa.gov/api/v3/account/%s", i)

	// Create HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	
	// Set required headers for Kion API
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", token)
	
	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	
	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var prolabels projectLabels
	err = json.Unmarshal(body, &prolabels)
	if err != nil {
		fmt.Println(err)
	}
	return prolabels
}

// GetRoles retrieves all cloud access roles available to the authenticated user.
// These roles represent the AWS accounts and IAM roles the user can assume.
// Returns:
//   - SessionRoles: Struct containing list of available roles and account information
func GetRoles() SessionRoles {
	url := "https://login.mcp.nasa.gov/api/v3/me/cloud-access-role"
	
	// Create HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	
	// Set required headers
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", token)
	
	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	
	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var sessroles SessionRoles
	err = json.Unmarshal(body, &sessroles)
	if err != nil {
		fmt.Println(err)
	}
	return sessroles
}

// GetKeys requests temporary AWS credentials from Kion for a specific account and role.
// These credentials are short-lived and provide secure access to AWS resources.
// Parameters:
//   - ac: AWS account number
//   - role: IAM role name to assume
// Returns:
//   - SessionKeys: Struct containing temporary AWS access key, secret key, and session token
func GetKeys(ac, role string) SessionKeys {
	url := "https://login.mcp.nasa.gov/api/v3/temporary-credentials"
	
	// Create request body with account and role information
	requestbody, err := json.Marshal(map[string]string{
		"account_number": ac,
		"iam_role_name":  role,
	})
	if err != nil {
		fmt.Println(err)
	}
	
	// Create HTTP client and POST request
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestbody))
	if err != nil {
		fmt.Println(err)
	}
	
	// Set required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", token)
	
	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	
	// Read and parse the response containing temporary credentials
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var sesskeys SessionKeys
	err = json.Unmarshal(body, &sesskeys)
	if err != nil {
		fmt.Println(err)
	}
	return sesskeys
}
