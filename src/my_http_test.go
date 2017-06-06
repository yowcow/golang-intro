package hello

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func myhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world: %s", r.URL)
}

func TestSimpleHttpRequest(t *testing.T) {
	assert := assert.New(t)

	r := httptest.NewRequest("GET", "http://foobar/foo/bar", nil)
	w := httptest.NewRecorder()
	myhandler(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal("Hello world: http://foobar/foo/bar", string(body))
}
