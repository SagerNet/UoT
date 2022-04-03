package uot

import (
	"encoding/binary"
	"io"
	"net"

	"github.com/sagernet/uot/common/socksaddr"
)

type ClientConn struct {
	net.Conn
}

func NewClientConn(conn net.Conn) net.PacketConn {
	return &ClientConn{conn}
}

func (c *ClientConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	address, port, err := AddrParser.ReadAddressAndPort(c)
	if err != nil {
		return 0, nil, err
	}
	var length uint16
	err = binary.Read(c, binary.BigEndian, &length)
	if err != nil {
		return 0, nil, err
	}
	if len(p) < int(length) {
		return 0, nil, io.ErrShortBuffer
	}
	n, err = io.ReadAtLeast(c, p, int(length))
	if err != nil {
		return 0, nil, err
	}
	addr = &net.UDPAddr{
		IP:   address.Addr().AsSlice(),
		Port: int(port),
	}
	return
}

func (c *ClientConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	address, port := socksaddr.AddressFromNetAddr(addr)
	err = AddrParser.WriteAddressAndPort(c, address, port)
	if err != nil {
		return
	}
	err = binary.Write(c, binary.BigEndian, uint16(len(p)))
	if err != nil {
		return
	}
	return c.Write(p)
}
