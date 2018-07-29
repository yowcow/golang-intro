package main

import (
	"bytes"
	"io"
	"testing"
	"testing/iotest"
)

func TestDataErrReader(t *testing.T) {
	buf := bytes.NewBufferString("hogefuga")
	r := iotest.DataErrReader(buf)

	tmp := make([]byte, 4)
	n, err := r.Read(tmp)
	if n != 4 {
		t.Error("expected 4 but got", n)
	}
	if err != nil {
		t.Error("expected no error but got", err)
	}
	if string(tmp) != "hoge" {
		t.Error("expected 'hoge' but got", string(tmp))
	}

	n, err = r.Read(tmp)
	if n != 4 {
		t.Error("expected 4 but got", n)
	}
	if err == nil {
		t.Error("expected error but got nil")
	}
	if string(tmp) != "fuga" {
		t.Error("expected 'fuga' but got", string(tmp))
	}
}

func TestHalfReader(t *testing.T) {
	buf := bytes.NewBufferString("hogefuga")
	r := iotest.HalfReader(buf)

	tmp := make([]byte, 8)
	n, err := r.Read(tmp)
	if n != 4 {
		t.Error("expected 4 but got", n)
	}
	if err != nil {
		t.Error("expected no error but got", err)
	}
	if string(tmp[:n]) != "hoge" {
		t.Error("expected 'hoge' but got", string(tmp[:n]))
	}

	n, err = r.Read(tmp)
	if n != 4 {
		t.Error("expected 4 but got", n)
	}
	if err != nil {
		t.Error("expected no error but got", err)
	}
	if string(tmp[:n]) != "fuga" {
		t.Error("expected 'fuga' but got", string(tmp[:n]))
	}

	n, err = r.Read(tmp)
	if n != 0 {
		t.Error("expected 0 but got", n)
	}
	if err != io.EOF {
		t.Error("expected EOF but got", err)
	}
}

func TestTruncateWriter(t *testing.T) {
	var buf bytes.Buffer
	w := iotest.TruncateWriter(&buf, 4)

	n, err := io.WriteString(w, "hoge")
	if n != 4 {
		t.Error("expected 4 but got", n)
	}
	if err != nil {
		t.Error("expected no error but got", err)
	}
	if buf.String() != "hoge" {
		t.Error("expected 'hoge' but got", buf.String())
	}

	n, err = io.WriteString(w, "fuga")
	if n != 4 {
		t.Error("expected 4 but got", n)
	}
	if err != nil {
		t.Error("expected no error but got", err)
	}
	if buf.String() != "hoge" {
		t.Error("expected 'hoge' but got", buf.String())
	}
}
