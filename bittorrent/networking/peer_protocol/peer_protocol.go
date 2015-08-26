package peer_protocol

import (
	"net"
	"time"
)

var (
	timeout, _ = time.ParseDuration("3s")
)

func dial(address string) (conn net.Conn, err error) {
	conn, err = net.DialTimeout("tcp", address, timeout)
	return
}
