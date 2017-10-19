package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/mainflux/mainflux/manager"
)

var endPointC = "/channels"

// CreateChannel - creates new channel and generates UUID
func CreateChannel(msg string, token string) string {
	url := serverAddr + endPointC
	req, err := http.NewRequest("POST", url, strings.NewReader(msg))
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHTTPResp(resp, err)

	return s
}

// GetChannels - gets all channels
func GetChannels(limit int, token string) string {
	url := serverAddr + "/channels?climit=" + strconv.Itoa(limit)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHTTPResp(resp, err)

	return s
}

// GetChannel - gets channel by ID
func GetChannel(id string, token string) string {
	url := serverAddr + "/channels/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHTTPResp(resp, err)

	return s
}

// UpdateChannel - publishes SenML message on the channel
func UpdateChannel(id string, msg string, token string) string {
	url := serverAddr + "/channels/" + id
	req, err := http.NewRequest("PUT", url, strings.NewReader(msg))
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHTTPResp(resp, err)

	return s
}

// DeleteChannel - removes channel
func DeleteChannel(id string, token string) string {
	url := serverAddr + "/channels/" + id
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHTTPResp(resp, err)

	return s
}

// DeleteAllChannels - removes all channels
func DeleteAllChannels(token string) string {
	url := serverAddr + endPointC
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}

	var list struct {
		Channels []manager.Channel `json:"channels,omitempty"`
	}
	json.Unmarshal([]byte(body), &list)

	s := ""
	for i := 0; i < len(list.Channels); i++ {
		s = s + DeleteChannel(list.Channels[i].ID, token) + "\n\n"
	}

	return s
}
