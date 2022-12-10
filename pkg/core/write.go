package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//
// Link: <https://one.example.com>; rel="preconnect", <https://two.example.com>; rel="preconnect", <https://three.example.com>; rel="preconnect"

// NB - URLS in headers are not percent encoded.  From what I can tell from
// https://www.rfc-editor.org/rfc/rfc2616, header values don't require encodings
// because they are clearly delimeted by the header name and a terminating
// newline.  So their values can be ordinary octets.  That's why no percent
// encoding takes place here.
func linkValue(claxon Claxon) (string, error) {
	segments := []string{}

	// Add schema as a described by link
	if claxon.Schema != "" {
		segment := fmt.Sprintf("<%s>; rel=\"describedby\"", claxon.Schema)
		segments = append(segments, segment)
	}
	// Add links, if there are any
	for i, link := range claxon.Links {
		if link.Rel == "" {
			return "", fmt.Errorf("Link %d had empty 'rel' field", i)
		}
		if link.Href == "" {
			return "", fmt.Errorf("Link %d had empty 'href' field", i)
		}
		if link.Schema == "" {
			segment := fmt.Sprintf("<%s>; rel=\"%s\"", link.Href, link.Rel)
			segments = append(segments, segment)
		} else {
			segment := fmt.Sprintf("<%s>; rel=\"%s\"; schema=\"%s\"", link.Href, link.Rel, link.Schema)
			segments = append(segments, segment)
		}
	}
	// Add actions, if there are any
	for i, action := range claxon.Actions {
		if action.Id == "" {
			return "", fmt.Errorf("action %d had empty 'id' field", i)
		}
		if action.Href == "" {
			return "", fmt.Errorf("action %d had empty 'href' field", i)
		}
		parts := []string{fmt.Sprintf("<%s>", action.Href), `type="action"`, fmt.Sprintf("id=\"%s\"", action.Id)}
		if action.Method != "" {
			parts = append(parts, fmt.Sprintf("method=\"%s\"", action.Method))
		}
		if action.RequestSchema != "" {
			parts = append(parts, fmt.Sprintf("reqs=\"%s\"", action.RequestSchema))
		}
		if action.ResponseSchema != "" {
			parts = append(parts, fmt.Sprintf("ress=\"%s\"", action.ResponseSchema))
		}
		segment := strings.Join(parts, "; ")
		segments = append(segments, segment)
	}
	return strings.Join(segments, ", "), nil
}

func Write(w http.ResponseWriter, v interface{}, claxon Claxon) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}

	val, err := linkValue(claxon)
	if err != nil {
		return err
	}
	w.Header().Add("Link", val)
	_, err = w.Write(body)
	return err
}
