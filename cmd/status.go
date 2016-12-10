package cmd

import (
	"fmt"
	"io/ioutil"
)

func Status() string {
	url := UrlHTTP + "/status"
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
