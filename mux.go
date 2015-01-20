package main

import (
	"io"
	"log"
	"net"
)

type Protocol interface {
	// address to proxy to
	Address() string

	// identify protocol from header
	Identify(header []byte) bool
}

type Mux struct {
	Handlers []Protocol
}

// create a new Mux assignment
func NewMux() *Mux {
	return &Mux{}
}

// add a protocol to mux handler set
func (m *Mux) Handle(p Protocol) {
	m.Handlers = append(m.Handlers, p)
}

// match protocol to handler
// returns address to proxy to
func (m *Mux) Identify(header []byte) (address string) {
	if len(m.Handlers) < 1 {
		return ""
	}

	for _, handler := range m.Handlers {
		if handler.Identify(header) {
			return handler.Address()
		}
	}

	// return address of last handler, default
	return m.Handlers[len(m.Handlers)-1].Address()
}

// create a server on given address and handle incoming connections
func (m *Mux) ListenAndServe(addr string) error {
	server, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			return err
		}

		go m.Serve(conn)
	}

	return nil
}

// serve takes an incomming connection, applies configured protocol
// handlers and proxies the connection based on result
func (m *Mux) Serve(conn net.Conn) error {
	defer conn.Close()

	// get first 3 bytes of connection as header
	header := make([]byte, 3)
	if _, err := io.ReadAtLeast(conn, header, 3); err != nil {
		return err
	}

	// identify protocol from header
	address := m.Identify(header)

	log.Printf("[INFO] proxy: from=%s to=%s\n", conn.RemoteAddr(), address)

	// connect to remote
	remote, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("[ERROR] remote: %s\n", err)
		return err
	}
	defer remote.Close()

	// write header we chopped back to remote
	remote.Write(header)

	// proxy between us and remote server
	err = Shovel(conn, remote)
	if err != nil {
		return err
	}

	return nil
}
