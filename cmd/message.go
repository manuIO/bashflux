package cmd

import (
	"net/http"
	"strings"
)

// GetMsg - gets messages from the channel
func GetMsg(id string, startTime string, endTime string) string {
	url := serverAddr + "/channels/" + id + "/messages" +
		"?start_time=" + startTime + "&end_time=" + endTime
	resp, err := netClient.Get(url)
	s := PrettyHTTPResp(resp, err)

	return s
}

// SendMsg - publishes SenML message on the channel
func SendMsg(id string, msg string, token string) string {
	url := serverAddr + "/channels/" + id + "/messages"
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
