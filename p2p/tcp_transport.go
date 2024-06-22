package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	listenerAddress string
	listener        net.Listener

	mutex sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenerAddress: listenAddr,
	}

}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenerAddress)

	if err != nil {
		return err
	}

	go t.StartAcceptLoop()
	return nil
}

func (t *TCPTransport) StartAcceptLoop() {
	for {
		conn, err := t.listener.Accept()

		if err != nil {
			fmt.Println("TCP accept error: %s\n", err)
		}

	}
}
