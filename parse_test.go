package core

import (
	"fmt"
	"testing"

	"github.com/mtiller/rfc8288"
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
				Rel:  "item",
				Type: "application/json",
			}},
		Actions: []Action{
			{Id: "load",
				Href: "./load",
			}},
	}
	s2 := Claxon{
		Links: []Link{
			{Href: "/bar",
				Rel: "item",
			}},
		Actions: []Action{
			{Id: "clear",
				Href: "./clear",
			}},
	}

	l1, err := ToRFC8288Links(s1)
	require.NoError(err)

	l2, err := ToRFC8288Links(s2)
	require.NoError(err)

	h1 := rfc8288.LinkHeader(l1...)
	h2 := rfc8288.LinkHeader(l2...)

	header := fmt.Sprintf("%s\r\n%s\r\n", h1, h2)

	c := ParseLinkHeader(header)

	assert.Equal(Claxon{
		Schema: "#/me",
		Links: []Link{
			{
				Href: "/foo",
				Rel:  "item",
				Type: "application/json",
			},
			{
				Href: "/bar",
				Rel:  "item",
			}},
		Actions: []Action{
			{Id: "load",
				Href: "./load",
			},
			{Id: "clear",
				Href: "./clear",
			}}}, c)
}
