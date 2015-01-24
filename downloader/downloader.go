package downloader

import (
	"fmt"
	"github.com/alexmarchant/torgo/bittorrent"
)

var (
	Port   = 6881
	PeerID = "15620985492012023883"
)

type Downloader struct {
	Torrent      *bittorrent.Torrent
	DownloadPath string
	Peers        []*bittorrent.Peer
}

func NewDownloader(torrentPath string, downloadPath string) (downloader *Downloader, err error) {
	var torrent *bittorrent.Torrent

	torrent, err = bittorrent.NewTorrent(torrentPath)
	if err != nil {
		return
	}

	downloader = &Downloader{
		Torrent:      torrent,
		DownloadPath: downloadPath,
		Peers:        []*bittorrent.Peer{},
	}

	return
}

func (d *Downloader) StartDownload() error {
	err := d.getPeers()
	fmt.Println(len(d.Peers), "peers found.")
	return err
}

func (d *Downloader) getPeers() error {
	for _, tracker := range d.Torrent.Trackers() {
		peers, err := tracker.GetPeersForTorrent(d.Torrent)
		if err != nil {
			fmt.Println("Error getting peers:", err, "... moving to the next tracker")
			continue
		}
		d.addPeers(peers)
	}

	return nil
}

func (d *Downloader) addPeers(newPeers []*bittorrent.Peer) {
	for _, newPeer := range newPeers {
		if !contains(d.Peers, newPeer) {
			d.Peers = append(d.Peers, newPeer)
		}
	}
}

func contains(peers []*bittorrent.Peer, newPeer *bittorrent.Peer) bool {
	for _, peer := range peers {
		if peer.IPAddress == newPeer.IPAddress && peer.TCPPort == newPeer.TCPPort {
			return true
		}
	}
	return false
}
