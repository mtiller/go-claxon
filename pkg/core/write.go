package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mtiller/rfc8288"
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
		link, err := rfc8288.NewLink(claxon.Schema)
		link.Rel = "describedby"
		if err != nil {
			return "", err
		}
		segments = append(segments, link.String())
	}
	// Add links, if there are any
	for i, link := range claxon.Links {
		if link.Rel == "" {
			return "", fmt.Errorf("Link %d had empty 'rel' field", i)
		}
		if link.Href == "" {
			return "", fmt.Errorf("Link %d had empty 'href' field", i)
		}
		l, err := rfc8288.NewLink(link.Href)
		if err != nil {
			return "", err
		}
		l.Rel = link.Rel
		if link.Schema != "" {
			l.Extend("schema", link.Schema)
		}
		segments = append(segments, l.String())
	}
	// Add actions, if there are any
	for i, action := range claxon.Actions {
		if action.Id == "" {
			return "", fmt.Errorf("action %d had empty 'id' field", i)
		}
		if action.Href == "" {
			return "", fmt.Errorf("action %d had empty 'href' field", i)
		}
		l, err := rfc8288.NewLink(action.Href)
		if err != nil {
			return "", err
		}
		l.Extend("claxon", "action")
		l.Extend("id", action.Id)
		if action.Method != "" {
			l.Extend("method", action.Method)
		}
		if action.RequestSchema != "" {
			l.Extend("reqs", action.RequestSchema)
		}
		if action.ResponseSchema != "" {
			l.Extend("ress", action.ResponseSchema)
		}
		segments = append(segments, l.String())
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
