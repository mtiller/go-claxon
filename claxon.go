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
