package gogamesrv

import (
	"net"
	"strconv"
)

type server struct {
	clients []*client
}

func NewServer() *server {
	return &server{}
}

func (s *server) listen(port int) error {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	defer ln.Close()
	//updater := make(chan *client)

	for {
		go s.addClient(newClient(s.accept(ln)))
	}
}

func (s *server) accept(l net.Listener) net.Conn {
	conn, _ := l.Accept()
	return conn
}

func (s *server) addClient(c *client) {
	s.clients = append(s.clients, c)
}
