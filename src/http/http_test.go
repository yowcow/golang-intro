package http

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestHTTPOverUnixSocket(t *testing.T) {
	dir, err := ioutil.TempDir("", "http-test")
	if err != nil {
		t.Fatal("expected nil but got", err)
	}
	defer os.RemoveAll(dir)

	address := filepath.Join(dir, "http-test.sock")
	ln, err := net.Listen("unix", address)
	if err != nil {
		t.Fatal("expected nil but got", err)
	}
	defer ln.Close()

	handler := http.NewServeMux()
	handler.HandleFunc("/foo/bar/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	server := http.Server{
		Handler: handler,
	}

	ready := make(chan struct{})
	finish := make(chan struct{})
	done := make(chan struct{})

	go func() {
		close(ready)
		if err := server.Serve(ln); err != http.ErrServerClosed {
			t.Fatal("expected http.ErrServerClosed but got", err)
		}
		<-finish // wait server to finish
		close(done)
	}()

	<-ready // server is ready

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", address)
			},
		},
	}

	resp, err := client.Get("http://hoge.fuga/foo/bar/")
	if err != nil {
		t.Fatal("expected nil but got", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("expected nil but got", err)
	}

	if err := server.Shutdown(context.Background()); err != nil {
		t.Fatal("expected nil but got", err)
	}

	close(finish)
	<-done // server is done

	// now do tests
	if resp.StatusCode != 200 {
		t.Error("expected 200 but got", resp.StatusCode)
	}
	if string(body) != "Hello world" {
		t.Error("expected 'Hello world' but got", string(body))
	}
}
