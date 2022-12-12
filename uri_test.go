package claxon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type StringID string

func (x StringID) MarshalJSON() ([]byte, error) {
	return MarshalURN(string(x))
}

type CustomID StringID

func (x CustomID) MarshalJSON() ([]byte, error) {
	return MarshalURN(string(x))
}

type SampleDocument struct {
	Id     StringID `json:"id"`
	Parent CustomID `json:"parent"`
}

func TestURIMarshal(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	s1 := SampleDocument{
		Id:     "urn:foobar:abc",
		Parent: "urn:parent:xyz",
	}

	j1, err := json.Marshal(s1)
	require.NoError(err)
	assert.Equal(`{"id":"urn:foobar:abc","parent":"urn:parent:xyz"}`, string(j1))

	UsingMap(map[string]string{
		"urn:foobar:": "https://example.com/account/",
	}, func() {
		j1, err = json.Marshal(s1)
	})
	require.NoError(err)
	assert.Equal(`{"id":"https://example.com/account/abc","parent":"urn:parent:xyz"}`, string(j1))

	UsingMap(map[string]string{
		"urn:parent:": "/link/",
	}, func() {
		j1, err = json.Marshal(s1)
	})
	require.NoError(err)
	assert.Equal(`{"id":"urn:foobar:abc","parent":"/link/xyz"}`, string(j1))

	UsingMap(map[string]string{
		"urn:foobar:": "https://example.com/account/",
		"urn:parent:": "/link/",
	}, func() {
		j1, err = json.Marshal(s1)
	})
	require.NoError(err)
	assert.Equal(`{"id":"https://example.com/account/abc","parent":"/link/xyz"}`, string(j1))
}
