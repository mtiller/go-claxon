package claxon

type Claxon struct {
	Schema  string   `json:"$schema,omitempty"`
	Links   []Link   `json:"$links,omitempty"`
	Actions []Action `json:"$actions,omitempty"`
}

type Link struct {
	Href  string `json:"href"`
	Rel   string `json:"rel"`
	Title string `json:"title"`
	Type  string `json:"type,omitempty"`
}

type Action struct {
	Href           string `json:"href"`
	Id             string `json:"id"`
	Method         string `json:"method,omitempty"`
	RequestSchema  string `json:"reqs,omitempty"`
	ResponseSchema string `json:"ress,omitempty"`
}

// The AddLink method takes the two required arguments (rel and href).
// If a third argument is present, it is interpreted as the title
// and the fourth argument as the content type
func (c *Claxon) AddLink(rel string, href string, more ...string) *Claxon {
	link := Link{
		Rel:  rel,
		Href: href,
	}
	if len(more) >= 1 {
		link.Title = more[0]
	}
	if len(more) >= 2 {
		link.Type = more[1]
	}
	links := c.Links
	if links == nil {
		links = []Link{}
	}
	links = append(links, link)
	c.Links = links
	return c
}

// The AddAction method takes the two required arguments (id and href).
// If a third argument is present, it is interpreted as the method.
// If a fourth argument is present, it is interpreted as the **response** schema
// If a fifth argument is present, it is interpreted as the **request** schema
func (c *Claxon) AddAction(id string, href string, more ...string) *Claxon {
	action := Action{
		Id:   id,
		Href: href,
	}
	if len(more) >= 1 {
		action.Method = more[0]
	}
	if len(more) >= 2 {
		action.ResponseSchema = more[1]
	}
	if len(more) >= 3 {
		action.RequestSchema = more[2]
	}
	actions := c.Actions
	if actions == nil {
		actions = []Action{}
	}
	actions = append(actions, action)
	c.Actions = actions
	return c
}
