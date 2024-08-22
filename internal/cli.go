package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const token = "Bearer <YOUR API TOKEN>"

func Getlabels(id int) projectLabels {
	i := strconv.Itoa(id)
	url := fmt.Sprintf("https://login.mcp.nasa.gov/api/v3/account/%s", i)

	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
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
func GetRoles() SessionRoles {

	url := "https://login.mcp.nasa.gov/api/v3/me/cloud-access-role"
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
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
func GetKeys(ac, role string) SessionKeys {

	url := "https://login.mcp.nasa.gov/api/v3/temporary-credentials"
	requestbody, err := json.Marshal(map[string]string{
		"account_number": ac,
		"iam_role_name":  role,
	})
	if err != nil {
		fmt.Println(err)
	}
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestbody))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
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
