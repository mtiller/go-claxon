package core

import (
	"io"
	"net/http/httptest"
	"testing"

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
				Rel:    "item",
				Schema: "#/components/item",
			}},
		Actions: []Action{
			{Id: "load",
				Href: "./load",
			}},
	}

	val, err := linkValue(hyper)
	require.NoError(err)
	assert.Equal(`<#/me>; rel="describedby", </foo>; rel="item"; schema="#/components/item", <./load>; type="action"; id="load"`, val)

	foo := SampleProperties{
		X: 5,
		Y: "hello",
		Z: true,
	}

	w := httptest.NewRecorder()
	err = Write(w, foo, hyper)
	require.NoError(err)
	res := w.Result()
	aval := res.Header.Get("Link")
	assert.Equal(val, aval)
	body, err := io.ReadAll(res.Body)
	require.NoError(err)
	defer res.Body.Close()
	assert.Equal(`{"x":5,"y":"hello","z":true}`, string(body))
}
