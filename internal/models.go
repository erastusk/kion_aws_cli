// Package internal defines data models for Kion API responses and AWS credentials.
// These structs are used for JSON marshaling/unmarshaling of API responses.
package internal

import "time"

// SessionKeys represents the response from the Kion temporary credentials API.
// Contains the temporary AWS access credentials.
type SessionKeys struct {
	Status int      `json:"status"` // HTTP status code of the response
	Data   dataKeys `json:"data"`   // The actual credential data
}

// dataKeys contains the temporary AWS credentials returned by Kion.
type dataKeys struct {
	AccessKey       string `json:"access_key"`        // AWS access key ID
	SecretAccessKey string `json:"secret_access_key"` // AWS secret access key
	SessionToken    string `json:"session_token"`     // AWS session token for temporary credentials
}

// SessionRoles represents the response from the Kion cloud access roles API.
// Contains all AWS roles available to the authenticated user.
type SessionRoles struct {
	Status int         `json:"status"` // HTTP status code of the response
	Data   []dataRoles `json:"data"`   // Array of available cloud access roles
}

// dataRoles represents a single cloud access role configuration in Kion.
// Contains details about an AWS account and the IAM role that can be assumed.
type dataRoles struct {
	AccountID           int       `json:"account_id"`            // Internal Kion account ID
	AccountNumber       string    `json:"account_number"`        // AWS account number (12 digits)
	AccountType         string    `json:"account_type"`          // Type of AWS account
	ApplyToAllAccounts  bool      `json:"apply_to_all_accounts"` // Whether role applies to all accounts
	AwsIamPath          string    `json:"aws_iam_path"`          // IAM path for the role
	AwsIamRoleName      string    `json:"aws_iam_role_name"`     // Name of the IAM role to assume
	CloudAccessRoleType string    `json:"cloud_access_role_type"`// Type of cloud access role
	CreatedAt           createdAt `json:"created_at"`            // Role creation timestamp
	DeletedAt           deletedAt `json:"deleted_at"`            // Role deletion timestamp (if deleted)
	FutureAccounts      bool      `json:"future_accounts"`       // Apply to future accounts
	ID                  int       `json:"id"`                    // Unique role identifier
	LongTermAccessKeys  bool      `json:"long_term_access_keys"` // Whether long-term keys are enabled
	Name                string    `json:"name"`                  // Display name for the role
	ProjectID           int       `json:"project_id"`            // Associated Kion project ID
	ShortTermAccessKeys bool      `json:"short_term_access_keys"`// Whether short-term keys are enabled
	UpdatedAt           updatedAt `json:"updated_at"`            // Last update timestamp
	WebAccess           bool      `json:"web_access"`            // Whether web console access is enabled
}

// createdAt represents a nullable timestamp for when a resource was created.
type createdAt struct {
	Time  time.Time `json:"Time"`  // The actual timestamp value
	Valid bool      `json:"Valid"` // Whether the timestamp is valid/set
}

// deletedAt represents a nullable timestamp for when a resource was deleted.
type deletedAt struct {
	Time  time.Time `json:"Time"`  // The actual timestamp value
	Valid bool      `json:"Valid"` // Whether the timestamp is valid/set
}

// updatedAt represents a nullable timestamp for when a resource was last updated.
type updatedAt struct {
	Time  time.Time `json:"Time"`  // The actual timestamp value
	Valid bool      `json:"Valid"` // Whether the timestamp is valid/set
}

// userInfo is a flexible map for storing user-related information.
// Currently unused but available for future enhancements.
type userInfo map[string]string

// projectLabels represents the response from the Kion account details API.
// Contains detailed information about a specific AWS account.
type projectLabels struct {
	Status int               `json:"status"` // HTTP status code of the response
	Data   projectlablesData `json:"data"`   // Account details
}

// projectlablesData contains comprehensive metadata about an AWS account in Kion.
type projectlablesData struct {
	AccountEmail              string `json:"account_email"`               // Email associated with the account
	AccountName               string `json:"account_name"`                // Human-readable account name
	AccountNumber             string `json:"account_number"`              // AWS account number
	AccountTypeID             int    `json:"account_type_id"`             // Account type identifier
	CarExternalID             string `json:"car_external_id"`             // Cloud Access Role external ID
	CreatedAt                 string `json:"created_at"`                  // Account creation timestamp
	DeletedAt                 string `json:"deleted_at"`                  // Account deletion timestamp (if deleted)
	ID                        int    `json:"id"`                          // Internal Kion account ID
	IncludeLinkedAccountSpend bool   `json:"include_linked_account_spend"`// Include linked account costs
	LinkedAccountNumber       string `json:"linked_account_number"`       // Linked AWS account number
	LinkedRole                string `json:"linked_role"`                 // Linked IAM role
	PayerID                   int    `json:"payer_id"`                    // Payer account ID for billing
	ProjectID                 int    `json:"project_id"`                  // Associated Kion project ID
	ServiceExternalID         string `json:"service_external_id"`         // Service external ID
	SkipAccessChecking        bool   `json:"skip_access_checking"`        // Whether to skip access validation
	StartDatecode             string `json:"start_datecode"`              // Account start date
	UseOrgAccountInfo         bool   `json:"use_org_account_info"`        // Use organization account info
}

// AccessKeys is an alternative structure for AWS credentials.
// Currently unused in favor of dataKeys structure.
type AccessKeys struct {
	Access_key_id     string `json:"access_key_id"`     // AWS access key ID
	Secret_access_key string `json:"secret_access_key"` // AWS secret access key
	Session_token     string `json:"session_token"`     // AWS session token
}
