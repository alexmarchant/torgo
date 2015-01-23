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
	Ip           string
	PeerId       string
	Port         int
	State        peerState
	AmChoking    bool
	AmInterested bool
	Choked       bool
	Interested   bool
}

func NewPeer() *Peer {
	return &Peer{
		State:        peerNotConnected,
		AmChoking:    true,
		AmInterested: false,
		Choked:       true,
		Interested:   false,
	}
}
