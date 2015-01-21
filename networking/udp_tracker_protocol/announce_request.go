package udp_tracker_protocol

import (
	"math/rand"
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

func NewAnnounceRequest(connectionID uint64, infoHash string, peerID string, event eventCode, port uint16) *AnnounceRequest {
	return &AnnounceRequest{
		ConnectionID:  connectionID,
		Action:        announceAction,
		TransactionID: rand.Uint32(),
		InfoHash:      infoHash,
		PeerID:        peerID,
		Downloaded:    0,
		Left:          0,
		Uploaded:      0,
		Event:         event,
		IPAddress:     0,
		Key:           0,
		NumWant:       -1,
		Port:          port,
	}
}
