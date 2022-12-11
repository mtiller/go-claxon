package claxon

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/mtiller/rfc8288"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinkHeaders(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	hyper := Claxon{
		Schema: "#/me",
		Links: []Link{
			{Href: "/foo",
				Rel:   "item",
				Title: "Foo",
				Type:  "application/json",
			}},
		Actions: []Action{
			{Id: "load",
				Href: "./load",
			}},
	}

	links, err := ToRFC8288Links(hyper)
	require.NoError(err)
	header := rfc8288.LinkHeader(links...)
	assert.Equal(`Link: <#/me>; rel="describedby", </foo>; rel="item"; title="Foo"; type="application/json", <./load>; title="load"; claxon="action"`, header)

	foo := SampleProperties{
		X: 5,
		Y: "hello",
		Z: true,
	}

	w := httptest.NewRecorder()
	err = WriteAsHeaders(w, foo, hyper)
	require.NoError(err)
	res := w.Result()
	aval := res.Header.Get("Link")
	assert.Equal(header[6:], aval)
	body, err := io.ReadAll(res.Body)
	require.NoError(err)
	defer res.Body.Close()
	assert.Equal(`{"x":5,"y":"hello","z":true}`, string(body))
}
