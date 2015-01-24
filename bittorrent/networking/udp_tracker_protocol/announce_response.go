package udp_tracker_protocol

import (
	"encoding/binary"
	"errors"
)

type AnnounceResponse struct {
	Action        actionCode
	TransactionID uint32
	Interval      uint32
	Leechers      uint32
	Seeders       uint32
	PeerAddresses []*PeerAddress
}

func NewAnnounceResponse(data []byte, request *AnnounceRequest) (response *AnnounceResponse, err error) {
	if len(data) < 20 {
		err = errors.New("Response too short")
		return
	}

	action := actionCode(binary.BigEndian.Uint32(data[0:4]))

	if action != announceAction {
		err = errors.New("Unexpected response action")
		return
	}

	transactionID := binary.BigEndian.Uint32(data[4:8])

	if transactionID != request.TransactionID {
		err = errors.New("Wrong transaction ID")
		return
	}

	interval := binary.BigEndian.Uint32(data[8:12])
	leechers := binary.BigEndian.Uint32(data[12:16])
	seeders := binary.BigEndian.Uint32(data[16:20])

	peerAddresses := []*PeerAddress{}
	peerAddressCount := len(data[20:]) / 6

	for i := 0; i < peerAddressCount; i++ {
		offset := 6 * i
		start := 20 + offset
		end := 26 + offset

		newPeerAddress, peerAddressErr := NewPeerAddress(data[start:end])
		if peerAddressErr != nil {
			continue
		}

		peerAddresses = append(peerAddresses, newPeerAddress)
	}

	response = &AnnounceResponse{
		Action:        action,
		TransactionID: transactionID,
		Interval:      interval,
		Leechers:      leechers,
		Seeders:       seeders,
		PeerAddresses: peerAddresses,
	}

	return
}

type PeerAddress struct {
	IPAddress uint32
	TCPPort   uint16
}

func NewPeerAddress(data []byte) (peerAddress *PeerAddress, err error) {
	if len(data) != 6 {
		err = errors.New("Data should be 6 bytes")
		return
	}

	ipAddress := binary.BigEndian.Uint32(data[0:4])
	tcpPort := binary.BigEndian.Uint16(data[4:6])

	if ipAddress == 0 || tcpPort == 0 {
		err = errors.New("Bad peer address data")
		return
	}

	peerAddress = &PeerAddress{
		IPAddress: ipAddress,
		TCPPort:   tcpPort,
	}

	return
}
