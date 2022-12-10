package core

import (
	"strings"

	"github.com/mtiller/rfc8288"
)

func ParseLinkHeader(s ...string) Claxon {
	ret := Claxon{
		Links:   []Link{},
		Actions: []Action{},
	}
	all := strings.Join(s, "\r\n")
	links, err := rfc8288.ParseLinkHeaders(all)
	if err != nil {
		return ret
	}
	for _, link := range links {
		if link.Rel == "describedby" {
			ret.Schema = link.HREF.String()
			continue
		}
		key, exists := link.StringExtension("claxon")
		if exists && key == "action" {
			add := Action{
				Href: link.HREF.String(),
			}
			id, has := link.StringExtension("id")
			if has {
				add.Id = id
			}
			method, has := link.StringExtension("method")
			if has {
				add.Method = method
			}
			reqs, has := link.StringExtension("reqs")
			if has {
				add.RequestSchema = reqs
			}
			ress, has := link.StringExtension("ress")
			if has {
				add.ResponseSchema = ress
			}
			ret.Actions = append(ret.Actions, add)
		} else {
			add := Link{
				Href: link.HREF.String(),
				Rel:  link.Rel,
			}
			if link.Type != "" {
				add.Type = link.Type
			}
			ret.Links = append(ret.Links, add)
		}
	}
	return ret
}
