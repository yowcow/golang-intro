package hello

import (
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPRequestOut(t *testing.T) {
	query := url.Values{}
	query.Add("hoge-key", "hoge-value")
	query.Add("fuga-key", "fuga-value")

	req, _ := http.NewRequest("POST", "http://hogefuga.com/path/to/endpoint", strings.NewReader(query.Encode()))
	req.Header.Add("User-Agent", "HogeFuga/0.1")
	dump, _ := httputil.DumpRequestOut(req, true)

	assert.Equal(t, "POST /path/to/endpoint HTTP/1.1\r\nHost: hogefuga.com\r\nUser-Agent: HogeFuga/0.1\r\nContent-Length: 39\r\nAccept-Encoding: gzip\r\n\r\nfuga-key=fuga-value&hoge-key=hoge-value", string(dump))
}

func TestHTTPRequest(t *testing.T) {
	reqRe := regexp.MustCompile("POST / HTTP/1.1\r\nHost: 127.0.0.1:\\d+\r\nAccept-Encoding: gzip\r\nContent-Length: 39\r\nUser-Agent: HogeFuga/0.1\r\n\r\nfuga-key=fuga-value&hoge-key=hoge-value")
	respRe := regexp.MustCompile("HTTP/1.1 200 OK\r\nContent-Length: 12\r\nContent-Type: text/plain; charset=utf-8\r\nDate: .+ GMT\r\nFuga: fuga\r\nHoge: hoge\r\n\r\nこんちは")

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		dump, _ := httputil.DumpRequest(req, true)

		assert.True(t, reqRe.Match(dump))

		header := w.Header()
		header.Add("Hoge", "hoge")
		header.Add("Fuga", "fuga")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("こんちは"))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	query := url.Values{}
	query.Add("hoge-key", "hoge-value")
	query.Add("fuga-key", "fuga-value")

	req, _ := http.NewRequest("POST", server.URL, strings.NewReader(query.Encode()))
	req.Header.Add("User-Agent", "HogeFuga/0.1")

	client := http.Client{}
	resp, _ := client.Do(req)
	dump, _ := httputil.DumpResponse(resp, true)

	assert.True(t, respRe.Match(dump))
}
