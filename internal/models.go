package internal

import "time"

type SessionKeys struct {
	Status int      `json:"status"`
	Data   dataKeys `json:"data"`
}
type dataKeys struct {
	AccessKey       string `json:"access_key"`
	SecretAccessKey string `json:"secret_access_key"`
	SessionToken    string `json:"session_token"`
}
type SessionRoles struct {
	Status int         `json:"status"`
	Data   []dataRoles `json:"data"`
}
type dataRoles struct {
	AccountID           int       `json:"account_id"`
	AccountNumber       string    `json:"account_number"`
	AccountType         string    `json:"account_type"`
	ApplyToAllAccounts  bool      `json:"apply_to_all_accounts"`
	AwsIamPath          string    `json:"aws_iam_path"`
	AwsIamRoleName      string    `json:"aws_iam_role_name"`
	CloudAccessRoleType string    `json:"cloud_access_role_type"`
	CreatedAt           createdAt `json:"created_at"`
	DeletedAt           deletedAt `json:"deleted_at"`
	FutureAccounts      bool      `json:"future_accounts"`
	ID                  int       `json:"id"`
	LongTermAccessKeys  bool      `json:"long_term_access_keys"`
	Name                string    `json:"name"`
	ProjectID           int       `json:"project_id"`
	ShortTermAccessKeys bool      `json:"short_term_access_keys"`
	UpdatedAt           updatedAt `json:"updated_at"`
	WebAccess           bool      `json:"web_access"`
}
type createdAt struct {
	Time  time.Time `json:"Time"`
	Valid bool      `json:"Valid"`
}
type deletedAt struct {
	Time  time.Time `json:"Time"`
	Valid bool      `json:"Valid"`
}
type updatedAt struct {
	Time  time.Time `json:"Time"`
	Valid bool      `json:"Valid"`
}

type userInfo map[string]string

type projectLabels struct {
	Status int               `json:"status"`
	Data   projectlablesData `json:"data"`
}
type projectlablesData struct {
	AccountEmail              string `json:"account_email"`
	AccountName               string `json:"account_name"`
	AccountNumber             string `json:"account_number"`
	AccountTypeID             int    `json:"account_type_id"`
	CarExternalID             string `json:"car_external_id"`
	CreatedAt                 string `json:"created_at"`
	DeletedAt                 string `json:"deleted_at"`
	ID                        int    `json:"id"`
	IncludeLinkedAccountSpend bool   `json:"include_linked_account_spend"`
	LinkedAccountNumber       string `json:"linked_account_number"`
	LinkedRole                string `json:"linked_role"`
	PayerID                   int    `json:"payer_id"`
	ProjectID                 int    `json:"project_id"`
	ServiceExternalID         string `json:"service_external_id"`
	SkipAccessChecking        bool   `json:"skip_access_checking"`
	StartDatecode             string `json:"start_datecode"`
	UseOrgAccountInfo         bool   `json:"use_org_account_info"`
}
type AccessKeys struct {
	Access_key_id     string `json:"access_key_id"`
	Secret_access_key string `json:"secret_access_key"`
	Session_token     string `json:"session_token"`
}
