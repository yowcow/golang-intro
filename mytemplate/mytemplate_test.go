package mytemplate

import (
	"bytes"
	"html/template"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tpl, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		t.Fatal("expected no error but got", err)
	}

	var buf bytes.Buffer
	err = tpl.ExecuteTemplate(&buf, "T", "<script>pwned!</script>")
	if err != nil {
		t.Fatal("expected no error but got", err)
	}

	expected := `Hello, &lt;script&gt;pwned!&lt;/script&gt;!`
	if buf.String() != expected {
		t.Error("expected", expected, "but got", buf.String())
	}
}
