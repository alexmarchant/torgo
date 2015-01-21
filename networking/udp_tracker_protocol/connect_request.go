package udp_tracker_protocol

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/alexmarchant/torgo/bittorrent"
	"math/rand"
	"net"
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
	connectionID := make([]byte, 8)
	binary.BigEndian.PutUint64(connectionID, cr.ConnectionID)

	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(cr.Action))

	transactionID := make([]byte, 4)
	binary.BigEndian.PutUint32(transactionID, cr.TransactionID)

	message := append(connectionID, action...)
	message = append(message, transactionID...)
	return message
}

func SendConnectRequest(tr *bittorrent.TrackerRequest) error {
	response := make([]byte, 2048)

	conn, err := net.DialTimeout(tr.URL().Scheme, tr.URL().Host, timeout)
	if err != nil {
		return err
	}

	connectRequest := NewConnectRequest()
	conn.Write(connectRequest.Message())

	readDeadline := time.Now().Add(timeout)
	err = conn.SetReadDeadline(readDeadline)
	if err != nil {
		return err
	}

	fmt.Println("Sending data")

	_, err = bufio.NewReader(conn).Read(response)
	if err != nil {
		fmt.Println(err)
		return err
	}

	connectResponse, err := ParseConnectResponse(response, connectRequest)
	if err != nil {
		return err
	}

	fmt.Println(connectResponse)

	return nil
}
