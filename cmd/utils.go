package cmd

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/hokaccha/go-prettyjson"
)

// PrettyJSON - JSON pretty print
func PrettyJSON(body string) string {
	pj, err := prettyjson.Format([]byte(body))
	if err != nil {
		return err.Error()
	}

	return string(pj)
}

// PrettyHTTPResp - format http response
func PrettyHTTPResp(resp *http.Response, err error) string {
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}

	str := "Status code: " + strconv.Itoa(resp.StatusCode) + " - " +
		http.StatusText(resp.StatusCode) + "\n\n"

	if len(resp.Header.Get("Location")) != 0 {
		str = str + "Resource location: " + resp.Header.Get("Location")
	}

	if len(body) != 0 {
		str = str + "\n" + PrettyJSON(string(body))
	}

	return str
}
