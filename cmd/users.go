package cmd

import (
	"fmt"
	"strings"
)

// CreateUser - create user
func CreateUser(user string, pwd string) string {
	msg := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, user, pwd)
	url := serverAddr + "/users"
	resp, err := netClient.Post(url, "application/json", strings.NewReader(msg))
	s := PrettyHTTPResp(resp, err)

	return s
}

// CreateToken - create user token
func CreateToken(user string, pwd string) string {
	msg := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, user, pwd)
	url := serverAddr + "/tokens"
	resp, err := netClient.Post(url, "application/json", strings.NewReader(msg))
	s := PrettyHTTPResp(resp, err)

	return s
}
