package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// GET All
func GetDevices() string {
	url := UrlHTTP + "/devices"
	rsp, err := netClient.Get(url)
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, e := prettyJSON(body)
	if e != nil {
		return "ERROR JSON"
	}

	return string(b)
}

// GET
func GetDevice(id string) string {
	url := UrlHTTP + "/devices/" + id
	rsp, err := netClient.Get(url)
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, e := prettyJSON(body)
	if e != nil {
		return "ERROR JSON"
	}

	return string(b)
}

// PUT
func UpdateDevice(id string, msg string) string {
	var err error

	url := UrlHTTP + "/devices/" + id
	sr := strings.NewReader(msg)
	req, err := http.NewRequest("PUT", url, sr)
	if err != nil {
		return "ERROR"
	}

	req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(msg)))

	rsp, err := netClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, err := prettyJSON(body)
	if err != nil {
		return "ERROR JSON"
	}

	return string(b)
}

// POST
func CreateDevice(msg string) string {
	var err error

	url := UrlHTTP + "/devices"
	rsp, err := netClient.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, err := prettyJSON(body)
	if err != nil {
		return "ERROR JSON"
	}

	return string(b)
}

// DELETE
func DeleteDevice(id string) string {
	var err error

	url := UrlHTTP + "/devices/" + id
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return "ERROR"
	}
	rsp, err := netClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, err := prettyJSON(body)
	if err != nil {
		return "ERROR JSON"
	}

	return string(b)
}
