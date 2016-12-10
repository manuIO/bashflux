package cmd

import (
	"bytes"
	"encoding/json"
)

//dont do this, see above edit
func prettyJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}
