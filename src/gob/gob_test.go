package gob

import (
	"encoding/gob"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

type Request struct {
	Name string
	Age  int
}

type Response struct {
	ReplyFor *Request
}

func startServer() (string, <-chan bool) {
	done := make(chan bool)

	tmpdir, err := ioutil.TempDir("", "gob-sock")
	if err != nil {
		panic(err)
	}

	socket := filepath.Join(tmpdir, "socket")
	ln, err := net.Listen("unix", socket)
	if err != nil {
		panic(err)
	}

	go func() {
		defer func() {
			ln.Close()
			os.RemoveAll(tmpdir)
			close(done)
		}()

		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		enc := gob.NewEncoder(conn)
		dec := gob.NewDecoder(conn)

		var req Request
		if err := dec.Decode(&req); err != nil {
			panic(err)
		}
		if err := enc.Encode(Response{&req}); err != nil {
			panic(err)
		}
	}()

	return socket, done
}

func TestGobMessaging(t *testing.T) {
	socket, done := startServer()
	conn, err := net.Dial("unix", socket)
	if err != nil {
		panic(err)
	}

	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	req := Request{"hoge", 20}
	if err := enc.Encode(req); err != nil {
		t.Fatalf("expected nil but got %v", err)
	}

	var resp Response
	if err := dec.Decode(&resp); err != nil {
		t.Fatalf("expected nil but got %v", err)
	}

	<-done // wait for server to quit

	if !reflect.DeepEqual(req, *resp.ReplyFor) {
		t.Errorf("expected %v but got %v", req, *resp.ReplyFor)
	}
}
