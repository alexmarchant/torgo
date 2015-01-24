package udp_tracker_protocol

import (
	"encoding/binary"
	"errors"
)

type ConnectResponse struct {
	Action        actionCode
	TransactionID uint32
	ConnectionID  uint64
}

func NewConnectResponse(data []byte, request *ConnectRequest) (response *ConnectResponse, err error) {
	if len(data) < 16 {
		err = errors.New("Response too short")
		return
	}

	action := data[0:4]
	transactionID := data[4:8]
	connectionID := data[8:16]

	response = &ConnectResponse{
		Action:        actionCode(binary.BigEndian.Uint32(action)),
		TransactionID: binary.BigEndian.Uint32(transactionID),
		ConnectionID:  binary.BigEndian.Uint64(connectionID),
	}

	if response.TransactionID != request.TransactionID {
		err = errors.New("Wrong transaction ID")
		return
	}

	return
}
