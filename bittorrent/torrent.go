package bittorrent

import (
	"crypto/sha1"
	// "fmt"
	"github.com/zeebo/bencode"
	"io"
	"net/url"
	"os"
)

type Torrent struct {
	Info         InfoDict   `bencode:"info"`
	Announce     string     `bencode:"announce,omitempty"`
	AnnounceList [][]string `bencode:"announce-list,omitempty"`
	CreationDate int64      `bencode:"creation date,omitempty"`
	Comment      string     `bencode:"comment,omitempty"`
	CreatedBy    string     `bencode:"created by,omitempty"`
	UrlList      string     `bencode:"url-list,omitempty"`
}

type InfoDict struct {
	Name        string `bencode:"name"`
	Length      int    `bencode:"length"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
}

func NewTorrent(torrentPath string) (torrent *Torrent, err error) {
	var file *os.File

	file, err = os.Open(torrentPath)

	if err != nil {
		return
	}

	err = bencode.NewDecoder(file).Decode(&torrent)

	if err != nil {
		panic(err)
	}

	return
}

func (t *Torrent) InfoHash() string {
	str, err := bencode.EncodeString(&t.Info)

	if err != nil {
		panic(err)
	}

	h := sha1.New()
	io.WriteString(h, str)
	return string(h.Sum(nil))
}

func (t *Torrent) Trackers() []*Tracker {
	trackers := []*Tracker{}

	trackerURL, _ := url.Parse(t.Announce)
	trackers = append(trackers, NewTracker(trackerURL))

	for _, trackerAnnounce := range t.AnnounceList {
		trackerURL, _ = url.Parse(trackerAnnounce[0])
		newTracker := NewTracker(trackerURL)
		trackers = append(trackers, newTracker)
	}

	return trackers
}
