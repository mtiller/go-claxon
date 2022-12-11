package claxon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestURIMarshal(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	s1 := SampleDocument{
		Id:     "urn:foobar:abc",
		Parent: "urn:parent:xyz",
	}

	j1, err := json.Marshal(s1)
	require.NoError(err)
	assert.Equal(`{"id":"urn:foobar:abc"}`, string(j1))
}
