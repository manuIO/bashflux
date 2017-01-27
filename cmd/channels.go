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
	"strconv"
	"strings"
	"encoding/json"

	"github.com/mainflux/mainflux-core/models"
	"github.com/hokaccha/go-prettyjson"
)

// CreateChannel - creates new channel and generates UUID
func CreateChannel(msg string) string {
	var err error

	url := UrlHTTP + "/channels"
	resp, err := netClient.Post(url, "application/json", nil)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	s := ""
	if resp.StatusCode == 201 {
		s = fmt.Sprintf("Created resource %s", resp.Header.Get("Location"))
	} else {
		s = http.StatusText(resp.StatusCode)
	}

	return s
}

// GetChannels - gets all channels
func GetChannels() string {
	url := UrlHTTP + "/channels"
	resp, err := netClient.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	b, e := prettyjson.Format([]byte(body))
	if e != nil {
		return err.Error()
	}

	return string(b)
}

// GetChannel - gets channel by ID
func GetChannel(id string) string {
	url := UrlHTTP + "/channels/" + id
	resp, err := netClient.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	b, e := prettyjson.Format([]byte(body))
	if e != nil {
		return err.Error()
	}

	return string(b)
}

// UpdateChannel - publishes SenML message on the channel
func UpdateChannel(id string, msg string) string {
	var err error

	url := UrlHTTP + "/channels/" + id
	sr := strings.NewReader(msg)
	req, err := http.NewRequest("PUT", url, sr)
	if err != nil {
		return err.Error()
	}

	req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(msg)))

	resp, err := netClient.Do(req)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	b, e := prettyjson.Format([]byte(body))
	if e != nil {
		return err.Error()
	}

	return string(b)
}

// DeleteChannel - removes channel
func DeleteChannel(id string) string {
	var err error

	url := UrlHTTP + "/channels/" + id
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err.Error()
	}
	resp, err := netClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	b, err := prettyjson.Format([]byte(body))
	if err != nil {
		return err.Error()
	}

	return string(b)
}

// DeleteAllChannels - removes all channels
func DeleteAllChannels() string {
	var err error

	url := UrlHTTP + "/channels"
	rsp, err := netClient.Get(url)
	if err != nil {
		return err.Error()
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	var channels []models.Channel
	json.Unmarshal(body, &channels)
	s := ""
	for i := 0; i < len(channels); i++ {
		s = s + DeleteChannel(channels[i].ID)
	}

	return s
}

// PlugChannel - plugs list of devices into the channel
func PlugChannel(id string, devices string) string {
	var err error

	url := UrlHTTP + "/channels/" + id + "/plug"
	sr := strings.NewReader(devices)
	resp, err := netClient.Post(url, "application/json", sr)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	b, err := prettyjson.Format([]byte(body))
	if err != nil {
		return err.Error()
	}

	return string(b)
}
