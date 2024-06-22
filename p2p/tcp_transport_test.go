package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(test *testing.T) {
	listAddr := ":5000"

	trans := NewTCPTransport(listAddr)

	assert.Equal(test, trans.listenerAddress, listAddr)

}
