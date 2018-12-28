package mypipe

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

	// close input to subprocess and wait for it to finish
	w.Close()
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

func TestStdoutStdinPipelining(t *testing.T) {
	cmd1 := exec.Command("cat", "data.txt")
	cmd1Stdout, err := cmd1.StdoutPipe()
	if err != nil {
		t.Fatal("expected nil but got", err)
	}

	cmd2 := exec.Command("grep", "hoge")
	cmd2.Stdin = cmd1Stdout
	cmd2Stdout, err := cmd2.StdoutPipe()
	if err != nil {
		t.Fatal("expected nil but got", err)
	}

	cmd3 := exec.Command("grep", "-v", "fuga")
	cmd3.Stdin = cmd2Stdout
	r, err := cmd3.StdoutPipe()
	if err != nil {
		t.Fatal("expected nil but got", err)
	}

	if err := cmd1.Start(); err != nil {
		t.Fatal("expected nil but got", err)
	}

	if err := cmd2.Start(); err != nil {
		t.Fatal("expected nil but got", err)
	}

	if err := cmd3.Start(); err != nil {
		t.Fatal("expected nil but got", err)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal("expected nil but got", err)
	}

	// wait for all processes to finish
	cmd1.Wait()
	cmd2.Wait()
	cmd3.Wait()

	expected := "hogehoge\n"
	if string(data) != expected {
		t.Error("expected", expected, "but got", string(data))
	}
}
