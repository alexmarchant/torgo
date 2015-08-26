package downloader

import (
	"fmt"
	"github.com/alexmarchant/torgo/bittorrent"
	"reflect"
)

const (
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
	if err != nil {
		return err
	}

	fmt.Println(len(d.Peers), "peers found.")

	if len(d.Peers) > 0 {
		err = d.downloadFromPeer(d.Peers[0])
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Downloader) getPeers() error {
	for _, tracker := range d.Torrent.Trackers() {
		peers, err := tracker.GetPeersForTorrent(d.Torrent, PeerID)
		if err != nil {
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
		if reflect.DeepEqual(peer.IPAddress, newPeer.IPAddress) && peer.TCPPort == newPeer.TCPPort {
			return true
		}
	}
	return false
}

func (d *Downloader) downloadFromPeer(peer *bittorrent.Peer) error {
	return nil
}
