package bittorrent

import (
	"net/url"
	"strconv"
)

type Tracker struct {
	URL *url.URL
}

type TrackerRequest struct {
	Tracker    *Tracker
	InfoHash   string
	PeerID     string
	Port       int
	Uploaded   int
	Downloaded int
	Left       int
	Compact    bool
	NoPeerID   bool
	Event      string
}

type TrackerResponse struct {
}

func NewTracker(url *url.URL) *Tracker {
	return &Tracker{
		URL: url,
	}
}

func (t *Tracker) Request(infoHash string, peerID string, port int) *TrackerRequest {
	return &TrackerRequest{
		Tracker:    t,
		InfoHash:   infoHash,
		PeerID:     peerID,
		Port:       port,
		Uploaded:   0,
		Downloaded: 0,
		Left:       0,
		Compact:    true,
		NoPeerID:   false,
		Event:      "started",
	}
}

func (tr *TrackerRequest) URL() *url.URL {
	requestURL := tr.Tracker.URL

	parameters := url.Values{}
	parameters.Add("info_hash", tr.InfoHash)
	parameters.Add("peer_id", tr.PeerID)
	parameters.Add("port", strconv.Itoa(tr.Port))
	parameters.Add("uploaded", strconv.Itoa(tr.Uploaded))
	parameters.Add("downloaded", strconv.Itoa(tr.Downloaded))
	parameters.Add("left", strconv.Itoa(tr.Left))
	parameters.Add("compact", strconv.Itoa(boolToInt(tr.Compact)))
	parameters.Add("no_peer_id", strconv.Itoa(boolToInt(tr.NoPeerID)))
	parameters.Add("event", tr.Event)
	requestURL.RawQuery = parameters.Encode()

	return requestURL
}

func boolToInt(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
}
