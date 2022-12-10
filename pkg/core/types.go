package core

type ClaxonFields struct {
}

type ClaxonPayload struct {
	Schema  string          `json:"$schema,omitempty"`
	Links   []ClaxonLinks   `json:"$links,omitempty"`
	Actions []ClaxonActions `json:"$actions,omitempty"`
}

type ClaxonLinks struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	// Here we use `schema` and not `schema` because this field isn't describing
	// the data it appears alongside (i.e., the link data) but rather the schema
	// of the data we will find if follow the provided link.
	Schema string `json:"schema,omitempty"`
}

type ClaxonActions struct {
	Id             string `json:"id"`
	Href           string `json:"href"`
	Method         string `json:"method,omitempty"`
	RequestSchema  string `json:"reqs,omitempty"`
	ResponseSchema string `json:"ress,omitempty"`
}
