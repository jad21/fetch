package fetch

import (
	"encoding/json"
	"fmt"
	"strings"
)

// NewReader this function will return a new reader
// if the format is JSON Valid format it will be convert
// before send, if is not json will send as come.
// Use a POINT for input variable
func NewReader(input interface{}) *strings.Reader {
	bs, err := json.Marshal(input)
	if err != nil {
		return strings.NewReader(fmt.Sprintf("error to read: %T", input))
	}
	return strings.NewReader(string(bs))
}
