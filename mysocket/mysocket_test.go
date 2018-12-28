package mysocket_test

import (
	"bufio"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	socket "github.com/yowcow/golang-intro/mysocket"
)

func TestServer(t *testing.T) {
	s, err := socket.NewServer("tcp", ":12345")

	assert.Nil(t, err)
	assert.Nil(t, s.Close())
}

func TestConnReadWrite(t *testing.T) {
	s, _ := socket.NewServer("tcp", ":12345")
	defer s.Close()
	q := make(chan bool)

	go func(s *socket.Server, q chan<- bool) {
		if c, _ := s.Accept(); c != nil {
			line, err := c.Read()

			assert.Nil(t, err)
			assert.Equal(t, "hello world", string(line))

			err = c.Write([]byte("bye world"))

			assert.Nil(t, err)

			c.Close()
			q <- true
			return
		}
		q <- false
	}(s, q)

	conn, err := net.Dial("tcp", "localhost:12345")

	assert.Nil(t, err)

	conn.Write([]byte("hello world\r\n"))
	line, _, err := bufio.NewReader(conn).ReadLine()

	assert.Nil(t, conn.Close())

	assert.Equal(t, "bye world", string(line))
	assert.True(t, <-q)
}

func TestConnReadStringWriteString(t *testing.T) {
	s, err := socket.NewServer("tcp", ":12345")
	defer s.Close()
	q := make(chan bool)

	go func(s *socket.Server, q chan<- bool) {
		if c, _ := s.Accept(); c != nil {
			line, err := c.ReadString()

			assert.Nil(t, err)
			assert.Equal(t, "hello world", line)

			err = c.WriteString("bye world")

			assert.Nil(t, err)

			c.Close()
			q <- true
			return
		}
		q <- false
	}(s, q)

	conn, err := net.Dial("tcp", "localhost:12345")

	assert.Nil(t, err)

	w := bufio.NewWriter(conn)
	w.WriteString("hello world")
	w.WriteRune('\r')
	w.WriteRune('\n')

	assert.Nil(t, w.Flush())

	r := bufio.NewReader(conn)
	line, _, err := r.ReadLine()

	assert.Equal(t, "bye world", string(line))
	assert.True(t, <-q)
}
