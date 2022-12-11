package claxon

import (
	"encoding/json"
	"fmt"
	"strings"
)

type StringID string

func (x StringID) MarshalJSON() ([]byte, error) {
	return MarshalURN(string(x))
}

type CustomID StringID

type SampleDocument struct {
	Id     StringID `json:"id"`
	Parent CustomID `json:"parent"`
}

func MarshalURN(x string) ([]byte, error) {
	if strings.HasPrefix(x, "urn:foobar:") {
		return json.Marshal(fmt.Sprintf("https://example.com/accounts/%s", x[11:]))
	}
	return json.Marshal(x)
}

func init() {

}
