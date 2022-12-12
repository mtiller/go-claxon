package claxon

import (
	"encoding/json"
	"fmt"
	"strings"
)

var mapping map[string]string = nil

// Given a specified mapping of URN prefix to URL prefix, execute the specific
// function.  This will cause any JSON marshaling that happens as a result of
// that function invocation to transform the strings based on swapping prefixes.
// TODO: Create a more general version where a mapping function can be employed
// TODO: Allow the mapping to involve some template (i.e., assume the ending of
// the URN is an ID but allow the URL to be a template of some kind).
func UsingMap(x map[string]string, f func()) {
	old := mapping
	mapping = x
	f()
	mapping = old
}

// TODO: Implement the same functionality for Unmarshaling.  The key here is
// that whatever transformation is being applied must be reversible so we can
// easily translate between URNs for internal use and URLs for external
// references.

// QUESTION: Is this translation really necessary or useful.  If the API
// resources are not identical to the "nouns" of our system but are instead just
// agents that operate on those nouns, then we could use the URNs as arguments
// in our API calls and never have to translate because all the URLs would be
// those of the agents.

func MarshalURN(x string) ([]byte, error) {
	if mapping == nil {
		return json.Marshal(x)
	}
	for urn, url := range mapping {
		if strings.HasPrefix(x, urn) {
			id := strings.TrimPrefix(x, urn)
			path := fmt.Sprintf("%s%s", url, id)
			return json.Marshal(path)
		}
	}
	return json.Marshal(x)
}

func init() {

}
