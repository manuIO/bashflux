/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package cmd

import (
	"fmt"
	"strings"
)

// CreateUser - create user
func CreateUser(user string, pwd string) string {
	msg := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, user, pwd)
	sr := strings.NewReader(msg)
	url := UrlHTTP + "/users"
	resp, err := netClient.Post(url, "application/json", sr)
	s := PrettyHttpResp(resp, err)

	return s
}

// LogginInUser - loggin in user
func CreateToken(user string, pwd string) string {
	var err error

	msg := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, user, pwd)
	sr := strings.NewReader(msg)
	url := UrlHTTP + "/tokens"
	resp, err := netClient.Post(url, "application/json", sr)
	s := PrettyHttpResp(resp, err)

	return s
}
