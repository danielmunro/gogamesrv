package gogamesrv

import (
	"fmt"
	"net"
	"strconv"
)

// Server is the combination of a port listener and a proxy for clients.
type Server struct {
	clients  []*client
	listener chan *client
}

// NewServer creates a new server instance.
func NewServer() *Server {
	return &Server{
		listener: make(chan *client),
	}
}

// Listen starts the two main threads of the server: the client input listener,
// and the new connection listener.
func (s *Server) Listen(port int) {
	ln, _ := net.Listen("tcp", ":"+strconv.Itoa(port))
	defer ln.Close()
	go s.startBlockingClientInputListener()
	s.startBlockingNewClientLoop(ln)
}

func (s *Server) startBlockingNewClientLoop(ln net.Listener) {
	for {
		go s.addClient(newClient(s.accept(ln)))
	}
}

func (s *Server) startBlockingClientInputListener() {
	for {
		select {
		case c := <-s.listener:
			s.broadcast(c)
		}
	}
}

func (s *Server) broadcast(c *client) {
	s.forAllClients(func(client *client) {
		client.send(getPayload(c, client))
	})
}

func (s *Server) accept(l net.Listener) net.Conn {
	conn, _ := l.Accept()
	return conn
}

func (s *Server) addClient(c *client) {
	s.clients = append(s.clients, c)
	for {
		c.read()
		s.listener <- c
	}
}

func (s *Server) forAllClients(f func(*client)) {
	for _, client := range s.clients {
		f(client)
	}
}

func getPayload(sender *client, receiver *client) string {
	return fmt.Sprintf(
		`{sender: "%s", message: "%s"}`,
		getSenderName(sender, receiver),
		sender.message,
	)
}

func getSenderName(sender *client, receiver *client) string {
	if sender == receiver {
		return "you"
	}

	return sender.conn.RemoteAddr().String()
}
