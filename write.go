package claxon

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mtiller/rfc8288"
)

func ToRFC8288Links(claxon Claxon) ([]rfc8288.Link, error) {
	ret := []rfc8288.Link{}

	// Add schema as a described by link
	if claxon.Schema != "" {
		link, err := rfc8288.NewLink(claxon.Schema)
		if err != nil {
			return nil, err
		}
		link.Rel = "describedby"
		ret = append(ret, *link)
	}
	if claxon.Self != "" {
		link, err := rfc8288.NewLink(claxon.Self)
		if err != nil {
			return nil, err
		}
		link.Rel = "self"
		ret = append(ret, *link)
	}
	// Add links, if there are any
	for i, link := range claxon.Links {
		if link.Rel == "" {
			return nil, fmt.Errorf("Link %d had empty 'rel' field", i)
		}
		if link.Href == "" {
			return nil, fmt.Errorf("Link %d had empty 'href' field", i)
		}
		l, err := rfc8288.NewLink(link.Href)
		if err != nil {
			return nil, err
		}
		l.Rel = link.Rel
		if link.Title != "" {
			l.Title = link.Title
		}
		if link.Type != "" {
			l.Type = link.Type
		}
		ret = append(ret, *l)
	}
	// Add actions, if there are any
	for i, action := range claxon.Actions {
		if action.Id == "" {
			return nil, fmt.Errorf("action %d had empty 'id' field", i)
		}
		if action.Href == "" {
			return nil, fmt.Errorf("action %d had empty 'href' field", i)
		}
		l, err := rfc8288.NewLink(action.Href)
		if err != nil {
			return nil, err
		}
		l.Title = action.Id
		l.Extend("claxon", "action")
		if action.Method != "" {
			l.Extend("method", action.Method)
		}
		if action.RequestSchema != "" {
			l.Extend("reqs", action.RequestSchema)
		}
		if action.ResponseSchema != "" {
			l.Extend("ress", action.ResponseSchema)
		}
		ret = append(ret, *l)
	}
	return ret, nil
}

func WriteAsHeaders(w http.ResponseWriter, v interface{}, claxon Claxon) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}

	links, err := ToRFC8288Links(claxon)
	if err != nil {
		return err
	}

	val := rfc8288.LinkHeaderValue(links...)
	w.Header().Add("Link", val)
	_, err = w.Write(body)
	return err
}
