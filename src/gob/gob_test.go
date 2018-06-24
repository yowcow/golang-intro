package gob

import (
	"encoding/gob"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

type Request struct {
	Name string
	Age  int
}

type Response struct {
	Time     *time.Time
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
		now := time.Now()
		if err := enc.Encode(Response{&now, &req}); err != nil {
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

type Response2 struct {
	Number    int
	NumberPtr *int
	ReplyFor  *Request
}

func TestGobMessagingDecodeToDifferntType(t *testing.T) {
	socket, done := startServer()
	conn, err := net.Dial("unix", socket)
	if err != nil {
		panic(err)
	}

	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	req := Request{"fuga", 30}
	if err := enc.Encode(req); err != nil {
		t.Fatalf("expected nil but got %v", err)
	}

	var resp Response2
	if err := dec.Decode(&resp); err != nil {
		t.Fatalf("expected nil but got %v", err)
	}

	<-done // wait for server to quit

	if resp.Number != 0 {
		t.Errorf("expected 0 but got %v", resp.Number)
	}
	if resp.NumberPtr != nil {
		t.Errorf("expected nil but got %v", resp.NumberPtr)
	}
	if !reflect.DeepEqual(req, *resp.ReplyFor) {
		t.Errorf("expected %v but got %v", req, *resp.ReplyFor)
	}
}
