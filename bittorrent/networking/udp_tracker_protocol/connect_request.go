package udp_tracker_protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"math/rand"
	"net"
	"net/url"
	"time"
)

type ConnectRequest struct {
	ConnectionID  uint64
	Action        actionCode
	TransactionID uint32
}

func NewConnectRequest() *ConnectRequest {
	return &ConnectRequest{
		ConnectionID:  blankConnectionID,
		Action:        connectAction,
		TransactionID: rand.Uint32(),
	}
}

func (cr *ConnectRequest) Message() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, cr.ConnectionID)
	binary.Write(buf, binary.BigEndian, cr.Action)
	binary.Write(buf, binary.BigEndian, cr.TransactionID)

	return buf.Bytes()
}

func (cr *ConnectRequest) Send(url *url.URL) (response *ConnectResponse, err error) {
	var conn net.Conn

	conn, err = net.DialTimeout(url.Scheme, url.Host, timeout)
	if err != nil {
		return
	}
	defer conn.Close()

	cr.writeRequest(conn)
	response, err = cr.readResponse(conn)

	return
}

func (cr *ConnectRequest) writeRequest(conn net.Conn) {
	conn.Write(cr.Message())
}

func (cr *ConnectRequest) readResponse(conn net.Conn) (response *ConnectResponse, err error) {
	buf := make([]byte, 16)
	readDeadline := time.Now().Add(timeout)

	conn.SetReadDeadline(readDeadline)

	_, err = bufio.NewReader(conn).Read(buf)
	if err != nil {
		return
	}

	response, err = NewConnectResponse(buf, cr)

	return
}
