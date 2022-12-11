package claxon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type SampleProperties struct {
	X int    `json:"x"`
	Y string `json:"y"`
	Z bool   `json:"z"`
}

type SampleClaxson struct {
	Claxon
	SampleProperties
}

func TestSerialize(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Serialize ordinary data, if no Claxon data is present, then no impact on
	// the data being serialized
	foo := SampleProperties{
		X: 5,
		Y: "hello",
		Z: true,
	}
	data, err := Marshal(foo, Claxon{})
	require.NoError(err)
	assert.Equal(`{"x":5,"y":"hello","z":true}`, string(data))

	// You can use embedded structs to combine data and metadata and then just
	// use the `json` package ot serialize
	hyperfoo := SampleClaxson{
		Claxon{
			Schema: "https://example.com/schema",
			Self:   "/me",
		},
		SampleProperties{X: 5,
			Y: "hello",
			Z: true,
		},
	}
	data, err = json.Marshal(hyperfoo)
	require.NoError(err)
	assert.Equal(`{"$schema":"https://example.com/schema","$self":"/me","x":5,"y":"hello","z":true}`, string(data))

	// Finally, you can take ordinary data and annotate it with Claxon data
	// using a `Marhsal` function with two arguments.
	data, err = Marshal(foo, Claxon{
		Schema: "https://example.com/schema",
		Self:   "/me",
	})
	require.NoError(err)
	assert.Equal(`{"$schema":"https://example.com/schema","$self":"/me","x":5,"y":"hello","z":true}`, string(data))
}
