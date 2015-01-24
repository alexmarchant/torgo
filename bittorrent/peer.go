package bittorrent

import ()

type peerState int

const (
	peerNotConnected peerState = iota
	peerConnecting
	peerConnectionFailed
	peerConnected
	peerDownloading
)

type Peer struct {
	IPAddress    uint32
	PeerId       string
	TCPPort      uint16
	State        peerState
	AmChoking    bool
	AmInterested bool
	Choked       bool
	Interested   bool
}

func NewPeer(ipAddress uint32, port uint16) *Peer {
	return &Peer{
		IPAddress:    ipAddress,
		TCPPort:      port,
		State:        peerNotConnected,
		AmChoking:    true,
		AmInterested: false,
		Choked:       true,
		Interested:   false,
	}
}
