package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoundTrip(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	s1 := Claxon{
		Schema: "#/me",
		Links: []Link{
			{Href: "/foo",
				Rel:    "item",
				Schema: "#/components/item",
			}},
		Actions: []Action{
			{Id: "load",
				Href: "./load",
			}},
	}
	s2 := Claxon{
		Links: []Link{
			{Href: "/bar",
				Rel:    "item",
				Schema: "#/components/item",
			}},
		Actions: []Action{
			{Id: "clear",
				Href: "./clear",
			}},
	}

	v1, err := linkValue(s1)
	require.NoError(err)

	v2, err := linkValue(s2)
	require.NoError(err)

	header := fmt.Sprintf("Link: %s\r\nLink: %s\r\n", v1, v2)

	c := ParseLinkHeader(header)

	assert.Equal(Claxon{
		Schema: "#/me",
		Links: []Link{
			{
				Href:   "/foo",
				Rel:    "item",
				Schema: "#/components/item",
			},
			{
				Href:   "/bar",
				Rel:    "item",
				Schema: "#/components/item",
			}},
		Actions: []Action{
			{Id: "load",
				Href: "./load",
			},
			{Id: "clear",
				Href: "./clear",
			}}}, c)
}
