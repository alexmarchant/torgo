package bittorrent

import (
	"errors"
	"github.com/alexmarchant/torgo/bittorrent/networking/udp_tracker_protocol"
	"net/url"
)

type Tracker struct {
	URL *url.URL
}

func NewTracker(url *url.URL) *Tracker {
	return &Tracker{
		URL: url,
	}
}

func (t *Tracker) GetPeersForTorrent(torrent *Torrent, peerID string) (peers []*Peer, err error) {
	switch t.URL.Scheme {
	case "http":
		peers, err = t.GetPeersForTorrentHTTP(torrent, peerID)
	case "udp":
		peers, err = t.GetPeersForTorrentUDP(torrent, peerID)
	default:
		panic("Unrecognized URL scheme")
	}
	return
}

func (t *Tracker) GetPeersForTorrentHTTP(torrent *Torrent, peerID string) (peers []*Peer, err error) {
	err = errors.New("HTTP tracker requests are a WIP")
	return
}

func (t *Tracker) GetPeersForTorrentUDP(torrent *Torrent, peerID string) (peers []*Peer, err error) {
	var connectResponse *udp_tracker_protocol.ConnectResponse
	var announceResponse *udp_tracker_protocol.AnnounceResponse

	connectResponse, err = udp_tracker_protocol.NewConnectRequest().Send(t.URL)
	if err != nil {
		return
	}

	announceResponse, err = udp_tracker_protocol.NewAnnounceRequest(connectResponse, peerID, torrent.InfoHash(), 8888).Send(t.URL)
	if err != nil {
		return
	}

	peers = []*Peer{}

	for _, peerAddress := range announceResponse.PeerAddresses {
		newPeer := NewPeer(peerAddress.IPAddress, peerAddress.TCPPort)
		peers = append(peers, newPeer)
	}

	return
}
