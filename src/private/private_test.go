package private_test

import (
	"testing"

	"github.com/yowcow/golang-intro/src/private"
)

func TestSecret(t *testing.T) {
	s := private.NewSecret("hogehoge")

	if actual := s.ExportGetMessage(); actual != "hogehoge" {
		t.Error("expected hogehoge but got", actual)
	}

	// calling unbound func with first argument for binding receiver
	if actual := private.ExportGetMessageFunc(s); actual != "hogehoge" {
		t.Error("expected hogehoge but got", actual)
	}
}
