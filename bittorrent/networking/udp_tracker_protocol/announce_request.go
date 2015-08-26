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

type AnnounceRequest struct {
	ConnectionID  uint64
	Action        actionCode
	TransactionID uint32
	InfoHash      string
	PeerID        string
	Downloaded    uint64
	Left          uint64
	Uploaded      uint64
	Event         eventCode
	IPAddress     uint32
	Key           uint32
	NumWant       int32
	Port          uint16
}

func NewAnnounceRequest(connectResponse *ConnectResponse, peerID string, infoHash string, port uint16) *AnnounceRequest {
	return &AnnounceRequest{
		ConnectionID:  connectResponse.ConnectionID,
		Action:        announceAction,
		TransactionID: rand.Uint32(),
		InfoHash:      infoHash,
		PeerID:        peerID,
		Downloaded:    0,
		Left:          0,
		Uploaded:      0,
		Event:         noneEvent,
		IPAddress:     0,
		Key:           0,
		NumWant:       -1,
		Port:          port,
	}
}

func (ar *AnnounceRequest) Message() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, ar.ConnectionID)
	binary.Write(buf, binary.BigEndian, ar.Action)
	binary.Write(buf, binary.BigEndian, ar.TransactionID)
	binary.Write(buf, binary.BigEndian, []byte(ar.InfoHash))
	binary.Write(buf, binary.BigEndian, []byte(ar.PeerID))
	binary.Write(buf, binary.BigEndian, ar.Downloaded)
	binary.Write(buf, binary.BigEndian, ar.Left)
	binary.Write(buf, binary.BigEndian, ar.Uploaded)
	binary.Write(buf, binary.BigEndian, ar.Event)
	binary.Write(buf, binary.BigEndian, ar.IPAddress)
	binary.Write(buf, binary.BigEndian, ar.Key)
	binary.Write(buf, binary.BigEndian, ar.NumWant)
	binary.Write(buf, binary.BigEndian, ar.Port)

	return buf.Bytes()
}

func (ar *AnnounceRequest) Send(url *url.URL) (response *AnnounceResponse, err error) {
	var conn net.Conn

	conn, err = net.DialTimeout(url.Scheme, url.Host, timeout)
	if err != nil {
		return
	}
	defer conn.Close()

	ar.writeRequest(conn)
	response, err = ar.readResponse(conn)

	return
}

func (ar *AnnounceRequest) writeRequest(conn net.Conn) {
	conn.Write(ar.Message())
}

func (ar *AnnounceRequest) readResponse(conn net.Conn) (response *AnnounceResponse, err error) {
	buf := make([]byte, 2048)

	readDeadline := time.Now().Add(timeout)
	conn.SetReadDeadline(readDeadline)

	_, err = bufio.NewReader(conn).Read(buf)
	if err != nil {
		return
	}

	response, err = NewAnnounceResponse(buf, ar)

	return
}
