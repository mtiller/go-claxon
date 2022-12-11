package claxon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInline(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Serialize ordinary data, if no Claxon data is present, then no impact on
	// the data being serialized
	foo := SampleProperties{
		X: 5,
		Y: "hello",
		Z: true,
	}
	data, err := json.Marshal(Embed(foo, Claxon{}))
	require.NoError(err)
	assert.Equal(`{"x":5,"y":"hello","z":true}`, string(data))

	// You can use embedded structs to combine data and metadata and then just
	// use the `json` package ot serialize
	hyperfoo := Embed(
		SampleProperties{X: 5,
			Y: "hello",
			Z: true,
		},
		Claxon{
			Schema: "https://example.com/schema",
			Self:   "/me",
			Links: []Link{
				{
					Rel:   "item",
					Href:  "/child1",
					Title: "The Favorite",
				},
				{
					Rel:  "item",
					Href: "/child2",
				},
			},
		})
	data, err = json.Marshal(hyperfoo)
	require.NoError(err)
	assert.Equal(`{"$schema":"https://example.com/schema","$self":"/me","x":5,"y":"hello","z":true,"$links":{"item":[{"href":"/child1","title":"The Favorite"},{"href":"/child2"}]}}`, string(data))

	// Finally, you can take ordinary data and annotate it with Claxon data
	// using a `Marshal` function with two arguments.
	data, err = InlineMarshal(foo, Claxon{
		Schema: "https://example.com/schema",
		Self:   "/me",
	})
	require.NoError(err)
	assert.Equal(`{"$schema":"https://example.com/schema","$self":"/me","x":5,"y":"hello","z":true}`, string(data))
}
