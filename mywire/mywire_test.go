package mywire_test

import (
	"bytes"
	"testing"

	"github.com/yowcow/golang-intro/mywire"
)

func TestBareIntegration(t *testing.T) {
	var buf bytes.Buffer

	event, err := mywire.InitializeEvent("hoge!!", &buf)
	if err != nil {
		t.Fatal("expected no error but got", err)
	}

	event.Start()
	if buf.String() != "hoge!!\n" {
		t.Error("expected 'hoge!!\\n' but got", buf.String())
	}
}
