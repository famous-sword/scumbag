package engine

import "net"

type node struct {
	id       string
	master   bool
	typeName string
	address  net.Addr
}

func (n *node) Addr() net.Addr {
	return n.address
}
