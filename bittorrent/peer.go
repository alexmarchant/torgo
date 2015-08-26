package bittorrent

import (
	"net"
	"strconv"
)

type peerState int

const (
	peerNotConnected peerState = iota
	peerConnecting
	peerConnectionFailed
	peerConnected
	peerDownloading
)

type Peer struct {
	IPAddress    net.IP
	PeerId       string
	TCPPort      uint16
	State        peerState
	AmChoking    bool
	AmInterested bool
	Choked       bool
	Interested   bool
}

func NewPeer(ipAddress net.IP, port uint16) *Peer {
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

func (p *Peer) StringAddress() (address string) {
	address = p.IPAddress.String()
	address = address + ":"
	address = address + strconv.Itoa(int(p.TCPPort))
	return
}
