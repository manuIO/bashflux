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
	"io/ioutil"
	"net/http"
	"strings"
)

// GetMsg - gets messages from the channel
func GetMsg(id string) string {

	url := UrlHTTP + "/channels/" + id + "/msg"
	resp, err := netClient.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	b, e := prettyJSON(body)
	if e != nil {
		return err.Error()
	}

	return string(b)
}

// SendMsg - publishes SenML message on the channel
func SendMsg(id string, msg string) string {
	var err error

	url := UrlHTTP + "/channels/" + id + "/msg"
	sr := strings.NewReader(msg)
	resp, err := netClient.Post(url, "application/json", sr)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	s := ""
	if resp.StatusCode == 202 {
		s = fmt.Sprintf("Message sent")
	} else {
		s = http.StatusText(resp.StatusCode)
	}

	return s
}
