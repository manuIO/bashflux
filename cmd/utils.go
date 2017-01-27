package cmd

import (
	"bytes"
	"encoding/json"

	"github.com/hokaccha/go-prettyjson"
)

//dont do this, see above edit
func prettyJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}

func GetPrettyJson(body []byte) string {
	pj, err := prettyjson.Format([]byte(body))
	if err != nil {
		return err.Error()
	} else {
		return string(pj)
	}
}
