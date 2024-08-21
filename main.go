package main

import "brcp/cmd"

func main() {
	cmd.Execute()
}

// func main() {
// 	r := internal.GetRoles()
// 	for i := range r.Data {
// 		z := internal.Getlabels(r.Data[i].AccountID)
// 		fmt.Println(r.Data[i].AccountNumber, r.Data[i].AwsIamRoleName, z.Data.AccountName)
// 	}
// 	// k := GetKeys("874873015118", "mcp-tenantDeveloper")
// 	// _ = r
// 	// _ = k

// }
