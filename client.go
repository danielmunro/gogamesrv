package gogamesrv

import (
	"bufio"
	"net"
	"strings"
)

type client struct {
	conn    net.Conn
	message string
}

func newClient(conn net.Conn) *client {
	return &client{
		conn: conn,
	}
}

func (c *client) read() {
	c.message, _ = bufio.NewReader(c.conn).ReadString('\n')
	c.message = strings.Trim(c.message, "\r\n")
}

func (c *client) send(message string) {
	c.conn.Write([]byte(message))
}
