/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package cmd

import (
	"io/ioutil"
)

// Status - server health check
func Status() string {
	url := UrlHTTP + "/status"
	rsp, err := netClient.Get(url)
	if err != nil {
		return err.Error()
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	return GetPrettyJson(body)
}
