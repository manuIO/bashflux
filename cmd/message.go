/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package cmd

import (
	"strings"
	"net/http"
)

// GetMsg - gets messages from the channel
func GetMsg(id string, startTime string, endTime string) string {
	url := UrlHTTP + "/channels/" + id + "/messages" +
	                 "?start_time=" + startTime + "&end_time=" + endTime
	resp, err := netClient.Get(url)
	s := PrettyHttpResp(resp, err)

	return s
}

// SendMsg - publishes SenML message on the channel
func SendMsg(id string, msg string, token string) string {
	url := UrlHTTP + "/channels/" + id + "/messages"
	req, err := http.NewRequest("POST", url, strings.NewReader(msg))
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHttpResp(resp, err)

	return s
}
