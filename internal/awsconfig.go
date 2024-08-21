package internal

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func getSeperator() string {
	oss := runtime.GOOS
	switch oss {
	case "windows":
		return "\\"
	case "linux":
		return "/"
	default:
		log.Fatal("Could not determine os")
		return ""
	}
}
func GetFile() []byte {
	hp, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	sep := getSeperator()
	homePath := fmt.Sprintf("%s%s.aws%scredentials", hp, sep, sep)
	input, err := os.ReadFile(homePath)
	if err != nil {
		log.Fatalln(err)
	}
	return input
}

func UpdateFile(input []byte, profile string, a SessionKeys) error {
	a_key_id := fmt.Sprintf("aws_access_key_id = %s", a.Data.AccessKey)
	sec_key := fmt.Sprintf("aws_secret_access_key = %s", a.Data.SecretAccessKey)
	sess_tok := fmt.Sprintf("aws_session_token = %s", a.Data.SessionToken)
	fmt.Printf("%s\n%s\n%s\n", a_key_id, sec_key, sess_tok)
	return nil
}
