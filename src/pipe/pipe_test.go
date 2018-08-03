package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestPipe(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal("expected nil but got", err)
	}

	done := make(chan struct{})
	go func() {
		defer w.Close()
		defer close(done)
		w.WriteString("hoge")
		w.WriteString("fuga")
	}()

	cases := []struct {
		subtest        string
		expected       string
		expectedLength int
	}{
		{
			"1st output is 'hoge'",
			"hoge",
			4,
		},
		{
			"2nd output is 'fuga'",
			"fuga",
			4,
		},
	}

	for _, c := range cases {
		t.Run(c.subtest, func(t *testing.T) {
			tmp := make([]byte, c.expectedLength)
			n, err := r.Read(tmp)
			if err != nil {
				t.Error("expected nil but got", err)
			}
			if n != c.expectedLength {
				t.Error("expected", c.expectedLength, "but got", n)
			}
			if string(tmp) != c.expected {
				t.Error("expected", c.expected, "but got", string(tmp))
			}
		})
	}

	<-done
}

// TestSubprocessStdinPipe requires `cronolog` installed
func TestSubprocessStdinPipe(t *testing.T) {
	now := time.Now()
	expectedFile := fmt.Sprintf("test-%s.log", now.Format("2006-01-02"))
	os.Remove(expectedFile) // clean first

	// write
	cmd := exec.Command("cronolog", "test-%Y-%m-%d.log")
	w, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal("expected nil but got", err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatal("expected nil but got", err)
	}

	if _, err := io.WriteString(w, "hogehoge\n"); err != nil {
		t.Error("expected nil but got", err)
	}
	if _, err := io.WriteString(w, "fugafuga\n"); err != nil {
		t.Error("expected nil but got", err)
	}
	w.Close()

	// Await a little to make sure buffer is flushed to file.
	// This shouldn't be a good practice.
	time.Sleep(100 * time.Millisecond)

	if err := cmd.Process.Kill(); err != nil {
		t.Fatal("expected nil but got", err)
	}

	cmd.Wait()

	f, err := os.Open(expectedFile)
	if err != nil {
		t.Fatal("expected nil but got", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal("expected nil but got", err)
	}

	expectedOutput := `hogehoge
fugafuga
`
	if string(b) != expectedOutput {
		t.Error("expected", expectedOutput, "but got", string(b))
	}
}
