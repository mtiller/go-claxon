package core

type Claxon struct {
	Schema  string   `json:"$schema,omitempty"`
	Links   []Link   `json:"$links,omitempty"`
	Actions []Action `json:"$actions,omitempty"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	// Here we use `schema` and not `schema` because this field isn't describing
	// the data it appears alongside (i.e., the link data) but rather the schema
	// of the data we will find if follow the provided link.
	Schema string `json:"schema,omitempty"`
}

type Action struct {
	Href           string `json:"href"`
	Id             string `json:"id"`
	Method         string `json:"method,omitempty"`
	RequestSchema  string `json:"reqs,omitempty"`
	ResponseSchema string `json:"ress,omitempty"`
}
