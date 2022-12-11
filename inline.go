package claxon

import (
	"encoding/json"
	"strings"
)

type inlinePreamble struct {
	Schema string `json:"$schema,omitempty"`
	Self   string `json:"$self,omitempty"`
}

type inlineShortLink struct {
	Href  string `json:"href"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`
}

type inlineLinks map[string][]inlineShortLink

type inlinePostscript struct {
	Links   inlineLinks `json:"$links,omitempty"`
	Actions []Action    `json:"$actions,omitempty"`
}

func InlineMarshal(v interface{}, c Claxon) ([]byte, error) {
	pre := inlinePreamble{
		Schema: c.Schema,
		Self:   c.Self,
	}
	preamble, err := json.Marshal(pre)
	if err != nil {
		return preamble, err
	}
	links := inlineLinks{}
	for _, link := range c.Links {
		rel, exists := links[link.Rel]
		if !exists {
			rel = []inlineShortLink{}
		}
		rel = append(rel, inlineShortLink{
			Title: link.Title,
			Href:  link.Href,
			Type:  link.Type,
		})
		links[link.Rel] = rel
	}
	post := inlinePostscript{
		Links:   links,
		Actions: c.Actions,
	}
	postscript, err := json.Marshal(post)
	data, err := json.Marshal(v)
	if err != nil {
		return data, err
	}
	start := string(preamble[1 : len(preamble)-1])
	middle := string(data[1 : len(data)-1])
	end := string(postscript[1 : len(postscript)-1])

	parts := []string{}
	if strings.TrimSpace(start) != "" {
		parts = append(parts, start)
	}
	if strings.TrimSpace(middle) != "" {
		parts = append(parts, middle)
	}
	if strings.TrimSpace(end) != "" {
		parts = append(parts, end)
	}
	return []byte("{" + strings.Join(parts, ",") + "}"), nil
}
