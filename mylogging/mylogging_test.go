package mylogging_test

import (
	"bytes"
	"log"
	"regexp"
	"testing"

	"github.com/yowcow/golang-intro/mylogging"
)

func TestDebug(t *testing.T) {
	buf := new(bytes.Buffer)
	l := mylogging.New(2, buf, "", log.Lshortfile)
	l.Debug("hoge", "fuga")

	expected := regexp.MustCompile(`\Amylogging_test\.go\:\d+\: \[DEBUG\] hogefuga\n`)
	if !expected.MatchString(buf.String()) {
		t.Error("expected '", expected, "' but got '", buf.String())
	}
}

func TestDebugln(t *testing.T) {
	buf := new(bytes.Buffer)
	l := mylogging.New(2, buf, "", log.Lshortfile)
	l.Debugln("hoge", "fuga")

	expected := regexp.MustCompile(`\Amylogging_test\.go\:\d+\: \[DEBUG\] hoge fuga\n`)
	if !expected.MatchString(buf.String()) {
		t.Error("expected '", expected, "' but got '", buf.String())
	}
}

func TestDebugf(t *testing.T) {
	buf := new(bytes.Buffer)
	l := mylogging.New(2, buf, "", log.Lshortfile)
	l.Debugf("%s-%s", "hoge", "fuga")

	expected := regexp.MustCompile(`\Amylogging_test\.go\:\d+\: \[DEBUG\] hoge\-fuga\n`)
	if !expected.MatchString(buf.String()) {
		t.Error("expected '", expected, "' but got '", buf.String())
	}
}
