// Package internal provides core functionality for AWS credentials management.
// This file handles reading and writing AWS credentials to the local configuration file.
package internal

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// getSeperator returns the appropriate file path separator for the current OS.
// Returns "\\" for Windows and "/" for Unix-based systems (Linux, macOS).
func getSeperator() string {
	oss := runtime.GOOS
	switch oss {
	case "windows":
		return "\\"
	case "linux", "darwin": // darwin is macOS
		return "/"
	default:
		log.Fatal("Could not determine os")
		return ""
	}
}

// GetFile reads the AWS credentials file from the user's home directory.
// Returns the file contents as a byte slice.
// The credentials file is typically located at ~/.aws/credentials
func GetFile() []byte {
	// Get the user's home directory path
	hp, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	
	// Construct the path to AWS credentials file
	sep := getSeperator()
	homePath := fmt.Sprintf("%s%s.aws%scredentials", hp, sep, sep)
	
	// Read the credentials file
	input, err := os.ReadFile(homePath)
	if err != nil {
		log.Fatalln(err)
	}
	return input
}

// UpdateFile updates the AWS credentials file with new temporary credentials.
// Parameters:
//   - input: Current contents of the credentials file (currently unused)
//   - profile: The AWS profile name to update
//   - a: SessionKeys containing the new AWS access credentials
//
// TODO: This function currently only prints credentials to stdout.
// It should be implemented to actually write to the credentials file.
func UpdateFile(input []byte, profile string, a SessionKeys) error {
	// Format the credential key-value pairs
	a_key_id := fmt.Sprintf("aws_access_key_id = %s", a.Data.AccessKey)
	sec_key := fmt.Sprintf("aws_secret_access_key = %s", a.Data.SecretAccessKey)
	sess_tok := fmt.Sprintf("aws_session_token = %s", a.Data.SessionToken)
	
	// Print credentials (TODO: Write to file instead)
	fmt.Printf("%s\n%s\n%s\n", a_key_id, sec_key, sess_tok)
	
	// TODO: Implement actual file writing logic
	// Should update or append to the credentials file under [profile] section
	return nil
}
