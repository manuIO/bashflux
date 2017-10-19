package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mainflux/mainflux/manager"
)

var endPoint = "/clients"

// CreateClient - creates new client and generates client UUID
func CreateClient(msg string, token string) string {
	url := serverAddr + endPoint
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

// GetClients - gets all clients
func GetClients(token string) string {
	url := serverAddr + endPoint
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

// GetClient - gets client by ID
func GetClient(id string, token string) string {
	url := serverAddr + endPoint + "/" + id
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

// UpdateClient - updates client by ID
func UpdateClient(id string, msg string, token string) string {
	url := serverAddr + endPoint + "/" + id
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

// DeleteClient - removes client
func DeleteClient(id string, token string) string {
	url := serverAddr + endPoint + "/" + id
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

// DeleteAllClients - removes all clients
func DeleteAllClients(token string) string {
	url := serverAddr + endPoint
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

	var clients struct {
		List []manager.Client `json:"clients,omitempty"`
	}
	json.Unmarshal([]byte(body), &clients)

	s := ""
	for i := 0; i < len(clients.List); i++ {
		s = s + DeleteClient(clients.List[i].ID, token) + "\n\n"
	}

	return s
}
