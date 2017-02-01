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
	"net/http"
	"strings"
)

// CreateChannel - creates new channel and generates UUID
func CreateUser(user string, pwd string) string {
	var err error
	msg := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, user, pwd)

	sr := strings.NewReader(msg)

	url := UrlHTTP + "/users"
	resp, err := netClient.Post(url, "application/json", sr)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	var s = fmt.Sprintf("POST/users: Status %d\n", resp.StatusCode)
	if resp.StatusCode == 201 {
		body := GetHttpRespBody(resp, err)
		s = fmt.Sprintf("%s %s", s, GetPrettyJson(body))
	} else {
		s = fmt.Sprintf("%s %s", s, http.StatusText(resp.StatusCode))
	}

	return s
}

// CreateChannel - creates new channel and generates UUID
func InitSession(user string, pwd string) string {
	var err error
	msg := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, user, pwd)

	sr := strings.NewReader(msg)

	url := UrlHTTP + "/sessions"
	resp, err := netClient.Post(url, "application/json", sr)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	var s = fmt.Sprintf("POST/sessions: Status %d\n", resp.StatusCode)
	if resp.StatusCode == 201 {
		body := GetHttpRespBody(resp, err)
		s = fmt.Sprintf("%s %s", s, GetPrettyJson(body))
	} else {
		s = fmt.Sprintf("%s %s", s, http.StatusText(resp.StatusCode))
	}

	return s
}
