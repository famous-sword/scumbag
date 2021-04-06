package engine

import "net"

type node struct {
	id      string
	master  bool
	address net.Addr
}

func (n *node) Addr() net.Addr {
	return n.address
}
