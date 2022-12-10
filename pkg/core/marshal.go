package core

import "encoding/json"

func Marshal(v interface{}, claxon Claxon) ([]byte, error) {
	preamble, err := json.Marshal(claxon)
	if err != nil {
		return preamble, err
	}
	data, err := json.Marshal(v)
	if err != nil {
		return data, err
	}
	if len(preamble) <= 2 {
		return data, nil
	}
	if len(data) <= 2 {
		return preamble, nil
	}
	return []byte("{" + string(preamble[1:len(preamble)-1]) + "," + string(data[1:len(data)-1]) + "}"), nil
}
