package myprivate_test

import (
	"testing"

	"github.com/yowcow/golang-intro/myprivate"
)

func TestSecret(t *testing.T) {
	s := myprivate.NewSecret("hogehoge")

	if actual := s.ExportGetMessage(); actual != "hogehoge" {
		t.Error("expected hogehoge but got", actual)
	}

	// calling unbound func with first argument for binding receiver
	if actual := myprivate.ExportGetMessageFunc(s); actual != "hogehoge" {
		t.Error("expected hogehoge but got", actual)
	}
}
