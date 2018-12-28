package mysocket

import (
	"bufio"
	"net"
)

type Conn struct {
	conn net.Conn
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{conn}
}

func (c Conn) Write(msg []byte) error {
	w := bufio.NewWriter(c.conn)
	w.Write(msg)
	w.WriteRune('\r')
	w.WriteRune('\n')
	return w.Flush()
}

func (c Conn) WriteString(msg string) error {
	return c.Write([]byte(msg))
}

func (c Conn) Read() ([]byte, error) {
	r := bufio.NewReader(c.conn)
	msg, _, err := r.ReadLine()
	return msg, err
}

func (c Conn) ReadString() (string, error) {
	msg, err := c.Read()
	return string(msg), err
}

func (c Conn) Close() error {
	return c.conn.Close()
}
